package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync/x/sync/types"
)

func (k Keeper) HeaderAll(goCtx context.Context, req *types.QueryAllHeaderRequest) (*types.QueryAllHeaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var headers []types.Header
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	headerStore := prefix.NewStore(store, types.KeyPrefix(types.HeaderKey))

	pageRes, err := query.Paginate(headerStore, req.Pagination, func(key []byte, value []byte) error {
		var header types.Header
		if err := k.cdc.Unmarshal(value, &header); err != nil {
			return err
		}

		headers = append(headers, header)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHeaderResponse{Header: headers, Pagination: pageRes}, nil
}

func (k Keeper) Header(goCtx context.Context, req *types.QueryGetHeaderRequest) (*types.QueryGetHeaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	header, found := k.GetHeader(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHeaderResponse{Header: header}, nil
}
