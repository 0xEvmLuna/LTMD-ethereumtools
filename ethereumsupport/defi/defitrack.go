package defi

import (
	eth "discordbot/ethereumsupport"
	"encoding/json"
	"fmt"
	"math"
)

type DefiToken struct {
	Change1d      chan eth.Defi
	Tvl           chan eth.Defi
	LiquidityData chan eth.Defi
}

func NewDefi() *DefiToken {
	return &DefiToken{
		Change1d:      make(chan eth.Defi),
		Tvl:           make(chan eth.Defi),
		LiquidityData: make(chan eth.Defi),
	}
}

func (dt *DefiToken) DefiTrack() {

	resp := eth.Get("https://api.llama.fi/protocols", false, nil)

	var result []eth.Defi
	if err := json.NewDecoder(resp.Response.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	for _, v := range result {
		go TvlFilter(dt.Tvl, v)
		go Liquidity(dt.LiquidityData, v)
		go PriceChange(dt.Change1d, v)
	}

}

func TvlFilter(ch chan<- eth.Defi, defi eth.Defi) {

}

func Liquidity(ch chan<- eth.Defi, defi eth.Defi) {

}

func PriceChange(ch chan<- eth.Defi, defi eth.Defi) {
	decimal := func(value float64) float64 {
		return math.Trunc(value*1e2+0.5) * 1e-1
	}

	if decimal(defi.Change1d) >= decimal(float64(30)) {
		ch <- eth.Defi{
			Id:       defi.Id,
			Name:     defi.Name,
			Symbol:   defi.Symbol,
			Change1d: defi.Change1d,
			Address:  defi.Address,
			Url:      defi.Url,
		}
	}

}
