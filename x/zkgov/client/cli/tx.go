package cli

// import (
// 	"bytes"
// 	"fmt"
// 	"math/big"
// 	"math/rand"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"time"

// 	"github.com/consensys/gnark-crypto/accumulator/merkletree"
// 	"github.com/consensys/gnark-crypto/ecc"
// 	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
// 	"github.com/consensys/gnark/frontend"
// 	"github.com/consensys/gnark/std/accumulator/merkle"
// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// 	voting "github.com/vishal-kanna/zk/zk-gov/x/zkgov/circuit"
// 	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
// )

// var FlagSplit = "split"

// // NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
// func NewTxCmd() *cobra.Command {
// 	txCmd := &cobra.Command{
// 		Use:                        types.ModuleName,
// 		Short:                      "Bank transaction subcommands",
// 		DisableFlagParsing:         true,
// 		SuggestionsMinimumDistance: 2,
// 		RunE:                       client.ValidateCmd,
// 	}

// 	txCmd.AddCommand(
// 		// NewRegisterCmd(),
// 		NewUserRegisterCmd(),
// 	)

// 	return txCmd
// }
// func NewRegisterCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "register",
// 		Short: "Register a new Voter",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			userId := viper.GetInt64("userId")
// 			if userId > 0 {
// 				fmt.Println("UserId is required")
// 				return err
// 			}

// 			randomId := getRandomNumber()
// 			commitment, nullifier := createCommitmentAndNullifier(userId, randomId)

// 			cmtBuf, nulBuf := bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})
// 			cmtBuf.Write(commitment)
// 			nulBuf.Write(nullifier)

// 			// store the above voterInfo
// 			userPath := strconv.FormatInt(userId, 10)
// 			err = SaveInfo("commitment_"+userPath, cmtBuf.Bytes())
// 			if err != nil {
// 				fmt.Println("Error writing commitment to file:", err.Error())
// 				return err
// 			}

// 			err = SaveInfo("nullifier_"+userPath, nulBuf.Bytes())
// 			if err != nil {
// 				fmt.Println("Error writing nullifier to file:", err.Error())
// 				return err
// 			}

// 			err = SaveInfo("voterId_"+userPath, []byte(fmt.Sprintf("%d\n", randomId)))
// 			if err != nil {
// 				fmt.Println("Error writing to file:", err.Error())
// 				return err
// 			}
// 			msg := types.RegisterCommitmentRequest{Commitment: string(commitment)}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 		},
// 	}

// 	cmd.Flags().Bool(FlagSplit, false, "Send the equally split token amount to each address")
// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }
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

// // Generate a random 5-digit salt
// func getRandomNumber() int64 {
// 	seed := time.Now().UnixNano()
// 	rng := rand.New(rand.NewSource(seed))

// 	return int64(rng.Intn(10000))
// }

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
