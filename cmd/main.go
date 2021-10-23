package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/libsv/go-bn"
	"github.com/libsv/go-bt/v2"
)

func main() {
	c := bn.NewNodeClient(
		bn.WithHost("http://localhost:18332"),
		bn.WithCreds("bitcoin", "bitcoin"),
	)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tx := bt.NewTx()
			tx.AddP2PKHOutputFromAddress("mpzLdVLZhbRXxYpaT8YcHntWb2tyPJvUnz", 123456700)

			obj, err := c.FundRawTransaction(context.TODO(), tx, nil)
			if err != nil {
				panic(err)
			}
			bb, err := json.MarshalIndent(obj, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bb))
		}()
	}

	wg.Wait()
}
