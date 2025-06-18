package common

import (
	"maps"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func HandleError(c echo.Context, err error, customCode ...int) error {
	code := http.StatusBadRequest
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if len(customCode) > 0 {
		code = customCode[0]
	}

	return c.JSON(code, map[string]any{
		"error":   err.Error(),
		"code":    code,
		"message": "error",
		"success": false,
	})
}

func HandleSuccess(c echo.Context, statusCode int, message string, dynamicMap ...map[string]any) error {
	responseData := map[string]any{
		"error":   nil,
		"code":    statusCode,
		"message": message,
		"success": true,
	}

	if len(dynamicMap) > 0 {
		for _, dm := range dynamicMap {
			maps.Copy(responseData, dm)
		}
	}

	return c.JSON(statusCode, responseData)
}

func HandleErrorService(funcName string, status int, customErrMsg string, realErr error) error {
	if realErr != nil {
		log.Errorf("Func %s Error: %s", funcName, realErr)
	}
	return echo.NewHTTPError(status, customErrMsg)
}
