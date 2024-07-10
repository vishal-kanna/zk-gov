package types

import (
	"bytes"
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

func UpdateMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitments []byte) error {
	merklerootKey := MerkleRootStoreKey(proposalID)

	var buf bytes.Buffer
	buf.Write(commitments)

	hFunc := mimc.NewMiMC()

	// TODO: find root directly
	merkleroot, _, _, err := merkletree.BuildReaderProof(&buf, hFunc, COMMITMENT_SIZE, uint64(0))
	if err != nil {
		return err
	}

	store.Set(merklerootKey, merkleroot)
	return nil
}

func GetMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) (string, error) {
	merkleKey := MerkleRootStoreKey(proposalID)
	merkleRootBytes, err := store.Get(merkleKey)
	if err != nil {
		return "", err
	}
	return string(merkleRootBytes), nil
}
