package ethereumsupport

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Provide struct {
	Client *ethclient.Client
}

// Faucet data
type Faucet struct {
	Provider       *Provide
	SendPool       []common.Address `json:"send_pool"`
	SendPrivateKey []string         `json:"send_privatekey"`
	Nonce          []uint64         `json:"nonce"`
	Receiver       common.Address   `json:"receiver"`
	Tx             string           `json:"transaction"`
	Gaslimit       uint64           `json:"gas_limit"`
	GasPrice       *big.Int         `json:"gas_price"`
	Amount         *big.Int         `json:"amount"`
}

// Originization data
type Defi struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Address     string             `json:"address"`
	Symbol      string             `json:"symbol"`
	Url         string             `json:"url"`
	Description string             `json:"description"`
	Chain       string             `json:"chain"`
	Category    string             `json:"category"`
	Chains      []string           `json:"chains"`
	Twitter     string             `json:"twitter"`
	ChainTvls   map[string]float64 `json:"chainTvls"`
	Change1h    float64            `json:"change_1h"`
	Change1d    float64            `json:"change_1d"`
	Change7d    float64            `json:"change_7h"`
	Staking     float64            `json:"staking"`
	Fdv         float64            `json:"fdv"`
	Mcap        float64            `json:"mcap"`
}
