package icon721

import (
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const ContractName = "ICON721"

type ICON721 interface {
	contract.ERC721Enumerable
	contract.Ownable

	// InfoWallet returns allocated and remaining allocations number of `user` address
	// `allocated` is the total number of ICON721 that `user` can be mint
	// `remainingAllocation` is the remaining number of ICON721 that `user` can be mint
	InfoWallet(user common.Address) (allocated int64, remainingAllocation int64, err error)
}

func NewIcon721Contract(client *eth.Client, address common.Address, abi abi.ABI) ICON721 {
	return &icon721{
		ERC721Enumerable: contract.NewERC721Enumerable(client, address),
		address:          address,
		client:           client,
		abi:              abi,
	}
}

type icon721 struct {
	contract.ERC721Enumerable
	contract.Ownable
	address common.Address
	client  *eth.Client
	abi     abi.ABI
}

func (i *icon721) InfoWallet(user common.Address) (allocated int64, remainingAllocation int64, err error) {
	//TODO implement me
	panic("implement me")
}
