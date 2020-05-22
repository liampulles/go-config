package config

import (
	"os"
)

// EnvSource uses os.Env to implement Source
type EnvSource struct{}

// Compile check to make sure EnvSource implements Source
var _ Source = &EnvSource{}

// NewEnvSource is a constructor
func NewEnvSource() *EnvSource {
	return &EnvSource{}
}

// GetString implements the Source interface
func (es *EnvSource) GetString(property string) (string, error) {
	if property == "" {
		return "", ErrEmptyProperty
	}
	value, found := os.LookupEnv(property)
	if !found {
		return "", newErrPropertyNotSet(property)
	}
	return value, nil
}
