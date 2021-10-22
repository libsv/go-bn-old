package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/libsv/go-bn"
)

func main() {
	c := bn.NewNodeClient(
		bn.WithHost("http://localhost:18332"),
		bn.WithCreds("bitcoin", "bitcoin"),
	)

	if err := c.Ping(context.TODO()); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj, err := c.SetExcessiveBlock(context.TODO(), 400000000)
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
