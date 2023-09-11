package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sync module sentinel errors
var (
	ErrInvalidHash = sdkerrors.Register(ModuleName, 0, "invalid hash")
)
