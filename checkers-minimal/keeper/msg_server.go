package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
)

type msgServer struct {
	k Keeper
}

var _ checkers.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) checkers.MsgServer {
	return &msgServer{k: keeper}
}

// CreateGame defines the handler for the MsgCreateGame message.
func (ms msgServer) CheckersCreateGm(ctx context.Context, msg *checkers.CheckersTorram) (*checkers.ResCheckersTorram, error) {
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	// Get current time in Unix format
	currentTime := time.Now().Unix()

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		Black:     msg.Black,
		Red:       msg.Red,
		StartTime: currentTime,         // Set current time as start time
		EndTime:   currentTime + 30*60, // Or set to a duration from now if needed
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResCheckersTorram{}, nil
}
