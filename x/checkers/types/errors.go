package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1104, "game by id not found: %s")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1105, "message creator is not a player: %s")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1106, "player tried to play out of turn: %s")
	ErrWrongMove        = sdkerrors.Register(ModuleName, 1107, "wrong move")

	ErrRedAlreadyPlayed   = sdkerrors.Register(ModuleName, 1108, "red player has already played")
	ErrBlackAlreadyPlayed = sdkerrors.Register(ModuleName, 1109, "black player has already played")
)
