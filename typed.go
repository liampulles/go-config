package config

import "strconv"

// TypedSource provides typed configuration values from a normal Source
type TypedSource struct {
	backingSource Source
}

// Compile check to make sure TypedSource implements Source
var _ Source = &TypedSource{}

// NewTypedSource is a constructor
func NewTypedSource(backingSource Source) *TypedSource {
	return &TypedSource{
		backingSource: backingSource,
	}
}

// GetString implements the Source interface
func (ts *TypedSource) GetString(property string) (string, error) {
	return ts.backingSource.GetString(property)
}

// GetInt tries to parse an int from the property, else returns the default.
func (ts *TypedSource) GetInt(property string) (int, error) {
	value, err := ts.backingSource.GetString(property)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, newErrValueFormat(property, value, IntDesiredFormat)
	}
	return i, nil
}
