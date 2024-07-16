package store

import (
	"context"
	"encoding/binary"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func StoreProposal(ctx context.Context, store cosmosstore.KVStore, proposal types.MsgCreateProposal) (uint64, error) {

	proposalCounter := GetProposalCounter(ctx, store)
	proposalCounter++
	StoreProposalCounter(ctx, store, proposalCounter)

	proposalInfoStoreKey := types.ProposalInfoStoreKey(proposalCounter)

	proposalBytes, err := proposal.Marshal()
	if err != nil {
		return 0, err
	}
	err = store.Set(proposalInfoStoreKey, proposalBytes)
	if err != nil {
		return 0, err
	}

	return proposalCounter, nil
}

func GetProposalCounter(ctx context.Context, store cosmosstore.KVStore) uint64 {
	proposalCounterKey := types.ProposalCounterKey
	var proposalCounterBytes []byte
	if found, err := store.Has(proposalCounterKey); !found || err != nil {
		zero := 0
		zeroBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(zeroBytes, uint64(zero))
		proposalCounterBytes = zeroBytes
		store.Set(proposalCounterKey, proposalCounterBytes)
	}

	proposalCounterBytes, _ = store.Get(proposalCounterKey)

	proposalCounter := binary.BigEndian.Uint64(proposalCounterBytes)

	return proposalCounter
}

func StoreProposalCounter(ctx context.Context, store cosmosstore.KVStore, proposalCounter uint64) error {
	proposalCounterKey := types.ProposalCounterKey

	proposalCounterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalCounterBytes, proposalCounter)

	return store.Set(proposalCounterKey, proposalCounterBytes)
}
