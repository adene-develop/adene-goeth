package icon721

import (
	"context"
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

const ContractName = "ICON721"

type ICON721 interface {
	contract.ERC721Enumerable
	contract.Ownable

	// InfoWallet returns allocated and remaining allocations number of `user` address
	// `allocated` is the total number of ICON721 that `user` can be mint
	// `remainingAllocation` is the remaining number of ICON721 that `user` can be mint
	InfoWallet(ctx context.Context, user common.Address) (allocated int64, remainingAllocation int64, err error)
}

func NewIcon721Contract(client *eth.Client, address common.Address) ICON721 {
	return &icon721{
		ERC721Enumerable: contract.NewERC721Enumerable(client, address),
		address:          address,
		client:           client,
	}
}

type icon721 struct {
	contract.ERC721Enumerable
	contract.Ownable
	address common.Address
	client  *eth.Client
}

func (i *icon721) InfoWallet(ctx context.Context, user common.Address) (allocated int64, remainingAllocation int64, err error) {
	var result struct {
		Allocated           uint16
		RemainingAllocation uint16
	}

	err = i.client.CallContractViewFunction(ctx, ABI, i.address, &result, "infoWallet", user)
	if err != nil {
		return 0, 0, errors.Wrap(err, "icon721 call view `infoWallet` error")
	}

	return int64(result.Allocated), int64(result.RemainingAllocation), nil
}
