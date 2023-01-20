// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...any)
	// hooks to execute on mutations.
	hooks *hooks
	// interceptors to execute on queries.
	inters *inters
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Contact          []ent.Hook
		FinancialAccount []ent.Hook
		User             []ent.Hook
	}
	inters struct {
		Contact          []ent.Interceptor
		FinancialAccount []ent.Interceptor
		User             []ent.Interceptor
	}
)

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
