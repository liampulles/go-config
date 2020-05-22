package config_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/go-config"
)

type sourceConstructor func(property string, value *string) config.Source

func testGetString_PassingCases(constructor sourceConstructor, t *testing.T) {
	// Setup fixture
	var tests = []struct {
		propertyFixture string
		setValueFixture *string
		expected        string
	}{
		// Property with set value -> Return set value
		{
			"property",
			ptr("value"),
			"value",
		},
	}

	for _, test := range tests {
		ptrValue := "(unset)"
		if test.setValueFixture != nil {
			ptrValue = *test.setValueFixture
		}
		t.Run(fmt.Sprintf("(%s) -> %s -> %s",
			test.propertyFixture, ptrValue, test.expected), func(t *testing.T) {

			// Setup fixture
			sourceFixture := constructor(test.propertyFixture, test.setValueFixture)

			// Exercise SUT
			actual, err := sourceFixture.GetString(test.propertyFixture)

			// Verify result
			if err != nil {
				t.Errorf("Encountered error\n%v", err)
			}
			if actual != test.expected {
				t.Errorf("Unexpected Result.\nActual: %s\nExpected: %s", actual, test.expected)
			}
		})
	}
}

func testGetString_FailingCases(constructor sourceConstructor, t *testing.T) {
	// Setup fixture
	var tests = []struct {
		propertyFixture string
		setValueFixture *string
		expectedErr     error
	}{
		// Empty property
		{
			"",
			ptr("value"),
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
	}

	for _, test := range tests {
		ptrValue := "(unset)"
		if test.setValueFixture != nil {
			ptrValue = *test.setValueFixture
		}
		t.Run(fmt.Sprintf("(%s) -> %s -> %v",
			test.propertyFixture, ptrValue, test.expectedErr), func(t *testing.T) {

			// Setup fixture
			sourceFixture := constructor(test.propertyFixture, test.setValueFixture)

			// Exercise SUT
			actual, err := sourceFixture.GetString(test.propertyFixture)

			// Verify result
			if err == nil {
				t.Errorf("Expected error, but none was returned")
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("Unexpected Result.\nActual: %v\nExpected: %v", err, test.expectedErr)
			}
			if actual != "" {
				t.Errorf("Unexpected Result.\nActual: %s\nExpected: %s", actual, "")
			}
		})
	}
}

func ptr(val string) *string {
	return &val
}
