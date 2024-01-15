package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"healthcheck/x/healthcheck/types"
)

func (k Keeper) MonitoredChainsAll(goCtx context.Context, req *types.QueryAllMonitoredChainsRequest) (*types.QueryAllMonitoredChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var monitoredChainss []types.MonitoredChains
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	monitoredChainsStore := prefix.NewStore(store, types.KeyPrefix(types.MonitoredChainsKeyPrefix))

	pageRes, err := query.Paginate(monitoredChainsStore, req.Pagination, func(key []byte, value []byte) error {
		var monitoredChains types.MonitoredChains
		if err := k.cdc.Unmarshal(value, &monitoredChains); err != nil {
			return err
		}

		monitoredChainss = append(monitoredChainss, monitoredChains)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMonitoredChainsResponse{MonitoredChains: monitoredChainss, Pagination: pageRes}, nil
}

func (k Keeper) MonitoredChains(goCtx context.Context, req *types.QueryGetMonitoredChainsRequest) (*types.QueryGetMonitoredChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetMonitoredChains(
		ctx,
		req.ChainId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMonitoredChainsResponse{MonitoredChains: val}, nil
}
