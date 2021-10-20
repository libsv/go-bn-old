package bn

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-bn/models"
	"github.com/pkg/errors"
)

// Default jsonrpc fields.
const (
	ID      = "gobn"
	JSONRpc = "1.0"
)

type Client interface {
	Info(ctx context.Context) (*models.Info, error)
	BestBlockHash(ctx context.Context) (string, error)
	Block(ctx context.Context, hash string, opts models.BlockOptions) (bc.Block, error)
	BlockByHeight(ctx context.Context, height int, opts models.BlockOptions)
}

type client struct {
	host string
	c    *http.Client
}

func NewClient(host string, oo ...optFunc) Client {
	c := &client{
		c:    &http.Client{Timeout: 30 * time.Second},
		host: host,
	}
	for _, o := range oo {
		o(c)
	}
	return c
}

func (c *client) Info(ctx context.Context) (*models.Info, error) {
	var resp models.Info
	return &resp, c.performRPC(ctx, "getinfo", &resp)
}

func (c *client) BestBlockHash(ctx context.Context) (string, error) {
	var resp string
	return resp, c.performRPC(ctx, "getbestblockhash", &resp)
}

func (c *client) BlockHash(ctx context.Context, hash string, opts models.BlockOptions) {
	panic("not implemented") // TODO: Implement
}

func (c *client) BlockByHeight(ctx context.Context, height int, opts models.BlockOptions) {
	panic("not implemented") // TODO: Implement
}

func (c *client) performRPC(ctx context.Context, method string, out interface{}, params ...interface{}) error {
	data, err := json.Marshal(&models.Request{
		ID:      ID,
		JSONRpc: JSONRpc,
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.host,
		bytes.NewReader(data),
	)
	if err != nil {
		return err
	}
	req.SetBasicAuth("bitcoin", "bitcoin")
	req.Header.Add("Content-Type", "text/plain")

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	reply := models.Response{
		Result: out,
	}
	if err = json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return errors.Wrapf(err, "error decoding response")
	}

	if resp.StatusCode != http.StatusOK {
		return reply.Error
	}

	return nil
}
