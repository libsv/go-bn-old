package bn

import "time"

type optFunc func(c *client)

func WithTimeout(seconds time.Duration) optFunc {
	return func(c *client) {
		c.c.Timeout = seconds * time.Second
	}
}
