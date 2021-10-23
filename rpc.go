package bn

import (
	"net/http"
	"reflect"
	"time"

	"github.com/libsv/go-bn/config"
	"github.com/libsv/go-bn/service"
)

type NodeClient interface {
	BlockChainClient
	ControlClient
	MiningClient
	NetworkClient
	TransactionClient
}

type positionalOptionalArgs interface {
	Args() []interface{}
}

type client struct {
	rpc       service.RPC
	isMainnet bool
}

func NewNodeClient(oo ...optFunc) NodeClient {
	opts := &clientOpts{
		timeout:   30 * time.Second,
		host:      "http://localhost:8332",
		username:  "bitcoin",
		password:  "bitcoin",
		isMainnet: false,
	}
	for _, o := range oo {
		o(opts)
	}
	return &client{
		rpc: service.NewRPC(&config.RPC{
			Username: opts.username,
			Password: opts.password,
			Host:     opts.host,
		}, &http.Client{Timeout: opts.timeout}),
	}
}

func (c *client) argsFor(p positionalOptionalArgs, args ...interface{}) []interface{} {
	if reflect.ValueOf(p).IsNil() {
		return args
	}

	return append(args, p.Args()...)
}
