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

	fmt.Println(1)
	nullifier := votePropal.Nullifier
	proposalID := votePropal.ProposalId
	voteOption := votePropal.VoteOption

	zkProofBytes := votePropal.ZkProof

	store := k.storeKey.OpenKVStore(ctx)

	merkleRoot, err := storeImpl.GetMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return err
	}

	fmt.Println(2)

	err = storeImpl.StoreNullifier(ctx, store, proposalID, nullifier)
	if err != nil {
		return err
	}

	fmt.Println(3)

	err = storeImpl.StoreVote(ctx, store, proposalID, voteOption)
	if err != nil {
		return err
	}

	fmt.Println(4)

	publicWitness := circuit.PreparePublicWitness(nullifier, uint64(voteOption), merkleRoot)
	zkProof, err := circuit.UnMarshalZkProof(zkProofBytes[:])
	if err != nil {
		return err
	}

	fmt.Println(5)

	// verifier key should be initialized at genesis
	vkey, err := circuit.FetchVerifier(int(votePropal.MerkleproofSize))
	if err != nil {
		return err
	}

	fmt.Println(6)

	err = groth16.Verify(zkProof, vkey, publicWitness)
	if err != nil {
		return err
	}

	fmt.Println(7)

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

func (k *Keeper) GetProposalAllInfo(ctx context.Context, req *types.QueryProposalAllInfoRequest) (*types.QueryProposalAllInfoResponse, error) {
	store := k.storeKey.OpenKVStore(ctx)

	// querying the stored proposal
	proposal, err := storeImpl.GetProposal(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// query the registered users
	users, err := storeImpl.GetUsers(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// query the store votes
	votes, err := storeImpl.GetVotes(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// query the commitments
	commitments, err := storeImpl.GetCommitments(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// query the nullifier
	nullifiers, err := storeImpl.GetNullifiers(ctx, store, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// fmt.Println("commitments:", commitments)
	// fmt.Println("nullifiers:", nullifiers)
	// fmt.Println("votes:", votes)
	// fmt.Println("users:", users)

	usersInfo := GetUsersInfo(commitments, users)
	votesInfo := GetVotesInfo(nullifiers, votes)

	return &types.QueryProposalAllInfoResponse{
		Title:                proposal.Title,
		Description:          proposal.Description,
		RegistrationDeadline: proposal.RegistrationDeadline,
		VotingDeadline:       proposal.VotingDeadline,
		RegisteredUsers:      usersInfo,
		Votes:                votesInfo,
	}, nil
}

func GetVotesInfo(nullifiers []string, votes []types.VoteOption) []*types.VoteInfo {
	VotesInfo := make([]*types.VoteInfo, len(nullifiers))
	for i, vote := range votes {
		voteInfo := &types.VoteInfo{
			Nullifer:   nullifiers[i],
			VoteOption: vote,
		}
		VotesInfo[i] = voteInfo
	}

	return VotesInfo
}

func GetUsersInfo(commitments []string, users []string) []*types.UserInfo {
	usersInfo := make([]*types.UserInfo, len(users))
	for i, user := range users {
		userInfo := &types.UserInfo{
			Commitment:  commitments[i],
			UserAddress: user,
		}
		usersInfo[i] = userInfo

	}

	return usersInfo
}
