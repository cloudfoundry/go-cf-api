package ccerrors

import (
	"fmt"
	"net/http"
	"strings"
)

type CloudControllerErrors struct {
	Errors []CloudControllerError `json:"errors"`
}

func AsErrors(errors ...CloudControllerError) CloudControllerErrors {
	return CloudControllerErrors{
		Errors: errors,
	}
}

type CloudControllerError struct {
	HTTPStatus int    `json:"-"`
	Detail     string `json:"detail,omitempty"`
	Title      string `json:"title"`
	Code       int    `json:"code"`
}

func (e *CloudControllerError) Error() string {
	return fmt.Sprintf("Code [%d] (%s): %s", e.Code, e.Title, e.Detail)
}

func BadQueryParameter() CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusBadRequest,
		Code:       10005, //nolint:gomnd // CF API error code
		Title:      "CF-BadQueryParameter",
		Detail:     "An invalid query parameter was given",
	}
}

func InvalidAuthToken() CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       1000, //nolint:gomnd // CF API error code
		Title:      "CF-InvalidAuthToken",
		Detail:     "An invalid auth token was given",
	}
}

func NotAuthenticated() CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       10002, //nolint:gomnd // CF API error code
		Title:      "CF-NotAuthenticated",
		Detail:     "No auth token was given, but authentication is required for this endpoint",
	}
}

func NotAuthorized() CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusForbidden,
		Code:       10003, //nolint:gomnd // CF API error code
		Title:      "CF-NotAuthorized",
		Detail:     "The authenticated user does not have permission to perform this operation",
	}
}

func ResourceNotFound(resourceType string) error {
	return &CloudControllerError{
		HTTPStatus: http.StatusNotFound,
		Code:       10010, //nolint:gomnd // CF API error code
		Title:      "CF-ResourceNotFound",
		Detail:     fmt.Sprintf("%s not found", strings.Title(resourceType)),
	}
}

func UnprocessableEntity(detail string) CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusUnprocessableEntity,
		Code:       10008, //nolint:gomnd // CF API error code
		Title:      "CF-UnprocessableEntity",
		Detail:     detail,
	}
}

func UnknownError() CloudControllerError {
	return CloudControllerError{
		HTTPStatus: http.StatusInternalServerError,
		Code:       10001, //nolint:gomnd // CF API error code
		Title:      "UnknownError",
		Detail:     "An unexpected error occurred",
	}
}
