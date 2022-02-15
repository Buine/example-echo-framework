package utils

import (
	"net/http"
	"yuno/echo-example/common/utils"
	"yuno/echo-example/response"

	"github.com/labstack/echo/v4"
)

func MiddlewareAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if !utils.ContainsInArray(context.Path(), utils.SkipRoutes) {
			context.JSON(http.StatusForbidden, response.ErrorResponse{
				Code:    "InternalError",
				Message: "Unauthorized",
			})
		} else {
			next(context)
		}
		return nil
	}
}
