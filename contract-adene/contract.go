package adene

import (
	"context"
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/ethereum/go-ethereum/common"
)

type ADENE interface {
	contract.ERC20
	contract.Ownable

	ReflectionFromToken(ctx context.Context, tAmount int64, deductTransferFee bool) (int64, error)
	TokenFromReflection(ctx context.Context, rAmount int64) (int64, error)
	IsExcludedFromFee(ctx context.Context, account common.Address) (bool, error)
}
