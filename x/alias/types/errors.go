package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/alias module sentinel errors
var (
	ErrMaxEntries    = sdkerrors.Register(ModuleName, 1101, "max entries reached")
	ErrAliasNotFree  = sdkerrors.Register(ModuleName, 1102, "alias not free")
	ErrAliasNotFound = sdkerrors.Register(ModuleName, 1103, "alias not found")
	ErrInvalidAlias  = sdkerrors.Register(ModuleName, 1104, "alias not valid")
)
