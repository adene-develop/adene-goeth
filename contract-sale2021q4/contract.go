package sale2021q4

import (
	"context"
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

const ContractName = "SALE2021Q4"

type Box struct {
	Price       *big.Int
	TotalSupply *big.Int
	Stock       *big.Int
}

type BoxLevel int

const (
	BoxLevelCommon BoxLevel = iota
	BoxLevelRare
	BoxLevelLegendary
)

type SALE2021Q4 interface {
	contract.Contract
	contract.Pausable
	contract.Ownable

	// Info trả về thông tin của contract
	Info(ctx context.Context) (tokenAddress common.Address, icon721Address common.Address,
		commonBox Box, rareBox Box, legendaryBox Box, err error)

	// InfoWallet trả về thông tin của user
	InfoWallet(ctx context.Context, user common.Address) (commonBought int, rareBought int, legendaryBought int, err error)

	// BoxLevelOf trả về level của 1 box
	BoxLevelOf(ctx context.Context, tokenID int64) (level BoxLevel, err error)
}

func NewSale2021Q4Contract(client *eth.Client, address common.Address) SALE2021Q4 {
	return &sale2021q4{
		address: address,
		client:  client,
	}
}

type sale2021q4 struct {
	contract.Pausable
	contract.Ownable
	address common.Address
	client  *eth.Client
}

func (s *sale2021q4) Address() common.Address {
	return s.address
}

func (s *sale2021q4) Client() *eth.Client {
	return s.client
}

func (s *sale2021q4) Info(ctx context.Context) (tokenAddress common.Address, icon721Address common.Address, commonBox Box, rareBox Box, legendaryBox Box, err error) {
	var result struct {
		Token     common.Address
		Icon721   common.Address
		Common    Box
		Rare      Box
		Legendary Box
	}
	err = s.client.CallContractViewFunction(ctx, ABI, s.address, &result, "info")
	if err != nil {
		return [20]byte{}, [20]byte{}, Box{}, Box{}, Box{}, errors.Wrap(err, "sale2021q4 call view `info` error")
	}

	return result.Token, result.Icon721, result.Common, result.Rare, result.Legendary, nil
}

func (s *sale2021q4) InfoWallet(ctx context.Context, user common.Address) (commonBought int, rareBought int, legendaryBought int, err error) {
	var result struct {
		Common    uint16
		Rare      uint16
		Legendary uint16
	}

	err = s.client.CallContractViewFunction(ctx, ABI, s.address, &result, "infoWallet", user)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "sale2021q4 call view `infoWallet` error")
	}

	return int(result.Common), int(result.Rare), int(result.Legendary), nil
}

func (s *sale2021q4) BoxLevelOf(ctx context.Context, tokenID int64) (BoxLevel, error) {
	var result struct {
		BoxLevel *big.Int
	}

	err := s.client.CallContractViewFunction(ctx, ABI, s.address, &result, "boxLevelOf", big.NewInt(tokenID))
	if err != nil {
		return 0, errors.Wrap(err, "sale2021q4 BoxLevelOf call view error")
	}

	return BoxLevel(result.BoxLevel.Int64()), nil
}
