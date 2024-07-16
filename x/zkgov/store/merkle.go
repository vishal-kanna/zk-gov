package store

import (
	"bytes"
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func UpdateMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, commitments []byte) error {
	merklerootKey := types.MerkleRootStoreKey(proposalID)

	var buf bytes.Buffer
	buf.Write(commitments)

	hFunc := mimc.NewMiMC()

	// TODO: find root directly
	merkleroot, _, _, err := merkletree.BuildReaderProof(&buf, hFunc, types.COMMITMENT_SIZE, uint64(0))
	if err != nil {
		return err
	}

	store.Set(merklerootKey, merkleroot)
	return nil
}

func GetMerkleProof(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) ([][]byte, error) {
	commitmentsKey := types.CommitmentsStoreKey(proposalID)
	commitmentsBytes, err := store.Get(commitmentsKey)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.Write(commitmentsBytes)

	hFunc := mimc.NewMiMC()

	_, merkleproof, _, err := merkletree.BuildReaderProof(&buf, hFunc, types.COMMITMENT_SIZE, uint64(0))
	if err != nil {
		return nil, err
	}

	return merkleproof, nil
}

func GetMerkleRoot(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) (string, error) {
	merkleKey := types.MerkleRootStoreKey(proposalID)
	merkleRootBytes, err := store.Get(merkleKey)
	if err != nil {
		return "", err
	}
	return string(merkleRootBytes), nil
}

func GetMerkleProofFromBytes(rootBytes []byte, proofBytes [][]byte) merkle.MerkleProof {
	var merkleProof merkle.MerkleProof
	merkleProof.RootHash = rootBytes
	merkleProof.Path = make([]frontend.Variable, len(proofBytes))
	for i := 0; i < len(proofBytes); i++ {
		merkleProof.Path[i] = proofBytes[i]
	}
	return merkleProof
}
