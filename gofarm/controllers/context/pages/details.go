package pages

import (
	"net/http"

	database "github.com/edmundrotimi/gofarm/db/product"
	"github.com/labstack/echo/v4"
)

func DetailsContext(c echo.Context) error {
	getProductId := c.Param("productId")

	//get db record
	dbRecord, _ := database.GetProductDetails(getProductId)

	return c.Render(http.StatusOK, "detail.html", dbRecord)
}
