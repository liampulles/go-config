package config_test

import (
	"testing"

	"github.com/liampulles/go-config"
)

func TestMapSource_GetString_GivenInput_ShouldReturnAsExpected(t *testing.T) {
	testGetString_PassingCases(mapSourceConstructor, t)
}

func TestMapSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetString_FailingCases(mapSourceConstructor, t)
}

func mapSourceConstructor(property string, val *string) config.Source {
	if val == nil {
		return config.MapSource(nil)
	}
	return config.MapSource(map[string]string{property: *val})
}
