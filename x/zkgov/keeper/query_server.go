package keeper

import (
	"context"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CommitmentMerkleProof(ctx context.Context, req *types.QueryCommitmentMerkleProofRequest) (*types.QueryCommitmentMerkleProofResponse, error) {
	// TODO: implement
	merkleproof, err := k.MerkleProof(ctx, req.ProposalId)
	if err != nil {
		return &types.QueryCommitmentMerkleProofResponse{}, nil
	}

	return &types.QueryCommitmentMerkleProofResponse{
		MerkleProof: merkleproof,
	}, nil

}
