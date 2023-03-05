package options

import "time"

type Connection struct {
	addr    string
	timeout time.Duration
	cache   bool
}

const (
	defaultTimeout = 10
	defaultCache   = true
)

type options struct {
	timeout time.Duration
	cache   bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(options *options)

func (f optionFunc) apply(options *options) {
	f(options)
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(opts *options) {
		opts.timeout = timeout
	})
}

func WithCache(cache bool) Option {
	return optionFunc(func(options *options) {
		options.cache = cache
	})
}

func NewConnection(addr string, opts ...Option) (*Connection, error) {
	opt := &options{
		timeout: defaultTimeout,
		cache:   defaultCache,
	}
	for _, o := range opts {
		o.apply(opt)
	}

	return &Connection{
		addr:    addr,
		timeout: opt.timeout,
		cache:   opt.cache,
	}, nil
}
