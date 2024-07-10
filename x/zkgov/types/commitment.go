package types

import (
	"context"

	cosmosstore "cosmossdk.io/core/store"
)

func StoreCommitment(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitment string) error {
	commitmentsKey := CommitmentsStoreKey(proposalID)
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
