package keeper

import (
	"context"

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
