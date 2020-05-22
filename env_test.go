package config_test

import (
	"os"
	"testing"

	"github.com/liampulles/go-config"
)

func TestEnvSource_GetString_GivenInput_ShouldReturnAsExpected(t *testing.T) {
	testGetString_PassingCases(envSourceConstructor, t)
}

func TestEnvSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetString_FailingCases(envSourceConstructor, t)
}

func envSourceConstructor(property string, val *string) config.Source {
	if val == nil {
		os.Unsetenv(property)
	} else {
		os.Setenv(property, *val)
	}
	return config.NewEnvSource()
}
