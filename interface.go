package config

// Source provides configuration values
type Source interface {
	// Note: property must be in upper case format, like an environment variable.
	GetString(property string) (string, error)
}
