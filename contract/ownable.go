package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
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
	//TODO implement me
	panic("implement me")
}

type OwnableEvents interface {
	OwnershipTransferred(previousOwner common.Address, newOwner common.Address)
}

func ParseOwnableEvents(filterChanges []*eth.FilterChange, events OwnableEvents) error {
	return nil
}
