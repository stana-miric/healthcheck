package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"healthcheck/x/healthcheck/types"
)

func (k msgServer) CreateMonitoredChains(goCtx context.Context, msg *types.MsgCreateMonitoredChains) (*types.MsgCreateMonitoredChainsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetMonitoredChains(
		ctx,
		msg.ChainId,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var monitoredChains = types.MonitoredChains{
		Creator:      msg.Creator,
		ChainId:      msg.ChainId,
		ConnectionId: msg.ConnectionId,
	}

	k.SetMonitoredChains(
		ctx,
		monitoredChains,
	)
	return &types.MsgCreateMonitoredChainsResponse{}, nil
}

func (k msgServer) UpdateMonitoredChains(goCtx context.Context, msg *types.MsgUpdateMonitoredChains) (*types.MsgUpdateMonitoredChainsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMonitoredChains(
		ctx,
		msg.ChainId,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var monitoredChains = types.MonitoredChains{
		Creator:      msg.Creator,
		ChainId:      msg.ChainId,
		ConnectionId: msg.ConnectionId,
	}

	k.SetMonitoredChains(ctx, monitoredChains)

	return &types.MsgUpdateMonitoredChainsResponse{}, nil
}

func (k msgServer) DeleteMonitoredChains(goCtx context.Context, msg *types.MsgDeleteMonitoredChains) (*types.MsgDeleteMonitoredChainsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetMonitoredChains(
		ctx,
		msg.ChainId,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMonitoredChains(
		ctx,
		msg.ChainId,
	)

	return &types.MsgDeleteMonitoredChainsResponse{}, nil
}
