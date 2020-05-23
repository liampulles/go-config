package config

import (
	"testing"
)

func TestMapSource_GetString_GivenInput_ShouldReturnAsExpected(t *testing.T) {
	testGetStringPassingCases(mapSourceConstructor, t)
}

func TestMapSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetStringFailingCases(mapSourceConstructor, t)
}

func mapSourceConstructor(property string, val *string) Source {
	if val == nil {
		return MapSource(nil)
	}
	return MapSource(map[string]string{property: *val})
}
