package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// TestnetEndpoint là RPC URL của mạng testnet trên bsc
const TestnetEndpoint = "https://data-seed-prebsc-1-s1.binance.org:8545/"

// MainnetEndpoint là RPC URL của mạng mainnet trên bsc
const MainnetEndpoint = "https://bsc-dataseed.binance.org/"

type Client struct {
	eth *ethclient.Client
	rpc *rpc.Client
}

func NewClient(rpcEndpoint string) (*Client, error) {
	c := &Client{}
	var err error

	// connect to rpc rpcEndpoint
	c.eth, err = ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	c.rpc, err = rpc.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) EthClient() *ethclient.Client {
	return c.eth
}

func (c *Client) RpcClient() *rpc.Client {
	return c.rpc
}

// CallContractViewFunction gọi hàm view với tên hàm là `function` của contract với abi  là `abi`
// `args` là params truyền vào hàm abi.Value
func (c *Client) CallContractViewFunction(ctx context.Context, abi abi.ABI, contractAddress common.Address, function string, args ...interface{}) (result []byte, err error) {
	data, err := abi.Pack(function, args...)
	if err != nil {
		return nil, errors.Wrap(err, "client abi pack error")
	}

	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	res, err := c.eth.CallContract(ctx, callMsg, nil)
	if err != nil {
		return nil, errors.Wrap(err, "client call contract error")
	}

	return res, nil
}

func (c *Client) Close() {
	c.eth.Close()
	c.rpc.Close()
}
