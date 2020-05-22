package config_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/go-config"
)

func TestTypedSource_GetString_GivenFixture_ShouldReturnAsExpected(t *testing.T) {
	testGetString_PassingCases(typedSourceConstructorAsSource, t)
}

func TestTypedSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetString_FailingCases(typedSourceConstructorAsSource, t)
}

func TestTypedSource_GetInt_GivenFixture_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		propertyFixture string
		setValueFixture *string
		expected        int
	}{
		// Property with set value -> Return set value
		{
			"property",
			ptr("1"),
			1,
		},

		// Property with negative set value -> Return set value
		{
			"property",
			ptr("-1"),
			-1,
		},
	}

	for _, test := range tests {
		ptrValue := "(unset)"
		if test.setValueFixture != nil {
			ptrValue = *test.setValueFixture
		}
		t.Run(fmt.Sprintf("(%s) -> %s -> %d",
			test.propertyFixture, ptrValue, test.expected), func(t *testing.T) {

			// Setup fixture
			sourceFixture := typedSourceConstructor(test.propertyFixture, test.setValueFixture)

			// Exercise SUT
			actual, err := sourceFixture.GetInt(test.propertyFixture)

			// Verify result
			if err != nil {
				t.Errorf("Encountered error\n%v", err)
			}
			if actual != test.expected {
				t.Errorf("Unexpected Result.\nActual: %d\nExpected: %d", actual, test.expected)
			}
		})
	}
}

func TestTypedSource_GetInt_GivenFaultyData_ShouldReturnError(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		propertyFixture string
		setValueFixture *string
		expectedErr     error
	}{
		// Empty property
		{
			"",
			ptr("1"),
			config.ErrEmptyProperty,
		},

		// Unset property
		{
			"property",
			nil,
			&config.ErrPropertyNotSet{
				Property: "property",
			},
		},

		// Non-int property
		{
			"property",
			ptr("not an int"),
			&config.ErrValueFormat{
				Property:          "property",
				ValueString:       "not an int",
				DesiredFormatDesc: "int",
			},
		},
	}

	for _, test := range tests {
		ptrValue := "(unset)"
		if test.setValueFixture != nil {
			ptrValue = *test.setValueFixture
		}
		t.Run(fmt.Sprintf("(%s) -> %s -> %v",
			test.propertyFixture, ptrValue, test.expectedErr), func(t *testing.T) {

			// Setup fixture
			sourceFixture := typedSourceConstructor(test.propertyFixture, test.setValueFixture)

			// Exercise SUT
			actual, err := sourceFixture.GetInt(test.propertyFixture)

			// Verify result
			if err == nil {
				t.Errorf("Expected error, but none was returned")
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("Unexpected Result.\nActual: %v\nExpected: %v", err, test.expectedErr)
			}
			if actual != 0 {
				t.Errorf("Unexpected Result.\nActual: %d\nExpected: %d", actual, 0)
			}
		})
	}
}

func typedSourceConstructorAsSource(property string, val *string) config.Source {
	return typedSourceConstructor(property, val)
}

func typedSourceConstructor(property string, val *string) *config.TypedSource {
	return config.NewTypedSource(mapSourceConstructor(property, val))
}
