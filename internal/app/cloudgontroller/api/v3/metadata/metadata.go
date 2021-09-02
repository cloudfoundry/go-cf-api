package metadata

import (
	"fmt"
	"reflect"

	"github.com/volatiletech/null/v8"
)

type MetadataMap map[string]null.String

type Metadata struct {
	Labels      MetadataMap `json:"labels"`
	Annotations MetadataMap `json:"annotations"`
}

func GetMetadata(annotationsSlice, labelsSlice interface{}) (Metadata, error) {
	var metadata Metadata

	annotationsMap, err := getKeyAndValue(annotationsSlice, "Key", "Value")
	if err != nil {
		return metadata, err
	}
	metadata.Annotations = annotationsMap

	labelsMap, err := getKeyAndValue(labelsSlice, "KeyName", "Value")
	if err != nil {
		return metadata, err
	}
	metadata.Labels = labelsMap

	return metadata, nil
}

func getKeyAndValue(slice interface{}, keyField, valueField string) (keyValueMap MetadataMap, err error) {
	keyValueMap = make(MetadataMap)
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("could not extract key and value, %s", rec)
		}
	}()
	if slice == nil {
		return keyValueMap, nil
	}

	switch reflect.TypeOf(slice).Kind() { //nolint:exhaustive // Reflection to check if type is slice
	case reflect.Slice:
		slice := reflect.ValueOf(slice)
		for i := 0; i < slice.Len(); i++ {
			elem := slice.Index(i)

			keyElem := elem.Elem().FieldByName(keyField)
			key, ok := keyElem.Interface().(null.String)
			if !ok || key.IsZero() {
				return nil, fmt.Errorf("key is not a valid string")
			}

			valueElem := elem.Elem().FieldByName(valueField)
			value, ok := valueElem.Interface().(null.String)
			if !ok {
				return nil, fmt.Errorf("value is not type of null.String")
			}
			keyValueMap[key.String] = value
		}
		return keyValueMap, err
	default:
		return nil, fmt.Errorf("annotationSlice parameter is not a slice")
	}
}
