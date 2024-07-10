package keeper

import (
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/circuit"
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
func init() {

}
func (k *Keeper) Vote(ctx context.Context, votePropal types.MsgVoteProposal) error {

	nullifier := votePropal.Nullifier
	proposalID := votePropal.ProposalId
	voteOption := votePropal.VoteOption

	zkProofBytes := votePropal.ZkProof

	store := k.storeKey.OpenKVStore(ctx)

	merkleRoot, err := types.GetMerkleRoot(ctx, store, proposalID)
	if err != nil {
		return err
	}

	err = types.StoreNullifier(ctx, store, proposalID, nullifier)
	if err != nil {
		return err
	}

	//
	publicWitness := circuit.PreparePublicWitness(nullifier, uint64(voteOption), merkleRoot)
	zkProof, err := circuit.UnMarshalZkProof(zkProofBytes[:])
	if err != nil {
		return err
	}

	// verifier key should be initialized at genesis
	var vkey groth16.VerifyingKey
	groth16.Verify(zkProof, vkey, publicWitness)

	// TODO: process the vote...

	return nil
}
func (k *Keeper) RegisterUser(ctx context.Context, commitment string, user string, proposalID uint64) error {
	store := k.storeKey.OpenKVStore(ctx)
	if err := types.StoreUser(ctx, store, proposalID, user); err != nil {
		return err
	}

	if err := types.StoreCommitment(ctx, store, proposalID, commitment); err != nil {
		return err
	}

	return nil
}

// func (k *Keeper) nextSequence(ctx context.Context, key []byte) (uint64, error) {
// 	store := k.storeKey.OpenKVStore(ctx)
// 	found, err := store.Has(key)
// 	fmt.Println("The found value is>>>>>>>>>>.", found)
// 	if err != nil {
// 		return 0, err
// 	}
// 	var seq uint64 = 1
// 	if found {
// 		pvBytes, err := store.Get(key)
// 		if err != nil {
// 			return 0, err
// 		}
// 		seq = binary.BigEndian.Uint64(pvBytes) + 1
// 	}
// 	seqBytes := uint64ToBytes(seq)
// 	store.Set(key, seqBytes)
// 	return seq, nil
// }

// func (k *Keeper) SetSequence(ctx context.Context, seq uint64) {
// 	store := k.storeKey.OpenKVStore(ctx)
// 	seqBytes := uint64ToBytes(seq)
// 	store.Set(UserSeqPrefix, seqBytes)
// }
// func uint64ToBytes(value uint64) []byte {
// 	seqBytes := make([]byte, 8)
// 	binary.BigEndian.PutUint64(seqBytes, value)
// 	return seqBytes
// }
