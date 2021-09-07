package metadata_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/metadata"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

func TestMetadata(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		annotationsSlice, labelSlice interface{}
		expectedMetadata             Metadata
		err                          error
	}{
		"buildpack_metadata": {
			annotationsSlice: models.BuildAnnotationSlice{
				&models.BuildAnnotation{Key: null.StringFrom("key1"), Value: null.StringFrom("value1")},
				&models.BuildAnnotation{Key: null.StringFrom("key2"), Value: null.StringFrom("value2")},
			},
			labelSlice: models.BuildpackLabelSlice{
				&models.BuildpackLabel{KeyName: null.StringFrom("key-name1"), Value: null.StringFrom("value1")},
				&models.BuildpackLabel{KeyName: null.StringFrom("key-name2"), Value: null.StringFrom("value2")},
			},
			expectedMetadata: Metadata{
				Annotations: Map{"key1": null.StringFrom("value1"), "key2": null.StringFrom("value2")},
				Labels:      Map{"key-name1": null.StringFrom("value1"), "key-name2": null.StringFrom("value2")},
			},
			err: nil,
		},
		"generic_metadata": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.StringFrom("key"), null.StringFrom("value")}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{Annotations: Map{"key": null.StringFrom("value")}, Labels: Map{"keyName": null.StringFrom("value")}},
			err:              nil,
		},
		"no_slice": {
			annotationsSlice: "not-a-annotation-slice",
			labelSlice:       "not-a-label-slice",
			expectedMetadata: Metadata{},
			err:              fmt.Errorf("metadata component is not a slice"),
		},
		"invalid_value_string": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.StringFrom("key"), null.NewString("broken", false)}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{
				Annotations: Map{"key": null.NewString("broken", false)},
				Labels:      Map{"keyName": null.StringFrom("value")},
			},
			err: nil,
		},
		"invalid_key_string": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.NewString("invalid-key", false), null.StringFrom("value")}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{},
			err:              fmt.Errorf("key is not a valid string"),
		},
		"invalid_value": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.StringFrom("key"), null.StringFrom("value")}},
			labelSlice: []*struct {
				KeyName null.String
				Value   int
			}{{null.StringFrom("keyName"), 123}},
			expectedMetadata: Metadata{},
			err:              fmt.Errorf("value is not type of null.String"),
		},
		"nil_metadata": {
			annotationsSlice: nil, labelSlice: nil,
			expectedMetadata: Metadata{Labels: Map{}, Annotations: Map{}},
			err:              nil,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := Get(tc.annotationsSlice, tc.labelSlice)
			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedMetadata, got)
			}
		})
	}
}
