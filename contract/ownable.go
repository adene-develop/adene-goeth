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

type OwnableEvents interface {
	OwnershipTransferred(previousOwner common.Address, newOwner common.Address)
}

func ParseOwnableEvents(filterChanges []*eth.FilterChange, events OwnableEvents) error {

}
