package graphql

import (
	"context"

	"github.com/pkg/errors"
)

// errInappropriateContext is returned when the context has not been
// configured with graphql.NewContext.
var errInappropriateContext = errors.New("inappropriate context")

// contextKey provides unique keys for context values.
type contextKey string

// clientContextKey is the context value key for the Client.
var clientContextKey = contextKey("graphql client context")

// fromContext gets the client from the specified
// Context.
func fromContext(ctx context.Context) (*client, error) {
	c, ok := ctx.Value(clientContextKey).(*client)
	if !ok {
		return nil, errInappropriateContext
	}
	return c, nil
}

// NewContext makes a new context.Context that enables requests.
func NewContext(parent context.Context, endpoint string) context.Context {
	client := &client{
		endpoint: endpoint,
	}
	return context.WithValue(parent, clientContextKey, client)
}