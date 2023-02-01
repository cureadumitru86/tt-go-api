// This defines the handlers that will handle a request and return a response
package handlers

import (
	"fmt"
	"net/http"

	"github.com/cureadumitru86/tt-go-api/src/logic"
	"github.com/labstack/echo/v4"
)

// Response - Response object
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ProductRequest - Request object for Product
type ProductRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price" default:0.00`
}

// @Summary Errors.
// @Description Returns a 500 error.
// @Tags demo
// @Accept */*
// @Produce json
// @Success 500 {object} map[string]interface{}
// @Router /err [get]
func getErr(c echo.Context) error {
	return c.JSONPretty(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Message: "This API is all messed up!"}, " ")
}

// @Summary Home.
// @Description Home route.
// @Tags demo
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func home(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, &Response{Status: http.StatusOK, Message: "GET: Eat your vegetables"}, " ")
}

// @Summary Fibbonacci Nth number sequence.
// @Description Returns the random Nth fibbonacci number based on a preset upper and lower limit.
// @Tags demo
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /fibb [get]
func getFib(c echo.Context) error {
	num, i, err := logic.CalcFib(10, 50)
	if err != nil {
		c.Error(err)
	}
	return c.JSONPretty(http.StatusOK, &Response{Status: http.StatusOK, Message: fmt.Sprintf("The %vth fibonacci number is %v", num, i)}, " ")
}

// @Summary Add product.
// @Description Adds a product to the database.
// @Param product body ProductRequest true "Product definition"
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /product [post]
func addProduct(c echo.Context) error {
	cc := c.(*ServiceContext)
	p := ProductRequest{}
	c.Bind(&p)
	db, err := cc.DB()
	if err != nil {

		cc.JSONPretty(http.StatusInternalServerError, &Response{Status: http.StatusInternalServerError, Message: err.Error()}, " ")
	}
	id, err := logic.InsertProduct(db, p.Name, p.Price)
	if err != nil {
		cc.Error(err)
	}
	return cc.JSONPretty(http.StatusOK, &Response{Status: http.StatusOK, Message: fmt.Sprintf("Product ID %v was insert successfully.", id)}, " ")
}
