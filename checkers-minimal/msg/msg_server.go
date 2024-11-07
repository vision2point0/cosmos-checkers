package msg

import (
	"context"

	modulev1 "github.com/alice/checkers/api/module/v1"
	"github.com/alice/checkers/keeper"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

// MsgServer defines the gRPC service for the checkers module.
type MsgServer interface {
	CheckersCreateGm(ctx context.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error)
}

// msgServer implements MsgServer and embeds the UnimplementedCheckersTorramServer
type msgServer struct {
	keeper keeper.Keeper
}

// NewMsgServerImpl creates a new MsgServer implementation.
func NewMsgServerImpl(k keeper.Keeper) *msgServer {
	return &msgServer{keeper: k}
}

// CheckersCreateGm handles the RPC request to create a checkers game.
func (m msgServer) CheckersCreateGm(ctx context.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error) {
	// Convert context to SDK context if needed
	sdkCtx := types.UnwrapSDKContext(ctx)

	// Call the keeper's CreateGame method
	return m.keeper.CreateGame(sdkCtx, req) // Directly pass req as it is of type *modulev1.ReqCheckersTorram
}

// RegisterMsgServer registers the message server for the checkers module
func RegisterMsgServer(server grpc.ServiceRegistrar, srv MsgServer) {
	modulev1.RegisterCheckersTorramServer(server, srv) // Ensure this function is generated correctly
}

// RegisterInterfaces registers the interfaces for this module
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("alice.checkers.module.v1.CheckersTorram", (*MsgServer)(nil))
}
