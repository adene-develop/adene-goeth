package adene

import (
	"context"
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
)

type ADENE interface {
	contract.ERC20
	contract.Ownable

	ReflectionFromToken(ctx context.Context, tAmount int64, deductTransferFee bool) (int64, error)
	TokenFromReflection(ctx context.Context, rAmount int64) (int64, error)
	IsExcludedFromFee(ctx context.Context, account common.Address) (bool, error)
}

func NewADENEContract(client *eth.Client, address common.Address) ADENE {
	return &ADENEContract{
		ERC20:   contract.NewERC20(client, address),
		Ownable: contract.NewOwnable(client, address),
		client:  client,
		address: address,
	}
}

type ADENEContract struct {
	contract.ERC20
	contract.Ownable
	client  *eth.Client
	address common.Address
}

func (A *ADENEContract) ReflectionFromToken(ctx context.Context, tAmount int64, deductTransferFee bool) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (A *ADENEContract) TokenFromReflection(ctx context.Context, rAmount int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (A *ADENEContract) IsExcludedFromFee(ctx context.Context, account common.Address) (bool, error) {
	//TODO implement me
	panic("implement me")
}
