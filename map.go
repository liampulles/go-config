package config

// MapSource resolves config properties to a map
type MapSource map[string]string

// Compile check to make sure MapSource implements Source
var _ Source = MapSource{}

// GetString implements the Source interface
func (m MapSource) GetString(property string) (string, error) {
	if property == "" {
		return "", ErrEmptyProperty
	}
	value, found := m[property]
	if !found {
		return "", newErrPropertyNotSet(property)
	}
	return value, nil
}
