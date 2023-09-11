package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"sync/testutil/sample"
)

func TestMsgCreateHeader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateHeader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateHeader{
				Admin: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateHeader{
				Admin: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateHeader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateHeader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateHeader{
				Admin: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateHeader{
				Admin: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteHeader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteHeader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteHeader{
				Admin: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteHeader{
				Admin: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
