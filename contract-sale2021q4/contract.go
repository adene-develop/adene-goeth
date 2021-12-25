package sale2021q4

import (
	"context"
	"fmt"
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

const ContractName = "SALE2021Q4"

type Box struct {
	Price       int64
	TotalSupply int64
	Stock       int64
}

type BoxLevel int

const (
	BoxLevelCommon BoxLevel = iota
	BoxLevelRare
	BoxLevelLegendary
)

type SALE2021Q4 interface {
	contract.Contract
	contract.Pauseable
	contract.Ownable

	// Info trả về thông tin của contract
	Info(ctx context.Context) (tokenAddress common.Address, icon721Address common.Address,
		commonBox Box, rareBox Box, legendaryBox Box)

	// InfoWallet trả về thông tin của user
	InfoWallet(ctx context.Context, user common.Address) (commonBought int, rareBought int, legendaryBought int)

	// BoxLevelOf trả về level của 1 box
	BoxLevelOf(ctx context.Context, tokenID int64) (level BoxLevel, err error)
}

func NewSale2021Q4Contract(client *eth.Client, address common.Address, abi abi.ABI) SALE2021Q4 {
	return &sale2021q4{
		address: address,
		client:  client,
		abi:     abi,
	}
}

type sale2021q4 struct {
	contract.Pauseable
	contract.Ownable
	address common.Address
	client  *eth.Client
	abi     abi.ABI
}

func (s *sale2021q4) Address() common.Address {
	return s.address
}

func (s *sale2021q4) Client() *eth.Client {
	return s.client
}

func (s *sale2021q4) Info(ctx context.Context) (tokenAddress common.Address, icon721Address common.Address, commonBox Box, rareBox Box, legendaryBox Box) {
	//TODO implement me
	panic("implement me")
}

func (s *sale2021q4) InfoWallet(ctx context.Context, user common.Address) (commonBought int, rareBought int, legendaryBought int) {
	//TODO implement me
	panic("implement me")
}

func (s *sale2021q4) BoxLevelOf(ctx context.Context, tokenID int64) (level BoxLevel, err error) {
	res, err := s.client.CallContractViewFunction(ctx, s.abi, s.address, "boxLevelOf", big.NewInt(tokenID))
	if err != nil {
		return 0, errors.Wrap(err, "sale2021q4 BoxLevelOf call view error")
	}

	i, ok := new(big.Int).SetString(fmt.Sprintf("%x", res), 16)
	if !ok {
		return 0, errors.New("ERC721Contract BalanceOf parse result error")
	}
	return BoxLevel(i.Int64()), nil
}
