package config

import (
	"fmt"
	"testing"
)

func TestTypedSource_GetString_GivenFixture_ShouldReturnAsExpected(t *testing.T) {
	testGetStringPassingCases(typedSourceConstructorAsSource, t)
}

func TestTypedSource_GetString_GivenFaultyData_ShouldReturnError(t *testing.T) {
	testGetStringFailingCases(typedSourceConstructorAsSource, t)
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
			strPtr("1"),
			1,
		},

		// Property with negative set value -> Return set value
		{
			"property",
			strPtr("-1"),
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
			strPtr("1"),
			ErrEmptyProperty,
		},

		// Unset property
		{
			"property",
			nil,
			&ErrPropertyNotSet{
				Property: "property",
			},
		},

		// Non-int property
		{
			"property",
			strPtr("not an int"),
			&ErrValueFormat{
				Property:          "property",
				ValueString:       "not an int",
				DesiredFormatDesc: IntDesiredFormat,
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

func typedSourceConstructorAsSource(property string, val *string) Source {
	return typedSourceConstructor(property, val)
}

func typedSourceConstructor(property string, val *string) *TypedSource {
	return NewTypedSource(mapSourceConstructor(property, val))
}
