// +build unit

package logging_test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type VcapRequestIdTestSuite struct {
	suite.Suite
	req          *http.Request
	c            echo.Context
	middlewareFn echo.MiddlewareFunc
	h            echo.HandlerFunc
	rec          *httptest.ResponseRecorder
}

func (suite *VcapRequestIdTestSuite) SetupTest() {
	e := echo.New()
	suite.req = httptest.NewRequest(http.MethodGet, "/something", nil)
	suite.rec = httptest.NewRecorder()
	suite.c = e.NewContext(suite.req, suite.rec)

	suite.h = func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	suite.middlewareFn = logging.NewVcapRequestId()
	fmt.Println("setup test")
}

func (suite *VcapRequestIdTestSuite) RunMiddleware() error {
	return suite.middlewareFn(suite.h)(suite.c)
}

func (suite *VcapRequestIdTestSuite) TestGenerateVcapRequestId() {
	err := suite.RunMiddleware()
	assert.Nil(suite.T(), err)
	vcapRequestId := suite.rec.Header().Get(logging.HeaderVcapRequestId)
	assert.NotNil(suite.T(), vcapRequestId)
	assert.Equal(suite.T(), 36, len(vcapRequestId))
}

func (suite *VcapRequestIdTestSuite) TestVcapRequestIdFromOutside() {
	suite.req.Header.Set(logging.HeaderVcapRequestId, "123-abc")
	err := suite.RunMiddleware()
	assert.Nil(suite.T(), err)
	vcapRequestId := suite.rec.Header().Get(logging.HeaderVcapRequestId)
	assert.NotNil(suite.T(), vcapRequestId)
	assert.Equal(suite.T(), 45, len(vcapRequestId))
	assert.True(suite.T(), strings.HasPrefix(vcapRequestId, "123-abc::"))
}

func (suite *VcapRequestIdTestSuite) TestRequestIdFromOutside() {
	suite.req.Header.Set(echo.HeaderXRequestID, "123-abc")
	err := suite.RunMiddleware()
	assert.Nil(suite.T(), err)
	vcapRequestId := suite.rec.Header().Get(logging.HeaderVcapRequestId)
	assert.NotNil(suite.T(), vcapRequestId)
	assert.Equal(suite.T(), 45, len(vcapRequestId))
	assert.True(suite.T(), strings.HasPrefix(vcapRequestId, "123-abc::"))
}

func (suite *VcapRequestIdTestSuite) TestVcapRequestIdFromOutsideOnlyAlphanumeric() {
	suite.req.Header.Set(logging.HeaderVcapRequestId, "123-abc.?:")
	err := suite.RunMiddleware()
	assert.Nil(suite.T(), err)
	vcapRequestId := suite.rec.Header().Get(logging.HeaderVcapRequestId)
	assert.Equal(suite.T(), 45, len(vcapRequestId))
	assert.True(suite.T(), strings.HasPrefix(vcapRequestId, "123-abc::"))
}

func (suite *VcapRequestIdTestSuite) TestVcapRequestIdFromOutsideMax255() {
	suite.req.Header.Set(logging.HeaderVcapRequestId, strings.Repeat("x", 500))
	err := suite.RunMiddleware()
	assert.Nil(suite.T(), err)
	vcapRequestId := suite.rec.Header().Get(logging.HeaderVcapRequestId)
	assert.Equal(suite.T(), 255+2+36, len(vcapRequestId))
	assert.True(suite.T(), strings.HasPrefix(vcapRequestId, strings.Repeat("x", 255)+"::"))
}

func TestVcapRequestIdSuite(t *testing.T) {
	suite.Run(t, new(VcapRequestIdTestSuite))
}
