package ethereumsupport

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func GetBalance(p *Provide, account common.Address) (float64, error) {
	balance, err := p.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	ethValue, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return ethValue, nil
}
