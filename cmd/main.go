package main

import (
	"context"
	"fmt"

	"github.com/libsv/go-bn"
)

func main() {
	c := bn.NewClient("http://localhost:18332")

	hash, err := c.BestBlockHash(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(hash)
}
