package types

import (
	"context"
)

// contextKey is used for all keys in go context
// it does not support built in types like "string" as key
type contextKey string

const requestIDKey contextKey = "RequestID"

// Context defines the basic go context with essential functions
type Context struct {
	context.Context
}

// GetReqID returns the request id if set or will return empty string
func (c Context) GetReqID() string {
	requestID, _ := c.Value(requestIDKey).(string)
	return requestID
}

// GetNewContext returns an empty context
func GetNewContext() Context {
	return Context{}
}

// NewContext creates a Context  given requestId
func NewContext(ctx context.Context, requestID string) Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return Context{context.WithValue(ctx, requestIDKey, requestID)}
}

// Set creates a new Context with the given data set
func (c Context) Set(key string, value interface{}) Context {
	return Context{context.WithValue(c, contextKey(key), value)}
}

// Get gets a value from the context or nil if it does not exist
func (c Context) Get(key string) interface{} {
	value := c.Value(contextKey(key))
	return value
}
