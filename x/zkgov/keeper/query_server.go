package keeper

import (
	"context"
	"fmt"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) GetUser(ctx context.Context, req *types.QueryUserRequset) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	usr, err := q.GetUserInfo(ctx, int64(req.Userid))
	fmt.Println("The user is>>>>>>>>>>.", usr)
	if err != nil {
		return &types.QueryGetUserResponse{}, err
	}
	return &types.QueryGetUserResponse{Ust: &usr}, nil
}
