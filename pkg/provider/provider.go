package provider

import (
	"context"

	"github.com/spf13/cast"
)

// Provider defines the interface to a key/value provider (e.g. etcd)
type Provider interface {
	ListKeysWithContext(context.Context, string, bool) ([]*KVPair, error)
}

// AbstractProvider is the base provider with all providers inherit from
type AbstractProvider struct {
	Enable bool
}

// KVPair represents a { Key, Value } tuple
type KVPair struct {
	Key   string
	Value interface{}
}

// Get ...
func (k KVPair) Get() interface{} {
	return k.Value
}

// String ...
func (k KVPair) GetString() string {
	return cast.ToString(k.Get())
}

// Bool ...
func (k KVPair) GetBool() bool {
	return cast.ToBool(k.Get())
}
