//go:build unit
// +build unit

//nolint:testpackage // we have to assign package level vars due to models using static functions
package buildpacks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

func TestGetBuildpackState(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		buildpackFileName      null.String
		expectedBuildpackState string
	}{
		"buildpack in ready state":           {buildpackFileName: null.StringFrom("test.zip"), expectedBuildpackState: StateReady},
		"buildpack in awaiting upload state": {buildpackFileName: null.NewString("", false), expectedBuildpackState: StateAwaitingUpload},
	}

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			buildpack := models.Buildpack{Filename: testCase.buildpackFileName}
			buildpackState := GetBuildpackState(&buildpack)
			require.Equal(t, testCase.expectedBuildpackState, buildpackState)
		})
	}
}

func TestBuildpackResponseObject(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		buildpack                                                           models.Buildpack
		expectedBuildpackGUID, expectedBuildpackDate, expectedBuildpackName string
		expectedResourcePath, expectedResourcePathUpload                    pagination.Link
	}{
		"buildpack valid response": {
			models.Buildpack{
				GUID:      "123",
				CreatedAt: time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC),
				UpdatedAt: null.Time{Time: time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC)},
				Name:      "Testbuildpack",
			},
			"123", "2021-08-16T13:14:15Z", "Testbuildpack",
			pagination.Link{Href: "v3/buildpack", Method: ""},
			pagination.Link{Href: "v3/buildpack/upload", Method: "POST"},
		},
	}

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			response, _ := NewPresenter().ResponseObject(&testCase.buildpack, "v3/buildpack")
			require.Equal(t, testCase.expectedBuildpackGUID, response.GUID)
			require.Equal(t, testCase.expectedBuildpackDate, response.CreatedAt)
			require.Equal(t, testCase.expectedBuildpackDate, response.UpdatedAt)
			require.Equal(t, testCase.expectedBuildpackName, response.Name)
			require.Equal(t, testCase.expectedResourcePath, response.Links.Self)
			require.Equal(t, testCase.expectedResourcePathUpload, response.Links.Upload)
		})
	}
}

func TestBuildpackResponseObjectSlice(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		buildpack                                      models.BuildpackSlice
		expectedBuildpackGUID1, expectedBuildpackGUID2 string
	}{
		"buildpack valid response": {models.BuildpackSlice{
			&models.Buildpack{GUID: "123"},
			&models.Buildpack{GUID: "456"},
		}, "123", "456"},
	}

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			response, _ := NewPresenter().ListResponseObject(testCase.buildpack, int64(len(testCase.buildpack)), pagination.Default(), "v3/buildpack")
			require.Equal(t, testCase.expectedBuildpackGUID1, response.Resources[0].GUID)
			require.Equal(t, testCase.expectedBuildpackGUID2, response.Resources[1].GUID)
		})
	}
}
