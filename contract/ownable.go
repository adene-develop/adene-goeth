package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type Ownable interface {
	// Owner returns the address of the current owner.
	Owner(ctx context.Context) (common.Address, error)
}

func NewOwnable(client *eth.Client, address common.Address) Ownable {
	return &OwnableContract{
		client:  client,
		address: address,
	}
}

type OwnableContract struct {
	client  *eth.Client
	address common.Address
}

func (o *OwnableContract) Owner(ctx context.Context) (common.Address, error) {
	var result struct {
		Owner common.Address
	}
	if err := o.client.CallContractViewFunction(ctx, OwnableABI, o.address, &result, "owner"); err != nil {
		return common.Address{}, errors.Wrap(err, "OwnableContract call `owner` error")
	}

	return result.Owner, nil
}

type OwnableEvents interface {
	OwnershipTransferred(previousOwner common.Address, newOwner common.Address)
}

func ParseOwnableEvents(filterChanges []*eth.FilterChange, events OwnableEvents) error {
	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case OwnableABI.Events["OwnershipTransferred"].ID:
			if filterChanges[i].Topics == nil || len(filterChanges[i].Topics) < 3 {
				return errors.New("invalid topics")
			}
			previousOwner := common.BytesToAddress(filterChanges[i].Topics[1].Bytes())
			newOwner := common.BytesToAddress(filterChanges[i].Topics[2].Bytes())
			events.OwnershipTransferred(previousOwner, newOwner)
		default:
		}
	}
	return nil
}
