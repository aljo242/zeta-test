package keeper

import (
	"sync/x/sync/types"
)

var _ types.QueryServer = Keeper{}
