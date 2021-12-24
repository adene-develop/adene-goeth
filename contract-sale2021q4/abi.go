package sale2021q4

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

const SALE2021Q4ABIJSON = `[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "enum SALE2021Q4.BoxLevel",
          "name": "level",
          "type": "uint8"
        },
        {
          "indexed": false,
          "internalType": "uint16",
          "name": "amount",
          "type": "uint16"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "startTokenId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "toTokenId",
          "type": "uint256"
        }
      ],
      "name": "Bought",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "address",
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "Paused",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "Unpaused",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "_icon721Contract",
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
      "inputs": [],
      "name": "_paymentContract",
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
          "internalType": "uint256",
          "name": "tokenId",
          "type": "uint256"
        }
      ],
      "name": "boxLevelOf",
      "outputs": [
        {
          "internalType": "enum SALE2021Q4.BoxLevel",
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
          "internalType": "enum SALE2021Q4.BoxLevel",
          "name": "level",
          "type": "uint8"
        },
        {
          "internalType": "uint16",
          "name": "amount",
          "type": "uint16"
        }
      ],
      "name": "buy",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "info",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "infoWallet",
      "outputs": [
        {
          "internalType": "uint16",
          "name": "",
          "type": "uint16"
        },
        {
          "internalType": "uint16",
          "name": "",
          "type": "uint16"
        },
        {
          "internalType": "uint16",
          "name": "",
          "type": "uint16"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "paymentContract",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "icon721Contract",
          "type": "address"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "commonBox",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "rareBox",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "uint256",
              "name": "price",
              "type": "uint256"
            },
            {
              "internalType": "uint16",
              "name": "totalSupply",
              "type": "uint16"
            },
            {
              "internalType": "uint16",
              "name": "stock",
              "type": "uint16"
            }
          ],
          "internalType": "struct SALE2021Q4.Box",
          "name": "legendaryBox",
          "type": "tuple"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "owner",
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
      "inputs": [],
      "name": "pause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "paused",
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
      "name": "renounceOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        }
      ],
      "name": "transferToken",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "unpause",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

var SALE2021Q4ABI abi.ABI

func init() {
	var err error
	SALE2021Q4ABI, err = abi.JSON(strings.NewReader(SALE2021Q4ABIJSON))
	if err != nil {
		panic(err)
	}
}
