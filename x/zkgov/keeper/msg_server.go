package keeper

import (
	"context"
	"fmt"

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

func (k msgServer) RegisterUser(ctx context.Context, req *types.RegisterUserRequest) (*types.RegisterUserResponse, error) {
	fmt.Println("I'm here>>>>>>>>>>>.")
	err := k.RegisterUsr(ctx)
	if err != nil {
		return &types.RegisterUserResponse{}, err
	}
	return &types.RegisterUserResponse{}, nil
}

// generate userId it should in seq
// generate the random number to get the nullifier
//
