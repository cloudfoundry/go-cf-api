package metadata_test

import (
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
		err                          bool
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
				Annotations: MetadataMap{"key1": null.StringFrom("value1"), "key2": null.StringFrom("value2")},
				Labels:      MetadataMap{"key-name1": null.StringFrom("value1"), "key-name2": null.StringFrom("value2")},
			},
			err: false,
		},
		"generic_metadata": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.StringFrom("key"), null.StringFrom("value")}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{Annotations: MetadataMap{"key": null.StringFrom("value")}, Labels: MetadataMap{"keyName": null.StringFrom("value")}},
			err:              false,
		},
		"no_slice": {
			annotationsSlice: "not-a-annotation-slice",
			labelSlice:       "not-a-label-slice",
			expectedMetadata: Metadata{},
			err:              true,
		},
		"invalid_value_string": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.StringFrom("key"), null.NewString("broken", false)}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{
				Annotations: MetadataMap{"key": null.NewString("broken", false)},
				Labels:      MetadataMap{"keyName": null.StringFrom("value")},
			},
			err: false,
		},
		"invalid_key_string": {
			annotationsSlice: []*struct{ Key, Value null.String }{{null.NewString("invalid-key", false), null.StringFrom("value")}},
			labelSlice:       []*struct{ KeyName, Value null.String }{{null.StringFrom("keyName"), null.StringFrom("value")}},
			expectedMetadata: Metadata{},
			err:              true,
		},
		"nil_metadata": {
			annotationsSlice: nil, labelSlice: nil,
			expectedMetadata: Metadata{Labels: MetadataMap{}, Annotations: MetadataMap{}},
			err:              false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := GetMetadata(tc.annotationsSlice, tc.labelSlice)
			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedMetadata, got)
			}
		})
	}
}
