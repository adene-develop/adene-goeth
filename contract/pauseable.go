package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
)

type Pauseable interface {
	// Paused returns true if the contract is paused, and false otherwise.
	Paused(ctx context.Context) (bool, error)
}

type PauseableEvents interface {
	// Paused emitted when the pause is triggered by `account`.
	Paused(account common.Address)

	// Unpaused emitted when the pause is lifted by `account`.
	Unpaused(account common.Address)
}

func NewPauseable(client *eth.Client, address common.Address) Pauseable {
	return &PauseableContract{
		client:  client,
		address: address,
	}
}

type PauseableContract struct {
	client  *eth.Client
	address common.Address
}

func (p *PauseableContract) Paused(ctx context.Context) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func ParsePauseableEvents(filterChanges []*eth.FilterChange, events PauseableEvents) error {
	return nil
}
