package keeper

import (
	"context"

	"healthcheck/x/healthcheck/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MonitoredChainAll(goCtx context.Context, req *types.QueryAllMonitoredChainRequest) (*types.QueryAllMonitoredChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var MonitoredChains []types.MonitoredChain
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	MonitoredChainStore := prefix.NewStore(store, types.KeyPrefix(types.MonitoredChainKeyPrefix))

	pageRes, err := query.Paginate(MonitoredChainStore, req.Pagination, func(key []byte, value []byte) error {
		var MonitoredChain types.MonitoredChain
		if err := k.cdc.Unmarshal(value, &MonitoredChain); err != nil {
			return err
		}

		MonitoredChains = append(MonitoredChains, MonitoredChain)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMonitoredChainResponse{MonitoredChain: MonitoredChains, Pagination: pageRes}, nil
}

func (k Keeper) MonitoredChain(goCtx context.Context, req *types.QueryGetMonitoredChainRequest) (*types.QueryGetMonitoredChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetMonitoredChain(
		ctx,
		req.ChainId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMonitoredChainResponse{MonitoredChain: val}, nil
}
