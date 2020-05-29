package config

type property interface {
	required() bool
}

type propertyBase struct {
	Property  string
	_Required bool
}

func (pb *propertyBase) required() bool {
	return pb._Required
}

// StringProperty holds a string mapping
type StringProperty struct {
	propertyBase
	mapTo *string
}

// StrProp constructs StringProperty
func StrProp(property string, mapTo *string, required bool) *StringProperty {
	return &StringProperty{
		propertyBase{
			property,
			required,
		},
		mapTo,
	}
}

// IntProperty holds an int mapping
type IntProperty struct {
	propertyBase
	mapTo *int
}

// IntProp constructs IntProperty
func IntProp(property string, mapTo *int, required bool) *IntProperty {
	return &IntProperty{
		propertyBase{
			property,
			required,
		},
		mapTo,
	}
}

// LoadProperties will try load properties from source into variable reference
func LoadProperties(typedSource *TypedSource, properties ...property) error {
	for _, property := range properties {
		switch v := property.(type) {
		case *StringProperty:
			if err := mapStringProperty(typedSource, v); err != nil {
				return err
			}
		case *IntProperty:
			if err := mapIntProperty(typedSource, v); err != nil {
				return err
			}
		default:
			return NewErrUnknownPropertyType(property)
		}
	}
	return nil
}

func mapStringProperty(typedSource *TypedSource, property *StringProperty) error {
	result, err := typedSource.GetString(property.Property)
	if err == nil {
		*property.mapTo = result
		return nil
	}
	if _, notPresent := err.(*ErrPropertyNotSet); notPresent && !property.required() {
		return nil
	}
	return err
}

func mapIntProperty(typedSource *TypedSource, property *IntProperty) error {
	result, err := typedSource.GetInt(property.Property)
	if err == nil {
		*property.mapTo = result
		return nil
	}
	if _, notPresent := err.(*ErrPropertyNotSet); notPresent && !property.required() {
		return nil
	}
	return err
}
