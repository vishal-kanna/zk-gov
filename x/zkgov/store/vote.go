package store

import (
	"context"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func InitVotes(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	votesKey := types.VotesStoreKey(proposalID)

	votesBytes := []byte{}

	return store.Set(votesKey, votesBytes)

}

func StoreVote(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, voteOption types.VoteOption) error {
	votesKey := types.VotesStoreKey(proposalID)
	Votes, err := store.Get(votesKey)
	if err != nil {
		return err
	}

	voteOptionBytes := types.MarshalVoteOption(voteOption)

	Votes = append(Votes, voteOptionBytes...)
	if err := store.Set(votesKey, Votes); err != nil {
		return err
	}

	return nil
}
