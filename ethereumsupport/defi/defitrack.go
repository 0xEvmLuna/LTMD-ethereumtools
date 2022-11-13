package defi

import (
	eth "discordbot/ethereumsupport"
	"encoding/json"
	"fmt"
	"math"
	"sync"
)

type DefiTokens struct {
	Change1d []eth.Defi
	Change1h []eth.Defi
	Change7d []eth.Defi
}

func DefiTrack() *DefiTokens {

	resp := eth.Get("https://api.llama.fi/protocols", false, nil)

	var result []eth.Defi
	if err := json.NewDecoder(resp.Response.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	ch := make(chan eth.Defi, 5000)
	//stop := make(chan bool)

	for _, v := range result {
		ch <- v
	}

	var callback DefiTokens
	var wg sync.WaitGroup

	wg.Add(1)
	go priceChange(ch, &callback, &wg)

	close(ch)
	wg.Wait()

	return &callback
}

func priceChange(tokens <-chan eth.Defi, callback *DefiTokens, wg *sync.WaitGroup) {
	decimal := func(value float64) float64 {
		return math.Trunc(value*1e2+0.5) * 1e-1
	}

	defer wg.Done()
	for t := range tokens {
		if decimal(t.Change1d) >= decimal(float64(10)) || decimal(t.Change1h) <= decimal(float64(-10)) {
			callback.Change1d = append(callback.Change1d, t)
		}

		if decimal(t.Change1h) >= decimal(float64(10)) || decimal(t.Change1h) <= decimal(float64(-10)) {
			callback.Change1h = append(callback.Change1h, t)
		}

		if decimal(t.Change7d) >= decimal(float64(10)) || decimal(t.Change1h) <= decimal(float64(-10)) {
			callback.Change7d = append(callback.Change7d, t)
		}

	}
}
