package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"sync/x/sync/types"
)

func TestHeaderMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	admin := types.DefaultAdmin
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateHeader(ctx, &types.MsgCreateHeader{Admin: admin})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestHeaderMsgServerUpdate(t *testing.T) {
	admin := types.DefaultAdmin

	tests := []struct {
		desc    string
		request *types.MsgUpdateHeader
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateHeader{Admin: admin},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHeader{Admin: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateHeader{Admin: admin, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateHeader(ctx, &types.MsgCreateHeader{Admin: admin})
			require.NoError(t, err)

			_, err = srv.UpdateHeader(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHeaderMsgServerDelete(t *testing.T) {
	admin := types.DefaultAdmin

	tests := []struct {
		desc    string
		request *types.MsgDeleteHeader
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteHeader{Admin: admin},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteHeader{Admin: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteHeader{Admin: admin, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateHeader(ctx, &types.MsgCreateHeader{Admin: admin})
			require.NoError(t, err)
			_, err = srv.DeleteHeader(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
