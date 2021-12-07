package metadata

import (
	"fmt"
	"reflect"

	"github.com/volatiletech/null/v8"
)

type Map map[string]null.String

type Metadata struct {
	Labels      Map `json:"labels"`
	Annotations Map `json:"annotations"`
}

func Get(annotationsSlice, labelsSlice interface{}) (Metadata, error) {
	var metadata Metadata

	annotationsMap, err := getKeyValue(annotationsSlice, "Key", "Value")
	if err != nil {
		return metadata, err
	}
	metadata.Annotations = annotationsMap

	labelsMap, err := getKeyValue(labelsSlice, "KeyName", "Value")
	if err != nil {
		return metadata, err
	}
	metadata.Labels = labelsMap

	return metadata, nil
}

func getKeyValue(slice interface{}, keyField, valueField string) (keyValueMap Map, err error) {
	keyValueMap = make(Map)
	if slice == nil {
		return keyValueMap, nil
	}

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("could not extract key and value, %s", rec)
		}
	}()

	switch reflect.TypeOf(slice).Kind() { //nolint:exhaustive // Reflection to check if type is slice
	case reflect.Slice:
		slice := reflect.ValueOf(slice)
		for i := 0; i < slice.Len(); i++ {
			elem := slice.Index(i)

			keyElem := elem.Elem().FieldByName(keyField)
			key, valid := keyElem.Interface().(null.String)
			if !valid || key.IsZero() {
				return nil, fmt.Errorf("key is not a valid string")
			}

			valueElem := elem.Elem().FieldByName(valueField)
			value, valid := valueElem.Interface().(null.String)
			if !valid {
				return nil, fmt.Errorf("value is not type of null.String")
			}
			keyValueMap[key.String] = value
		}
		return keyValueMap, err
	default:
		return nil, fmt.Errorf("metadata component is not a slice")
	}
}
