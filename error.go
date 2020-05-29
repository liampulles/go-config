package config

import (
	"fmt"
)

// ErrEmptyProperty is returned when a config function is given an empty string
// property
var ErrEmptyProperty = fmt.Errorf("cannot find a value for an empty property")

// ErrPropertyNotSet is returned when a property has no corresponding value
type ErrPropertyNotSet struct {
	Property string
}

// Compile check to make sure ErrPropertyNotSet implements error
var _ error = &ErrPropertyNotSet{}

func newErrPropertyNotSet(property string) *ErrPropertyNotSet {
	return &ErrPropertyNotSet{
		Property: property,
	}
}

// Error implements the error interface
func (e *ErrPropertyNotSet) Error() string {
	return fmt.Sprintf("%s property is not set", e.Property)
}

// ErrValueFormat is returned when the value found for a property cannot be
// converted to the desired format
type ErrValueFormat struct {
	Property          string
	ValueString       string
	DesiredFormatDesc string
}

// Compile check to make sure ErrValueFormat implements error
var _ error = &ErrValueFormat{}

// IntDesiredFormat is a ErrValueFormat.DesiredFormatDesc variant
const IntDesiredFormat = "int"

// NewErrValueFormat constructor
func newErrValueFormat(property string, valueString string, desiredFormatDesc string) *ErrValueFormat {
	return &ErrValueFormat{
		Property:          property,
		ValueString:       valueString,
		DesiredFormatDesc: desiredFormatDesc,
	}
}

// Error implements the error interface
func (e *ErrValueFormat) Error() string {
	return fmt.Sprintf("value of %s property can not be converted to %s (is %s)", e.Property, e.DesiredFormatDesc, e.ValueString)
}

// ErrUnknownPropertyType is returned when a property is not known
type ErrUnknownPropertyType struct {
	Type string
}

// NewErrUnknownPropertyType is a constructor
func NewErrUnknownPropertyType(typ interface{}) *ErrUnknownPropertyType {
	return &ErrUnknownPropertyType{
		Type: fmt.Sprintf("%T", typ),
	}
}

// Error implements the error interface
func (e *ErrUnknownPropertyType) Error() string {
	return fmt.Sprintf("unknown property type %s", e.Type)
}

// Compile check to make sure ErrUnknownPropertyType implements error
var _ error = &ErrUnknownPropertyType{}
