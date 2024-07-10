package keeper

import (
	"context"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) CommitmentMerkleProof(ctx context.Context, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	// TODO: implement
	return &types.QueryCommitmentMerkleProofResponse{}, nil
}
