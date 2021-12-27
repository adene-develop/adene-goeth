package contract

import (
	"context"
	"fmt"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math"
	"math/big"
	"strings"
)

const ERC20ABIString = `[
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "symbol",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "initialSupply",
          "type": "uint256"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Approval",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Transfer",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        }
      ],
      "name": "allowance",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "decimals",
      "outputs": [
        {
          "internalType": "uint8",
          "name": "",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "subtractedValue",
          "type": "uint256"
        }
      ],
      "name": "decreaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "spender",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "addedValue",
          "type": "uint256"
        }
      ],
      "name": "increaseAllowance",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "name",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "symbol",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "totalSupply",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transfer",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "sender",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "transferFrom",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

var ERC20ABI abi.ABI

func init() {
	a, err := abi.JSON(strings.NewReader(ERC721ABIString))
	if err != nil {
		panic(err)
	}
	ERC721ABI = a
}

type ERC20 interface {
	// Name returns the name of the token.
	Name(ctx context.Context) (string, error)

	// Symbol returns the symbol of the token.
	Symbol(ctx context.Context) (string, error)

	// Decimals returns the decimals places of the token.
	Decimals(ctx context.Context) (uint8, error)

	// TotalSupply returns total supply of this tokens
	TotalSupply(ctx context.Context) (float64, error)

	// BalanceOf returns the amount of tokens owned by `account`.
	BalanceOf(ctx context.Context, account common.Address) (float64, error)

	// Allowance returns the remaining number of tokens that `spender` will be
	// allowed to spend on behalf of `owner` through {transferFrom}. This is
	// zero by default.
	Allowance(ctx context.Context, owner, spender common.Address) (float64, error)
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

func (e *ERC20Contract) TotalSupply(ctx context.Context) (float64, error) {
	var result struct {
		TotalSupply *big.Int
	}

	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "totalSupply"); err != nil {
		return 0, errors.Wrap(err, "ERC20Contract call view `totalSupply` error")
	}

	f, err := e.realNumberOfTokens(ctx, result.TotalSupply)
	if err != nil {
		return 0, errors.Wrap(err, "ERC20Contract TotalSupply get real total supply error")
	}
	return f, nil
}

func (e *ERC20Contract) BalanceOf(ctx context.Context, account common.Address) (float64, error) {
	var result struct {
		Balance *big.Int
	}
	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "balanceOf", account); err != nil {
		return 0, errors.Wrap(err, "ERC20Contract call view `balanceOf` error")
	}
	f, err := e.realNumberOfTokens(ctx, result.Balance)
	if err != nil {
		return 0, errors.Wrap(err, "ERC20Contract BalanceOf get real total supply error")
	}
	return f, nil
}

func (e *ERC20Contract) Allowance(ctx context.Context, owner, spender common.Address) (float64, error) {
	var result struct {
		Allowance *big.Int
	}
	if err := e.client.CallContractViewFunction(ctx, ERC20ABI, e.address, &result, "allowance", owner, spender); err != nil {
		return 0, errors.Wrap(err, "ERC20Contract call view `allowance` error")
	}
	f, err := e.realNumberOfTokens(ctx, result.Allowance)
	if err != nil {
		return 0, errors.Wrap(err, "ERC20Contract Allowance get real total supply error")
	}
	return f, nil
}

// return the real number of given tokens = total tokens / pow(10, decimals)
func (e *ERC20Contract) realNumberOfTokens(ctx context.Context, tokens *big.Int) (float64, error) {
	decimals, err := e.Decimals(ctx)
	if err != nil {
		return 0, err
	}
	x := new(big.Float).SetInt(tokens)
	y := big.NewFloat(math.Pow10(int(decimals)))
	f, _ := new(big.Float).Quo(x, y).Float64()
	return f, nil
}

type ERC20Events interface {
	// Transfer emitted when `value` tokens are moved from one account (`from`) to
	// another (`to`).
	Transfer(from, to common.Address, value float64)

	// Approval emitted when the allowance of a `spender` for an `owner` is set by
	// a call to {approve}. `value` is the new allowance.
	Approval(owner, spender common.Address, value float64)
}

func ParseERC20Events(filterChanges []*eth.FilterChange, events ERC20Events) error {
	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case ERC20ABI.Events["Transfer"].ID:
			if err := parseTransferEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC20Events")
			}
		case ERC20ABI.Events["Approval"].ID:
			if err := parseApprovalEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC20Events")
			}
		default:
		}
	}
	return nil
}

func parseTransferEvent(change *eth.FilterChange, events ERC20Events) error {
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	f, ok := new(big.Float).SetString(fmt.Sprintf("%x", change.Topics[3].Bytes()))
	if !ok {
		return errors.New("parse transfer event error")
	}
	value, _ := f.Float64()
	events.Transfer(from, to, value)
	return nil
}

func parseApprovalEvent(change *eth.FilterChange, events ERC20Events) error {
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	f, ok := new(big.Float).SetString(fmt.Sprintf("%x", change.Topics[3].Bytes()))
	if !ok {
		return errors.New("parse approval event error")
	}
	value, _ := f.Float64()
	events.Approval(from, to, value)
	return nil
}
