package keeper

import (
	"github.com/evmos/evmos/v6/x/atom/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

var _ types.QueryServer = Keeper{}


// AtomUsd implements the Query/AtomUsd gRPC method
func (k Keeper) AtomUsd(c context.Context, req *types.QueryAtomUsdRequest) (*types.QueryAtomUsdResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	atomUsd := k.GetAtomUsd(ctx)
	if atomUsd == nil {
		return nil, status.Errorf(codes.NotFound, "No results")
	}

	return &types.QueryAtomUsdResponse{AtomUsd: atomUsd}, nil
}