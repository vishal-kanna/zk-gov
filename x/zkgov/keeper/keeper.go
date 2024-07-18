package keeper

import (
	"context"
	"fmt"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/circuit"
	storeImpl "github.com/vishal-kanna/zk/zk-gov/x/zkgov/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type Keeper struct {
	storeKey cosmosstore.KVStoreService
	cdc      codec.BinaryCodec
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey cosmosstore.KVStoreService,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k *Keeper) Vote(ctx context.Context, votePropal types.MsgVoteProposal) error {

	nullifier := votePropal.Nullifier
	proposalID := votePropal.ProposalId
	voteOption := votePropal.VoteOption

	zkProofBytes := votePropal.ZkProof

	store := k.storeKey.OpenKVStore(ctx)

	merkleRoot, err := storeImpl.GetMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return err
	}

	err = storeImpl.StoreNullifier(ctx, store, proposalID, nullifier)
	if err != nil {
		return err
	}

	err = storeImpl.StoreVote(ctx, store, proposalID, voteOption)
	if err != nil {
		return err
	}

	publicWitness := circuit.PreparePublicWitness(nullifier, uint64(voteOption), merkleRoot)
	zkProof, err := circuit.UnMarshalZkProof(zkProofBytes[:])
	if err != nil {
		return err
	}

	// verifier key should be initialized at genesis
	vkey, err := circuit.FetchVerifier(int(votePropal.MerkleproofSize))
	if err != nil {
		return err
	}

	err = groth16.Verify(zkProof, vkey, publicWitness)
	if err != nil {
		return err
	}

	// TODO: process the vote...

	return nil
}

func (k *Keeper) RegisterUser(ctx context.Context, commitment string, user string, proposalID uint64) error {
	store := k.storeKey.OpenKVStore(ctx)
	if err := storeImpl.StoreUser(ctx, store, proposalID, user); err != nil {
		return err
	}

	if err := storeImpl.StoreCommitment(ctx, store, proposalID, commitment); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) CreatePropsal(ctx context.Context, proposal types.MsgCreateProposal) (uint64, error) {
	store := k.storeKey.OpenKVStore(ctx)

	proposalID, err := storeImpl.StoreProposal(ctx, store, proposal)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitCommitments(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitNullifiers(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitUsers(ctx, store, proposalID)
	if err != nil {
		return proposalID, err
	}

	err = storeImpl.InitVotes(ctx, store, proposalID)

	return proposalID, err

}

/* ------------------------- Queries ------------------------------*/

func (k *Keeper) MerkleProof(ctx context.Context, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	store := k.storeKey.OpenKVStore(ctx)
	return storeImpl.GetMerkleProof(ctx, store, req)
}

func (k *Keeper) GetAllDetails(ctx context.Context, req *types.QueryProposalAllInfoRequest) (*types.QueryProposalAllInfoResponse, error) {
	store := k.storeKey.OpenKVStore(ctx)

	// querying the stored proposal
	proposalStoreKey := types.ProposalInfoStoreKey(req.ProposalId)
	var proposal types.MsgCreateProposal
	proposalInfo, err := store.Get(proposalStoreKey)
	if err != nil {
		return nil, err
	}
	err = k.cdc.Unmarshal(proposalInfo, &proposal)
	if err != nil {
		return nil, err
	}

	// query the registered users
	userStoreKey := types.UsersStoreKey(req.ProposalId)
	storedUsers, err := store.Get(userStoreKey)
	if err != nil {
		return nil, err
	}
	users := make([]string, 0)
	for i := 0; i < len(storedUsers); i += types.USER_SIZE {
		user := storedUsers[i : i+types.USER_SIZE]
		users = append(users, string(user))
	}
	fmt.Println("users are>>>>>>>>>>>>.", users)

	// query the store votes
	votesStorekey := types.VotesStoreKey(req.ProposalId)
	storedVotes, err := store.Get(votesStorekey)
	if err != nil {
		return nil, err
	}
	votes := make([]types.VoteOption, 0)
	for i := 0; i < len(storedVotes); i += types.VOTE_SIZE {
		vote := storedUsers[i : i+types.VOTE_SIZE]
		v := types.UnMarshalVoteOption(vote)
		votes = append(votes, v)
	}

	// query the commitments
	commitmentStoreKey := types.CommitmentsStoreKey(req.ProposalId)
	StoredCommitments, err := store.Get(commitmentStoreKey)
	if err != nil {
		return nil, err
	}
	commitmentsArray := make([]string, 0)
	for i := 0; i < len(StoredCommitments); i += types.COMMITMENT_SIZE {
		commitmentBytes := StoredCommitments[i : i+types.COMMITMENT_SIZE]
		commitmentString := types.BytesToHexString(commitmentBytes)
		commitmentsArray = append(commitmentsArray, commitmentString)
	}

	// query the nullifier
	nullifierStorekey := types.NullifiersStoreKey(req.ProposalId)
	StoredNullifiers, err := store.Get(nullifierStorekey)
	if err != nil {
		return nil, err
	}
	nullifierArray := make([]string, 0)
	for i := 0; i < len(StoredNullifiers); i += types.NULLIFIER_SIZE {
		nullifierBytes := StoredNullifiers[i : i+types.NULLIFIER_SIZE]
		nullifierString := types.BytesToHexString(nullifierBytes)
		nullifierArray = append(nullifierArray, nullifierString)
	}

	fmt.Println("nullifier array is>>>>>>>>", nullifierArray)
	fmt.Println("commitments array is>>>>>>", commitmentsArray)
	fmt.Println("Store votes are>>>>>>>", votes)
	fmt.Println("Stored users >>>>>>>.", storedUsers)
	fmt.Println("The proposalInfo is>>>>>>>>>>>>", proposal)

	UserInfo := make([]*types.UserInfo, 0)
	for i, user := range users {
		UserInfo[i].UserAddress = user
	}
	for i, commitment := range commitmentsArray {
		UserInfo[i].Commitment = commitment
	}

	VoteInfo := make([]*types.VoteInfo, 0)
	for i, vote := range votes {
		VoteInfo[i].VoteOption = vote
	}

	for i, nullifer := range nullifierArray {
		VoteInfo[i].Nullifer = nullifer
	}
	return &types.QueryProposalAllInfoResponse{
		Title:                proposal.Title,
		Description:          proposal.Description,
		RegistrationDeadline: proposal.RegistrationDeadline,
		VotingDeadline:       proposal.VotingDeadline,
		RegisteredUsers:      UserInfo,
		Votes:                VoteInfo,
	}, nil
}
