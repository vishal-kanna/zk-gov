package cli

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cosmos_types "github.com/cosmos/gogoproto/types"
	"github.com/spf13/cobra"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/circuit"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

var FlagSplit = "split"

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Zk-gov transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		// NewRegisterCmd(),
		NewRegisterVoteCmd(),
		NewCreateProposalCmd(),
		NewVote(),
	)

	return txCmd
}

// takes the proposal id and his option vote
func NewRegisterVoteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-vote proposal-id voteOption ",
		Short:   "Register a new Voter",
		Example: "simd tx zk-gov register-vote 1 1 --from alice --keyring-backend test --chain-id testnet",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress().String()
			proposalID := args[0]
			voteOption := args[1]
			Pid, _ := strconv.Atoi(proposalID)
			vote, _ := strconv.Atoi(voteOption)

			randomSecret1 := getRandomNumber()
			randomSecret2 := getRandomNumber()

			commitment := circuit.CreateCommitment(randomSecret1, randomSecret2, int64(vote))
			nullifier := circuit.CreateNullifier(randomSecret2, int64(vote))

			err = circuit.SaveInfo(uint64(Pid), commitment, nullifier, uint64(vote), uint64(randomSecret1), uint64(randomSecret2))
			if err != nil {
				fmt.Println("Error while saving to file:", err.Error())
				return err
			}

			commitmentString := types.BytesToHexString(commitment)
			msg := types.MsgRegisterUser{ProposalId: uint64(Pid), Sender: sender, Commitment: commitmentString}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// create new proposal
func NewCreateProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-proposal [title] [description] ",
		Short:   "Create a new Proposal",
		Example: "simd tx zk-gov create-proposal dummy-proposal dummy-description --from alice --keyring-backend test --chain-id demo",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress().String()
			title := args[0]
			description := args[1]
			registration_deadline := time.Now().Add(time.Hour)
			registration_deadline_timestamp, err := cosmos_types.TimestampProto(registration_deadline)
			if err != nil {
				return err
			}

			voting_deadline := registration_deadline.Add(time.Hour)
			voting_deadline_timestamp, err := cosmos_types.TimestampProto(voting_deadline)
			if err != nil {
				return err
			}

			msg := types.MsgCreateProposal{
				Title:                title,
				Sender:               sender,
				Description:          description,
				RegistrationDeadline: registration_deadline_timestamp,
				VotingDeadline:       voting_deadline_timestamp,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// get the proposal_state_root zk_proof and proposal_state_root
// get nullifier and commitment from the store
func NewVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote [proposal-id]",
		Short: "Create a vote transaction for previously generated values (register-vote)",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			proposalID := args[0]
			Pid, _ := strconv.Atoi(proposalID)
			VoterInfo, err := circuit.FetchInfo(proposalID)
			if err != nil {
				fmt.Println("Error while fetching file:", err.Error())
				return err
			}

			nullifier := VoterInfo.Nullifier
			nullifierBytes, _ := types.HexStringToBytes(nullifier)
			voteOption := *big.NewInt(int64(VoterInfo.VoteOption))
			commitment := VoterInfo.Commitment
			commitmentBytes, _ := types.HexStringToBytes(commitment)
			randomSecret1 := *big.NewInt(int64(VoterInfo.RandomSecret1))
			randomSecret2 := *big.NewInt(int64(VoterInfo.RandomSecret2))
			sender := clientCtx.GetFromAddress().String()

			var opt types.VoteOption
			if VoterInfo.VoteOption == 0 {
				opt = types.VoteOption_VOTE_OPTION_YES
			} else {
				opt = types.VoteOption_VOTE_OPTION_NO
			}

			// merkle proof request
			var req types.QueryCommitmentMerkleProofRequest

			req.Commitment = commitment
			req.ProposalId = uint64(Pid)
			res, err := queryClient.CommitmentMerkleProof(cmd.Context(), &req)
			if err != nil {
				fmt.Println("Error while questing MerkleProof", err.Error())
			}

			merkleroot := res.GetRoot()
			merklerootString := types.BytesToHexString(merkleroot)
			merkleproofBytes := res.GetMerkleProof()
			merkleproof := store.GetMerkleProofFromBytes(merkleroot, merkleproofBytes)
			commitmentIndex := res.GetCommitmentIndex()
			merkleproofSize := len(merkleproof.Path)

			assignment := circuit.PrivateVotingCircuit{
				SecretUniqueId1: randomSecret1,
				SecretUniqueId2: randomSecret2,
				Commitment:      commitmentBytes,
				Nullifier:       nullifierBytes,
				VoteOption:      voteOption,
				CommitmentIndex: commitmentIndex,
				MerkleRoot:      merkleroot,
				MerkleProof:     merkleproof,
			}

			fmt.Println("secret1:", VoterInfo.RandomSecret1,

				"secret2:", randomSecret2,
				"comm", commitmentBytes,
				"null", nullifierBytes,
				"vote option", voteOption,
				"index", commitmentIndex,
				"merkleroot", merkleroot,
				"merkleproof", merkleproof,
			)

			circuit.TestZKProof(&assignment)

			zkProofBytes, err := circuit.GenerateProof(&assignment)

			msg := types.MsgVoteProposal{
				ProposalId:        uint64(Pid),
				Nullifier:         nullifier,
				VoteOption:        opt,
				Sender:            sender,
				ZkProof:           zkProofBytes,
				ProposalStateRoot: merklerootString,
				MerkleproofSize:   uint64(merkleproofSize),
			}

			fmt.Println("cli done.........")

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// func GenerateProofCmd() *cobra.Command {

// 	generateProofCmd := &cobra.Command{
// 		Use:   "generate-proof",
// 		Short: "Generate a proof",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			fmt.Println(">>>>>>>>>>>>>>.", clientCtx)
// 			userIdStr, _ := cmd.Flags().GetString("userId")
// 			userId, err := strconv.ParseInt(userIdStr, 10, 64)
// 			if userId <= 0 {
// 				fmt.Println("userId is required")
// 				return err
// 			}

// 			userPath := strconv.FormatInt(userId, 10)
// 			commitment, err := FetchInfo("commitment_" + userPath)
// 			if err != nil {
// 				fmt.Println("Error reading commitment", err)
// 				return err
// 			}

// 			nullifier, err := FetchInfo("nullifier_" + userPath)
// 			if err != nil {
// 				fmt.Println("Error reading nullifier", err)
// 				return err
// 			}

// 			voterId, err := FetchInfo("voterId_" + userPath)
// 			if err != nil {
// 				fmt.Println("Error reading voterId", err)
// 				return err
// 			}

// 			var buf bytes.Buffer
// 			// build merkle proof
// 			dataSegments := 4
// 			proofIndex := 0
// 			dataSize := len(commitment)
// 			hFunc := mimc.NewMiMC()
// 			for j := byte(1); j <= byte(dataSegments); j++ {
// 				data := commitment
// 				hFunc.Reset()
// 				hFunc.Write(data)
// 				hash := hFunc.Sum(nil)

// 				_, err := buf.Write(hash)
// 				if err != nil {
// 					fmt.Println("failed to write hash", err)
// 				}
// 			}

// 			root, proof, numLeaves, err := merkletree.BuildReaderProof(&buf, hFunc, dataSize, uint64(proofIndex))
// 			verified := merkletree.VerifyProof(hFunc, root, proof, uint64(proofIndex), numLeaves)
// 			if verified {
// 				fmt.Println("Proof is generated and verified")
// 			}

// 			// Define the inputs
// 			assignment := voting.Circuit{
// 				UniqueId1:   userId,
// 				ProofIndex:  0,
// 				UniqueId2:   voterId,
// 				Commitment:  commitment,
// 				Nullifier:   nullifier,
// 				MerkleProof: GetMerkleProofFromBytes(root, proof),
// 				MerkleRoot:  root,
// 				VoteOption:  1,
// 			}
// 			witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
// 			publicWitness, err := witness.Public()

// 			if err != nil {
// 				fmt.Println("failed to generate witness", err.Error())
// 				return err
// 			}
// 			err = SaveInfo("witness_"+userPath, []byte(fmt.Sprintf("%v", publicWitness)))
// 			if err != nil {
// 				fmt.Println("Error writing nullifier to file:", err.Error())
// 				return err
// 			}
// 			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 			return nil

// 		},
// 	}
// 	generateProofCmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
// 	flags.AddTxFlagsToCmd(generateProofCmd)
// 	return generateProofCmd
// }

// func NewUserRegisterCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "Register your user",
// 		Short: "Register a new Voter",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			sender := clientCtx.GetFromAddress().String()
// 			msg := types.RegisterUserRequest{Sender: sender}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 		},
// 	}

// 	cmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

// Generate a random 5-digit salt
func getRandomNumber() int64 {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	return int64(rng.Intn(10000))
}

// // To create a commitment and nullifier
// func createCommitmentAndNullifier(userId, randomId int64) ([]byte, []byte) {
// 	hFunc := mimc.NewMiMC()

// 	// Create commitment
// 	hFunc.Write(big.NewInt(userId).Bytes())
// 	hFunc.Write(big.NewInt(randomId).Bytes())
// 	commitment := hFunc.Sum(nil)
// 	hFunc.Reset()

// 	// Create nullifier
// 	hFunc.Write(big.NewInt(randomId).Bytes())
// 	nullifier := hFunc.Sum(nil)

// 	return commitment, nullifier
// }

// // To save voter info
// func SaveInfo(filename string, data []byte) error {
// 	return os.WriteFile(filepath.Join("voters", filename), data, 0666)
// }

// // to fetch voter info
// func FetchInfo(filename string) ([]byte, error) {
// 	info, err := os.ReadFile(filepath.Join("voters", filename))
// 	return info, err
// }

// // to get merkle proof from root+proof bytes
// func GetMerkleProofFromBytes(rootBytes []byte, proofBytes [][]byte) merkle.MerkleProof {
// 	var merkleProof merkle.MerkleProof
// 	merkleProof.RootHash = rootBytes
// 	merkleProof.Path = make([]frontend.Variable, len(proofBytes))
// 	for i := 0; i < len(proofBytes); i++ {
// 		merkleProof.Path[i] = proofBytes[i]
// 	}
// 	return merkleProof
// }
