package store

import (
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func StoreCommitment(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitment string) error {
	commitmentsKey := types.CommitmentsStoreKey(proposalID)
	commitmentsBytes, err := store.Get(commitmentsKey)
	if err != nil {
		return err
	}

	commitmentBytes := []byte(commitment)
	commitmentsBytes = append(commitmentsBytes, commitmentBytes...)
	store.Set(commitmentsKey, commitmentsBytes)

	UpdateMerkleRoot(ctx, store, proposalID, commitmentsBytes)
	return nil
}
