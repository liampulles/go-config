package config

import (
	"os"
	"testing"
)

func TestEnvSource_GetString_GivenInput_ShouldReturnAsExpected(t *testing.T) {
	testGetStringPassingCases(envSourceConstructor, t)
}

func TestEnvSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetStringFailingCases(envSourceConstructor, t)
}

func envSourceConstructor(property string, val *string) Source {
	if val == nil {
		os.Unsetenv(property)
	} else {
		os.Setenv(property, *val)
	}
	return NewEnvSource()
}
