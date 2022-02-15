package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Health Check
// @Title        Health Check
// @Description  Verify status microservice
// @Produce  json
// @Success 200 {object} response.HealthResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /health-check [get]
func HealthCheck(context echo.Context) error {
	context.JSON(http.StatusOK, echo.Map{
		"status": "OK",
	})
	return nil
}
