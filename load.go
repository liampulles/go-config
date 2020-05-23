package config

// LoadStringProperties will try load properties from source into variable reference
func LoadStringProperties(source Source, propertiesToReferences map[string]*string) error {
	for property, ref := range propertiesToReferences {
		value, err := source.GetString(property)
		if err != nil {
			return err
		}
		*ref = value
	}
	return nil
}

// LoadIntProperties will try load properties from source into variable reference
func LoadIntProperties(typedSource *TypedSource, propertiesToReferences map[string]*int) error {
	for property, ref := range propertiesToReferences {
		value, err := typedSource.GetInt(property)
		if err != nil {
			return err
		}
		*ref = value
	}
	return nil
}
