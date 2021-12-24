package contract

import (
	"context"
	"fmt"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const ERC721ABIString = `[
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name_",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "symbol_",
          "type": "string"
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
          "name": "approved",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "tokenId",
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
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "operator",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bool",
          "name": "approved",
          "type": "bool"
        }
      ],
      "name": "ApprovalForAll",
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
          "indexed": true,
          "internalType": "uint256",
          "name": "tokenId",
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
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "owner",
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
      "inputs": [
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "getApproved",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
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
          "name": "operator",
          "type": "address"
        }
      ],
      "name": "isApprovedForAll",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
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
      "inputs": [
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "ownerOf",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "safeTransferFrom",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        },
        {
          "internalType": "bytes",
          "name": "_data",
          "type": "bytes"
        }
      ],
      "name": "safeTransferFrom",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "operator",
          "type": "address"
        },
        {
          "internalType": "bool",
          "name": "approved",
          "type": "bool"
        }
      ],
      "name": "setApprovalForAll",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes4",
          "name": "interfaceId",
          "type": "bytes4"
        }
      ],
      "name": "supportsInterface",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
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
      "inputs": [
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "tokenURI",
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
      "inputs": [
        {
          "internalType": "address",
          "name": "from",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "transferFrom",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

var ERC721ABI abi.ABI

func init() {
	a, err := abi.JSON(strings.NewReader(ERC721ABIString))
	if err != nil {
		panic(err)
	}
	ERC721ABI = a
}

// ERC721 Standard view functions
type ERC721 interface {
	Contract
	// Name returns the token collection name.
	Name(ctx context.Context) (string, error)

	// Symbol returns the token collection symbol.
	Symbol(ctx context.Context) (string, error)

	// TokenURI returns the Uniform Resource Identifier (URI) for `tokenId` token.
	TokenURI(ctx context.Context, tokenID string) (string, error)

	// BalanceOf returns the number of tokens in ``owner``'s account.
	BalanceOf(ctx context.Context, owner common.Address) (balance int64, err error)

	// OwnerOf returns the owner of the `tokenId` token.
	OwnerOf(ctx context.Context, tokenID int64) (owner common.Address, err error)

	// GetApproved returns the account approved for `tokenId` token.
	GetApproved(ctx context.Context, tokenID int64) (operator common.Address, err error)

	// IsApprovedForAll returns if the `operator` is allowed to manage all of the assets of `owner`.
	IsApprovedForAll(ctx context.Context, owner common.Address, operator common.Address) (bool, error)
}

func NewERC721(client *eth.Client, address common.Address) ERC721 {
	return &ERC721Contract{
		address: address,
		client:  client,
	}
}

type ERC721Contract struct {
	address common.Address
	client  *eth.Client
}

func (e *ERC721Contract) Address() common.Address {
	return e.address
}

func (e *ERC721Contract) Client() *eth.Client {
	return e.client
}

func (e *ERC721Contract) Name(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ERC721Contract) Symbol(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ERC721Contract) TokenURI(ctx context.Context, tokenID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e *ERC721Contract) BalanceOf(ctx context.Context, owner common.Address) (balance int64, err error) {
	res, err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, "balanceOf", owner)
	if err != nil {
		return 0, errors.Wrap(err, "ERC721Contract BalanceOf call view error")
	}

	i, ok := new(big.Int).SetString(fmt.Sprintf("%x", res), 16)
	if !ok {
		return 0, errors.New("ERC721Contract BalanceOf parse result error")
	}
	return i.Int64(), nil
}

func (e *ERC721Contract) OwnerOf(ctx context.Context, tokenID int64) (owner common.Address, err error) {
	res, err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, "ownerOf", big.NewInt(tokenID))
	if err != nil {
		return common.Address{}, errors.Wrap(err, "ERC721Contract OwnerOf call view error")
	}
	return common.BytesToAddress(res), nil
}

func (e *ERC721Contract) GetApproved(ctx context.Context, tokenID int64) (operator common.Address, err error) {
	//TODO implement me
	panic("implement me")
}

func (e *ERC721Contract) IsApprovedForAll(ctx context.Context, owner common.Address, operator common.Address) (bool, error) {
	//TODO implement me
	panic("implement me")
}

type ERC721Events interface {
	// Transfer emitted when `tokenId` token is transferred from `from` to `to`.
	Transfer(from common.Address, to common.Address, tokenID int64)

	// Approval emitted when `owner` enables `approved` to manage the `tokenId` token.
	Approval(owner common.Address, approved common.Address, tokenID int64)

	// ApprovalForAll emitted when `owner` enables or disables (`approved`) `operator` to manage all of its assets.
	ApprovalForAll(owner common.Address, operator common.Address, approved bool)
}

func ParseERC721Events(filterChanges []*eth.FilterChange, events ERC721Events) error {

}
