package server

import (
	"errors"

	"golang.org/x/net/context"
)

// Endpoint is the fundamental building block of package server.
// It represents a single RPC method.
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Middleware is a chainable behavior modifier.
type Middleware func(Endpoint) Endpoint

// ErrBadCast indicates a type error during decoding or encoding.
var ErrBadCast = errors.New("bad cast")

// ErrContextCanceled indicates a controlling context was canceled before the
// request could be served.
var ErrContextCanceled = errors.New("context was canceled")
