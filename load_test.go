package config

import (
	"testing"
)

func TestLoadProperties_GivenEmptyProperties_ShouldNotFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
	}))

	// Exercise SUT
	err := LoadProperties(sourceFixture)

	// Verify Results
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
}

func TestLoadProperties_GivenUnsetStringPropertyForRequiredMapping_ShouldFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
	}))
	actual := "prev"
	nonExistantProperty := StrProp("DOES NOT EXIST", &actual, true)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err == nil {
		t.Errorf("Expected error but none returned")
	}
	if actual != "prev" {
		t.Errorf("Expected actual to not be mapped, but is %s", actual)
	}
}

func TestLoadProperties_GivenUnsetIntPropertyForRequiredMapping_ShouldFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "1",
	}))
	actual := -1
	nonExistantProperty := IntProp("DOES NOT EXIST", &actual, true)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err == nil {
		t.Errorf("Expected error but none returned")
	}
	if actual != -1 {
		t.Errorf("Expected actual to not be mapped, but is %d", actual)
	}
}

func TestLoadProperties_GivenUnsetStringPropertyForNonRequiredMapping_ShouldPass(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
	}))
	actual := "prev"
	nonExistantProperty := StrProp("DOES NOT EXIST", &actual, false)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
	if actual != "prev" {
		t.Errorf("Expected actual to not be mapped, but is %s", actual)
	}
}

func TestLoadProperties_GivenUnsetIntPropertyForNonRequiredMapping_ShouldPass(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
	}))
	actual := -1
	nonExistantProperty := IntProp("DOES NOT EXIST", &actual, false)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
	if actual != -1 {
		t.Errorf("Expected actual to not be mapped, but is %d", actual)
	}
}

func TestLoadProperties_GivenEmptyStringProperty_ShouldFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
	}))
	actual := ""
	nonExistantProperty := StrProp("", &actual, false)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err == nil {
		t.Errorf("Expected error but none returned")
	}
	if actual != "" {
		t.Errorf("Expected actual to not be mapped, but is %s", actual)
	}
}

func TestLoadProperties_GivenEmptyIntProperty_ShouldFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "1",
	}))
	actual := -1
	nonExistantProperty := IntProp("", &actual, false)

	// Exercise SUT
	err := LoadProperties(sourceFixture, nonExistantProperty)

	// Verify Results
	if err == nil {
		t.Errorf("Expected error but none returned")
	}
	if actual != -1 {
		t.Errorf("Expected actual to not be mapped, but is %d", actual)
	}
}

func TestLoadProperties_GivenUnknownPropertyType_ShouldFail(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "1",
	}))
	expectedErr := "unknown property type *config.unknownProperty"

	// Exercise SUT
	err := LoadProperties(sourceFixture, &unknownProperty{})

	// Verify Results
	if err == nil {
		t.Errorf("Expected error but none returned")
	} else if err.Error() != expectedErr {
		t.Errorf("Unexpected error message\nActual: %s\nExpected: %s", err.Error(), expectedErr)
	}
}

func TestLoadProperties_GivenSetValidProperties_ShouldNotFailAndShouldMap(t *testing.T) {
	// Setup fixture
	sourceFixture := NewTypedSource(MapSource(map[string]string{
		"PROPERTY1": "VALUE1",
		"PROPERTY2": "VALUE2",
		"PROPERTY3": "3",
	}))
	actualStr := ""
	strProp := StrProp("PROPERTY2", &actualStr, false)
	actualInt := -1
	intProp := IntProp("PROPERTY3", &actualInt, false)

	// Exercise SUT
	err := LoadProperties(sourceFixture, strProp, intProp)

	// Verify Results
	if err != nil {
		t.Errorf("Unexpected error: %#v", err)
	}
	if actualStr != "VALUE2" {
		t.Errorf("Unexpected strProp\nActual: %s\nExpected: %s", actualStr, "VALUE2")
	}
	if actualInt != 3 {
		t.Errorf("Unexpected intProp\nActual: %d\nExpected: %d", actualInt, 3)
	}
}

type unknownProperty struct {
}

func (up *unknownProperty) required() bool {
	return false
}
