// +build unit

package ratelimiter_test

import (
	"database/sql/driver"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/ratelimiter"
)

type Time struct {
	time.Time
}

// Match satisfies sqlmock.Argument interface
func (a Time) Match(v driver.Value) bool {
	t, ok := v.(time.Time)
	if !ok {
		return false
	}
	return t.UTC().Truncate(time.Second) == a.UTC().Truncate(time.Second)
}

func TestNoRequestCountExists(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	rateLimiter := RateLimiter{
		GeneralLimit:  10,
		ResetInterval: time.Minute * 1,
		DB:            db,
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "request_counts" WHERE (user_guid=$1);`)).
		WithArgs("123").
		WillReturnRows(sqlmock.NewRows(nil))

	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "request_counts" ("user_guid","count","valid_until") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs("123", 1, Time{time.Now().Add(rateLimiter.ResetInterval)}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := rateLimiter.Allow("123", context)
	assert.NoError(t, err)
	assert.True(t, result)
}

//nolint
func TestRequestCountExceedsRateLimit(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	rateLimiter := RateLimiter{
		GeneralLimit:  10,
		ResetInterval: time.Minute * 1,
		DB:            db,
	}

	expiryTime := time.Now().Add(1 * time.Hour)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "request_counts" WHERE (user_guid=$1);`)).
		WithArgs("123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_guid", "count", "valid_until"}).AddRow(1, "123", 10, expiryTime))

	mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "request_counts" SET "user_guid"=$1,"count"=$2,"valid_until"=$3 WHERE "id"=$4`)).
		WithArgs("123", 11, expiryTime, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := rateLimiter.Allow("123", context)
	assert.NoError(t, err)
	assert.False(t, result)
}

//nolint
func TestRequestCountWithinRateLimit(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	rateLimiter := RateLimiter{
		GeneralLimit:  10,
		ResetInterval: time.Minute * 1,
		DB:            db,
	}

	expiryTime := time.Now().Add(1 * time.Hour)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "request_counts" WHERE (user_guid=$1);`)).
		WithArgs("123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_guid", "count", "valid_until"}).AddRow(1, "123", 3, expiryTime))

	mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "request_counts" SET "user_guid"=$1,"count"=$2,"valid_until"=$3 WHERE "id"=$4`)).
		WithArgs("123", 4, expiryTime, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := rateLimiter.Allow("123", context)
	assert.NoError(t, err)
	assert.True(t, result)
}

// When we're above the rate limit but its reset time is in the past
func TestRequestCountExpired(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/something", nil)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	rateLimiter := RateLimiter{
		GeneralLimit:  10,
		ResetInterval: time.Minute * 1,
		DB:            db,
	}

	testTime := time.Now().UTC()
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "request_counts" WHERE (user_guid=$1);`)).
		WithArgs("123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_guid", "count", "valid_until"}).AddRow(1, "123", 57, testTime.Add(-1*time.Hour)))

	mock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "request_counts" SET "user_guid"=$1,"count"=$2,"valid_until"=$3 WHERE "id"=$4`)).
		WithArgs("123", 1, Time{time.Now().Add(rateLimiter.ResetInterval)}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := rateLimiter.Allow("123", context)
	assert.NoError(t, err)
	assert.True(t, result)
}
