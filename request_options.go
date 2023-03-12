package jumpserver

import (
	"context"
	"net/http"
)

type RequestOption interface {
	apply(r *http.Request) error
}
type requestOptionFunc func(*http.Request) error

func (fn requestOptionFunc) apply(r *http.Request) error {
	return fn(r)
}

// WithContext
func WithContext(ctx context.Context) RequestOption {
	return requestOptionFunc(func(r *http.Request) error {
		r = r.WithContext(ctx)
		return nil
	})
}

// WithHeader
func WithHeader(key, val string) RequestOption {
	return requestOptionFunc(func(r *http.Request) error {
		r.Header.Set(key, val)
		return nil
	})
}

// WithHeaders
func WithHeaders(headers map[string]string) RequestOption {
	return requestOptionFunc(func(r *http.Request) error {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return nil
	})
}

// WithUserAgent
func WithUserAgent(agent string) RequestOption {
	return requestOptionFunc(func(r *http.Request) error {
		return nil
	})
}
