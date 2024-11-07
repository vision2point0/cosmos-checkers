package keeper

import (
	"context"

	modulev1 "github.com/alice/checkers/api/module/v1"
	"github.com/cosmos/cosmos-sdk/types"
)

// CheckersCreateGm handles the RPC request
func (k Keeper) CheckersCreateGm(ctx context.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error) {
	// Convert context.Context to types.Context
	sdkCtx := types.UnwrapSDKContext(ctx)

	// Call the CreateGame method with the converted context
	return k.CreateGame(sdkCtx, req) // Pass req directly without dereferencing
}
