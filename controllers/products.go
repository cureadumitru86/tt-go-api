package products

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func getProducts(c echo.Context) error {
	return c.String(http.StatusOK, "Product listing")

	/*
		data, err := db.GetAllProducts()
		if err != nil {
			return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
		}
		return c.JSON(http.StatusOK, data)
	*/
}

func getProduct(c echo.Context) error {
	id, err: strconv.Atoi(c.Param(id));

	/*
	data, err := db.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
	}
	*/

	return c.JSON(http.StatusOK, "Product single", id)
}
