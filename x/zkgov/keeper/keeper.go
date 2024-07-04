package keeper

import (
	"context"
	"encoding/binary"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
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

func (k *Keeper) StoreCommitment(ctx context.Context, commitment string) error {

	store := k.storeKey.OpenKVStore(ctx)
	seq, err := k.nextPlanSequence(ctx)
	if err != nil {
		return err
	}
	commitments := types.Commitment{
		Commitment:   commitment,
		CommitmentId: seq,
	}
	bz, err := k.cdc.Marshal(&commitments)
	if err != nil {
		return err
	}
	return store.Set(CommitmentStoreKey(seq), bz)
}

func (k *Keeper) nextPlanSequence(ctx context.Context) (uint64, error) {
	store := k.storeKey.OpenKVStore(ctx)
	found, err := store.Has(CommitmentSeqPrefix)
	if err != nil {
		return 0, err
	}
	var seq uint64 = 1
	if found {
		pvBytes, err := store.Get(CommitmentSeqPrefix)
		if err != nil {
			return 0, err
		}
		seq = binary.BigEndian.Uint64(pvBytes) + 1
	}
	seqBytes := uint64ToBytes(seq)
	store.Set(CommitmentSeqPrefix, seqBytes)
	return seq, nil
}
func (k *Keeper) setPlanSequence(ctx context.Context, seq uint64) {
	store := k.storeKey.OpenKVStore(ctx)
	seqBytes := uint64ToBytes(seq)
	store.Set(CommitmentSeqPrefix, seqBytes)
}
func uint64ToBytes(value uint64) []byte {
	seqBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(seqBytes, value)
	return seqBytes
}
