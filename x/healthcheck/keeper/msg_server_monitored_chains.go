package keeper

import (
	"context"

	"healthcheck/x/healthcheck/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMonitoredChain(goCtx context.Context, msg *types.MsgCreateMonitoredChain) (*types.MsgCreateMonitoredChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMonitoredChain(
		ctx,
		msg.ChainId,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var MonitoredChain = types.MonitoredChain{
		ChainId:      msg.ChainId,
		ConnectionId: msg.ConnectionId,
	}

	k.SetMonitoredChain(
		ctx,
		MonitoredChain,
	)
	return &types.MsgCreateMonitoredChainResponse{}, nil
}

func (k msgServer) UpdateMonitoredChain(goCtx context.Context, msg *types.MsgUpdateMonitoredChain) (*types.MsgUpdateMonitoredChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetMonitoredChain(
		ctx,
		msg.ChainId,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var MonitoredChain = types.MonitoredChain{
		//Creator:      msg.Creator,
		ChainId:      msg.ChainId,
		ConnectionId: msg.ConnectionId,
	}

	k.SetMonitoredChain(ctx, MonitoredChain)

	return &types.MsgUpdateMonitoredChainResponse{}, nil
}

func (k msgServer) DeleteMonitoredChain(goCtx context.Context, msg *types.MsgDeleteMonitoredChain) (*types.MsgDeleteMonitoredChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetMonitoredChain(
		ctx,
		msg.ChainId,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveMonitoredChain(
		ctx,
		msg.ChainId,
	)

	return &types.MsgDeleteMonitoredChainResponse{}, nil
}
