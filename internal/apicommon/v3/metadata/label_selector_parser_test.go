//go:build unit
// +build unit

package metadata_test

import (
	"errors"
	"testing"

	. "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/metadata"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		inputLabelSelector     string
		expectedLabelSelectors LabelSelectorSlice
		expectedErr            error
	}{
		"exists/not exists": {
			inputLabelSelector: "code,!chargeback-code,a",
			expectedLabelSelectors: []LabelSelector{
				{Key: "code", Operator: Exists, Values: nil},
				{Key: "chargeback-code", Operator: NotExists, Values: nil},
				{Key: "a", Operator: Exists, Values: nil},
			},
		},
		"equal/equal alt/not equal": {
			inputLabelSelector: "foo=bar,baz==hi,boo!=faz",
			expectedLabelSelectors: []LabelSelector{
				{Key: "foo", Operator: Eq, Values: []string{"bar"}},
				{Key: "baz", Operator: EqAlt, Values: []string{"hi"}},
				{Key: "boo", Operator: NotEq, Values: []string{"faz"}},
			},
		},
		"in/notin": {
			inputLabelSelector: "boo in (bar,baz),foo notin (far,faz)",
			expectedLabelSelectors: []LabelSelector{
				{Key: "boo", Operator: In, Values: []string{"bar", "baz"}},
				{Key: "foo", Operator: NotIn, Values: []string{"far", "faz"}},
			},
		},
		"all": {
			inputLabelSelector: "env=dev,a.blah==b,nope.x!=yy,code,!chargeback-code,a,tier in (backend,worker),other-foo notin (x,z)",
			expectedLabelSelectors: []LabelSelector{
				{Key: "env", Operator: Eq, Values: []string{"dev"}},
				{Key: "a.blah", Operator: EqAlt, Values: []string{"b"}},
				{Key: "nope.x", Operator: NotEq, Values: []string{"yy"}},
				{Key: "code", Operator: Exists, Values: nil},
				{Key: "chargeback-code", Operator: NotExists, Values: nil},
				{Key: "a", Operator: Exists, Values: nil},
				{Key: "tier", Operator: In, Values: []string{"backend", "worker"}},
				{Key: "other-foo", Operator: NotIn, Values: []string{"x", "z"}},
			},
			expectedErr: nil,
		},
		"bad in": {
			inputLabelSelector: "boo in (bar,baz",
			expectedErr:        errors.New("could not parse label selectors 'boo in (bar,baz'"),
		},
	}

	lsp := NewLabelSelectorParser()
	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			got, err := lsp.Parse(testCase.inputLabelSelector)
			if testCase.expectedErr != nil {
				assert.Equal(t, testCase.expectedErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedLabelSelectors, got)
		})
	}
}
