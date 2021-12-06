package v3

import (
	"fmt"
	"net/http"
	"strings"
)

type CfApiErrors struct {
	Errors []CfApiError `json:"errors"`
}

func AsErrors(errors ...CfApiError) CfApiErrors {
	return CfApiErrors{
		Errors: errors,
	}
}

type CfApiError struct {
	HTTPStatus int    `json:"-"`
	Detail     string `json:"detail,omitempty"`
	Title      string `json:"title"`
	Code       int    `json:"code"`
	Err        error  `json:"-"`
}

func (e *CfApiError) Error() string {
	return fmt.Sprintf("Code [%d] (%s): %s", e.Code, e.Title, e.Detail)
}

func (e *CfApiError) Unwrap() error {
	return e.Err
}

func BadQueryParameter(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusBadRequest,
		Code:       10005, //nolint:gomnd // CF API error code
		Title:      "CF-BadQueryParameter",
		Detail:     "An invalid query parameter was given",
		Err:        err,
	}
}

func InvalidAuthToken(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       1000, //nolint:gomnd // CF API error code
		Title:      "CF-InvalidAuthToken",
		Detail:     "An invalid auth token was given",
		Err:        err,
	}
}

func NotAuthenticated(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       10002, //nolint:gomnd // CF API error code
		Title:      "CF-NotAuthenticated",
		Detail:     "No auth token was given, but authentication is required for this endpoint",
		Err:        err,
	}
}

func NotAuthorized(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusForbidden,
		Code:       10003, //nolint:gomnd // CF API error code
		Title:      "CF-NotAuthorized",
		Detail:     "The authenticated user does not have permission to perform this operation",
		Err:        err,
	}
}

func ResourceNotFound(resourceType string, err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusNotFound,
		Code:       10010, //nolint:gomnd // CF API error code
		Title:      "CF-ResourceNotFound",
		Detail:     fmt.Sprintf("%s not found", strings.Title(resourceType)),
		Err:        err,
	}
}

func UnprocessableEntity(detail string, err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       10008, //nolint:gomnd // CF API error code
		Title:      "CF-UnprocessableEntity",
		Detail:     detail,
		Err:        err,
	}
}

func TooManyRequests(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusTooManyRequests,
		Code:       10013, //nolint:gomnd // CF API error code
		Title:      "CF-RateLimitExceeded",
		Detail:     "Rate Limit Exceeded",
		Err:        err,
	}
}

func UnknownError(err error) *CfApiError {
	return &CfApiError{
		HTTPStatus: http.StatusInternalServerError,
		Code:       10001, //nolint:gomnd // CF API error code
		Title:      "UnknownError",
		Detail:     "An unexpected error occurred",
		Err:        err,
	}
}
