package auth

import "github.com/labstack/echo/v4"

type Scope string

const (
	Read          Scope = "cloud_controller.read"
	Write         Scope = "cloud_controller.write"
	Admin         Scope = "cloud_controller.admin"
	AdminReadOnly Scope = "cloud_controller.admin_read_only"
	GlobalAuditor Scope = "cloud_controller.global_auditor"
)

func CanRead(c echo.Context) bool {
	return hasScope(Read, c)
}

func CanWrite(c echo.Context) bool {
	return hasScope(Write, c)
}

func IsAdmin(c echo.Context) bool {
	return hasScope(Admin, c)
}

func IsAdminReadOnly(c echo.Context) bool {
	return hasScope(AdminReadOnly, c)
}

func IsGlobalAuditor(c echo.Context) bool {
	return hasScope(GlobalAuditor, c)
}

func hasScope(scope Scope, c echo.Context) bool {
	scopes, ok := c.Get(Scopes).([]string)
	if !ok {
		return false
	}
	for _, s := range scopes {
		if s == string(scope) {
			return true
		}
	}
	return false
}
