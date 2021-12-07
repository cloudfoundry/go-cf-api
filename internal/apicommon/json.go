package apicommon

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

const INDENT = " "

type JSONSerializer struct {
	echo.DefaultJSONSerializer
}

func (d JSONSerializer) Serialize(c echo.Context, iface interface{}, indent string) error {
	enc := json.NewEncoder(c.Response())
	enc.SetEscapeHTML(false)
	enc.SetIndent("", INDENT)
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(iface)
}

func (d JSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	return d.DefaultJSONSerializer.Deserialize(c, i)
}
