package bn

import "time"

type optFunc func(c *clientOpts)

type clientOpts struct {
	timeout  time.Duration
	host     string
	username string
	password string
	cache    bool
}

func WithTimeout(seconds time.Duration) optFunc {
	return func(c *clientOpts) {
		c.timeout = seconds
	}
}

func WithCache() optFunc {
	return func(c *clientOpts) {
		c.cache = true
	}
}

func WithHost(host string) optFunc {
	return func(c *clientOpts) {
		c.host = host
	}
}

func WithCreds(username, password string) optFunc {
	return func(c *clientOpts) {
		c.username = username
		c.password = password
	}
}
