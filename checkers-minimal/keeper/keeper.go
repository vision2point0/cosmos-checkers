package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/alice/checkers"
	modulev1 "github.com/alice/checkers/api/module/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
)

// Define a custom type for StoreKey
type CheckersStoreKey struct{}

// Implement the Name method to satisfy the StoreKey interface
func (CheckersStoreKey) Name() string {
	return "checkers"
}

// Implement the String method to satisfy the StoreKey interface
func (CheckersStoreKey) String() string {
	return "checkers"
}

// Keeper manages the state of the checkers module
type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec
	authority    string
	Schema       collections.Schema
	Params       collections.Item[checkers.Params]
	StoreKey     CheckersStoreKey // Use the custom StoreKey type here
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, checkers.ParamsKey, "params", codec.CollValue[checkers.Params](cdc)),
		StoreKey:     CheckersStoreKey{}, // Initialize the custom StoreKey
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// CreateGame stores game data in the state
func (k Keeper) CreateGame(ctx types.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error) {
	store := ctx.KVStore(k.StoreKey) // Use the custom StoreKey

	// Create a unique game ID
	gameID := fmt.Sprintf("game_%d", time.Now().UnixNano())

	// Create the game data struct
	gameData := checkers.GameData{
		ID:        gameID,
		Player1:   req.Player1,
		Player2:   req.Player2,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	// Store the game data using MustMarshal
	dataBytes := k.cdc.MustMarshal(&gameData) // Marshal the game data
	store.Set([]byte(gameID), dataBytes)      // Store the marshaled data

	// Return the response
	return &modulev1.ResCheckersTorram{
		Success: true,
		Message: "Game created successfully",
	}, nil
}

// MsgServer interface for handling Checkers messages
type MsgServer interface {
	CheckersCreateGm(ctx context.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error)
}

// msgServer struct implementing the MsgServer interface
type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl creates a new MsgServer implementation
func NewMsgServerImpl(k Keeper) MsgServer {
	return &msgServer{keeper: k}
}

// Implement the CheckersCreateGm method
func (m msgServer) CheckersCreateGm(ctx context.Context, req *modulev1.ReqCheckersTorram) (*modulev1.ResCheckersTorram, error) {
	// Convert context to SDK context if needed
	sdkCtx := types.UnwrapSDKContext(ctx)

	// Call the keeper's CreateGame method
	return m.keeper.CreateGame(sdkCtx, req) // Directly pass req without dereferencing
}
