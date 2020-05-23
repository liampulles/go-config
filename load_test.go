package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadStringProperties_GivenSetProperties_ShouldMapWithoutError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		setPropertiesFixture map[string]string
		mappingFixture       map[string]*string
		expected             map[string]*string
	}{
		// Empty case
		{
			map[string]string{},
			map[string]*string{},
			map[string]*string{},
		},
		// Empty mappings in -> Empty mappings out
		{
			map[string]string{"PROPERTY": "VALUE"},
			map[string]*string{},
			map[string]*string{},
		},
		// One property match
		{
			map[string]string{
				"PROPERTY1": "VALUE1",
				"PROPERTY2": "VALUE2",
			},
			map[string]*string{"PROPERTY1": strPtr("")},
			map[string]*string{"PROPERTY1": strPtr("VALUE1")},
		},
		// A few property matches
		{
			map[string]string{
				"PROPERTY1": "VALUE1",
				"PROPERTY2": "VALUE2",
				"PROPERTY3": "VALUE3",
			},
			map[string]*string{
				"PROPERTY1": strPtr(""),
				"PROPERTY2": strPtr(""),
			},
			map[string]*string{
				"PROPERTY1": strPtr("VALUE1"),
				"PROPERTY2": strPtr("VALUE2"),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Setup fixture
			sourceFixture := MapSource(test.setPropertiesFixture)

			// Exercise SUT
			err := LoadStringProperties(sourceFixture, test.mappingFixture)

			// Verify result
			if err != nil {
				t.Errorf("Unexpected error:\n%#v", err)
			}
			if !reflect.DeepEqual(test.mappingFixture, test.expected) {
				t.Errorf("Unexpected Result.\nActual: %#v\nExpected: %#v", test.mappingFixture, test.expected)
			}
		})
	}
}

func TestLoadStringProperties_GivenUnsetOrBadProperties_ShouldReturnError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		setPropertiesFixture map[string]string
		mappingFixture       map[string]*string
		expectedErr          error
	}{
		// Empty property
		{
			map[string]string{"PROPERTY": "VALUE"},
			map[string]*string{"": strPtr("something")},
			ErrEmptyProperty,
		},
		// Unset property
		{
			map[string]string{
				"PROPERTY1": "VALUE1",
				"PROPERTY2": "VALUE2",
			},
			map[string]*string{
				"PROPERTY1":             strPtr(""),
				"NON_EXISTENT_PROPERTY": strPtr(""),
			},
			&ErrPropertyNotSet{Property: "NON_EXISTENT_PROPERTY"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Setup fixture
			sourceFixture := MapSource(test.setPropertiesFixture)

			// Exercise SUT
			err := LoadStringProperties(sourceFixture, test.mappingFixture)

			// Verify result
			if err == nil {
				t.Errorf("Expected error, but none was returned")
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("Unexpected Result.\nActual: %v\nExpected: %v", err, test.expectedErr)
			}
		})
	}
}

func TestLoadIntProperties_GivenSetProperties_ShouldMapWithoutError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		setPropertiesFixture map[string]string
		mappingFixture       map[string]*int
		expected             map[string]*int
	}{
		// Empty case
		{
			map[string]string{},
			map[string]*int{},
			map[string]*int{},
		},
		// Empty mappings in -> Empty mappings out
		{
			map[string]string{"PROPERTY": "1"},
			map[string]*int{},
			map[string]*int{},
		},
		// One property match
		{
			map[string]string{
				"PROPERTY1": "1",
				"PROPERTY2": "2",
			},
			map[string]*int{"PROPERTY1": intPtr(0)},
			map[string]*int{"PROPERTY1": intPtr(1)},
		},
		// A few property matches
		{
			map[string]string{
				"PROPERTY1": "1",
				"PROPERTY2": "2",
				"PROPERTY3": "3",
			},
			map[string]*int{
				"PROPERTY1": intPtr(0),
				"PROPERTY2": intPtr(0),
			},
			map[string]*int{
				"PROPERTY1": intPtr(1),
				"PROPERTY2": intPtr(2),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Setup fixture
			sourceFixture := NewTypedSource(MapSource(test.setPropertiesFixture))

			// Exercise SUT
			err := LoadIntProperties(sourceFixture, test.mappingFixture)

			// Verify result
			if err != nil {
				t.Errorf("Unexpected error:\n%#v", err)
			}
			if !reflect.DeepEqual(test.mappingFixture, test.expected) {
				t.Errorf("Unexpected Result.\nActual: %#v\nExpected: %#v", test.mappingFixture, test.expected)
			}
		})
	}
}

func TestLoadIntProperties_GivenUnsetOrBadProperties_ShouldReturnError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		setPropertiesFixture map[string]string
		mappingFixture       map[string]*int
		expectedErr          error
	}{
		// Empty property
		{
			map[string]string{"PROPERTY": "1"},
			map[string]*int{"": intPtr(0)},
			ErrEmptyProperty,
		},
		// Unset property
		{
			map[string]string{
				"PROPERTY1": "1",
				"PROPERTY2": "2",
			},
			map[string]*int{
				"PROPERTY1":             intPtr(0),
				"NON_EXISTENT_PROPERTY": intPtr(0),
			},
			&ErrPropertyNotSet{Property: "NON_EXISTENT_PROPERTY"},
		},
		// Value in wrong format
		{
			map[string]string{"PROPERTY": "VALUE"},
			map[string]*int{"PROPERTY": intPtr(0)},
			&ErrValueFormat{Property: "PROPERTY", ValueString: "VALUE", DesiredFormatDesc: IntDesiredFormat},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Setup fixture
			sourceFixture := NewTypedSource(MapSource(test.setPropertiesFixture))

			// Exercise SUT
			err := LoadIntProperties(sourceFixture, test.mappingFixture)

			// Verify result
			if err == nil {
				t.Errorf("Expected error, but none was returned")
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("Unexpected Result.\nActual: %v\nExpected: %v", err, test.expectedErr)
			}
		})
	}
}
