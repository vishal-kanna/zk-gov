package keeper

import (
	"context"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}
func (k msgServer) RegisterCommitment(ctx context.Context, req *types.RegisterCommitmentRequest) (*types.RegisterCommitmentResponse, error) {
	if req.Commitment == "" {
		return nil, types.EmptyCommitment
	}
	err := k.StoreCommitment(ctx, req.Commitment)
	if err != nil {
		return nil, nil
	}
	return &types.RegisterCommitmentResponse{}, nil
}
