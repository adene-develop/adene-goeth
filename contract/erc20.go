package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

type ERC20 interface {
	// Name returns the name of the token.
	Name(ctx context.Context) (string, error)

	// Symbol returns the symbol of the token.
	Symbol(ctx context.Context) (string, error)

	// Decimals returns the decimals places of the token.
	Decimals(ctx context.Context) (uint8, error)

	// TotalSupply returns total supply of this tokens
	TotalSupply(ctx context.Context) (*big.Int, error)

	// BalanceOf returns the amount of tokens owned by `account`.
	BalanceOf(ctx context.Context, account common.Address) (*big.Int, error)

	// Allowance returns the remaining number of tokens that `spender` will be
	// allowed to spend on behalf of `owner` through {transferFrom}. This is
	// zero by default.
	Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error)
}

func NewERC20(client *eth.Client, address common.Address) ERC20 {
	return &ERC20Contract{
		address: address,
		client:  client,
	}
}

type ERC20Contract struct {
	address common.Address
	client  *eth.Client
}

func (e *ERC20Contract) Name(ctx context.Context) (string, error) {
	var result struct {
		Name string
	}

	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "name"); err != nil {
		return "", errors.Wrap(err, "ERC20Contract call view `name` error")
	}

	return result.Name, nil
}

func (e *ERC20Contract) Symbol(ctx context.Context) (string, error) {
	var result struct {
		Symbol string
	}

	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "symbol"); err != nil {
		return "", errors.Wrap(err, "ERC20Contract call view `symbol` error")
	}

	return result.Symbol, nil
}

func (e *ERC20Contract) Decimals(ctx context.Context) (uint8, error) {
	var result struct {
		Decimals uint8
	}

	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "decimals"); err != nil {
		return 0, errors.Wrap(err, "ERC20Contract call view `decimals` error")
	}
	return result.Decimals, nil
}

func (e *ERC20Contract) TotalSupply(ctx context.Context) (*big.Int, error) {
	var result struct {
		TotalSupply *big.Int
	}

	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "totalSupply"); err != nil {
		return nil, errors.Wrap(err, "ERC20Contract call view `totalSupply` error")
	}

	return result.TotalSupply, nil
}

func (e *ERC20Contract) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {
	var result struct {
		Balance *big.Int
	}
	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "balanceOf", account); err != nil {
		return nil, errors.Wrap(err, "ERC20Contract call view `balanceOf` error")
	}
	return result.Balance, nil
}

func (e *ERC20Contract) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	var result struct {
		Allowance *big.Int
	}
	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "allowance", owner, spender); err != nil {
		return nil, errors.Wrap(err, "ERC20Contract call view `allowance` error")
	}
	return result.Allowance, nil
}

type ERC20Events interface {
	// Transfer emitted when `value` tokens are moved from one account (`from`) to
	// another (`to`).
	Transfer(from, to common.Address, value *big.Int)

	// Approval emitted when the allowance of a `spender` for an `owner` is set by
	// a call to {approve}. `value` is the new allowance.
	Approval(owner, spender common.Address, value *big.Int)
}

func ParseERC20Events(filterChanges []*eth.FilterChange, events ERC20Events) error {
	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case ERC20ABI.Events["Transfer"].ID:
			if err := parseERC20TransferEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC20Events")
			}
		case ERC20ABI.Events["Approval"].ID:
			if err := parseERC20ApprovalEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC20Events")
			}
		default:
		}
	}
	return nil
}

func parseERC20TransferEvent(change *eth.FilterChange, events ERC20Events) error {
	if change.Topics == nil || len(change.Topics) < 3 {
		return errors.New("invalid topics")
	}
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())

	var r struct {
		Value *big.Int
	}

	if err := ERC20ABI.UnpackIntoInterface(&r, "Transfer", change.Data); err != nil {
		return err
	}
	events.Transfer(from, to, r.Value)
	return nil
}

func parseERC20ApprovalEvent(change *eth.FilterChange, events ERC20Events) error {
	if change.Topics == nil || len(change.Topics) < 3 {
		return errors.New("invalid topics")
	}
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	var r struct {
		Value *big.Int
	}

	if err := ERC20ABI.UnpackIntoInterface(&r, "Approval", change.Data); err != nil {
		return err
	}
	events.Approval(from, to, r.Value)
	return nil
}
