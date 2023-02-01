package handlers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/cureadumitru86/tt-go-api/src/docs"
)

// @title GoDemo API
// @version 1.0
// @description GoDemo server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func registerRoutes(e *echo.Echo) {
	e.Match([]string{"POST", "GET"}, "/", home)
	e.GET("/fibb", getFib)
	e.GET("/err", getErr)
	e.POST("/product", addProduct)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
