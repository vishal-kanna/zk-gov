package store

import (
	"bytes"
	"context"
	"errors"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func InitNullifiers(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	nullifiersKey := types.NullifiersStoreKey(proposalID)

	nullifiersBytes := []byte{}

	return store.Set(nullifiersKey, nullifiersBytes)

}

func StoreNullifier(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, nullifier string) error {
	nullifiersKey := types.NullifiersStoreKey(proposalID)
	nullifiersBytes, err := store.Get(nullifiersKey)
	if err != nil {
		return err
	}

	nullifierBytes, err := types.HexStringToBytes(nullifier)
	if err != nil {
		return err
	}

	// if nullifier already stored, the vote is already processed
	for i := 0; i < len(nullifiersBytes); i += types.NULLIFIER_SIZE {
		storedNullifier := nullifiersBytes[i*types.NULLIFIER_SIZE : (i+1)*types.NULLIFIER_SIZE]
		if bytes.Equal(storedNullifier, nullifierBytes) {
			return errors.New("the user is already voted")
		}
	}

	nullifiersBytes = append(nullifiersBytes, nullifierBytes...)
	if err := store.Set(nullifiersKey, nullifiersBytes); err != nil {
		return err
	}

	return nil
}
