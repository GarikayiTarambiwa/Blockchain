package pages

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	database "github.com/edmundrotimi/gofarm/db/product"
	"github.com/labstack/echo/v4"
)

func FormContext(c echo.Context) error {
	productID := c.FormValue("product_id")
	quantity := c.FormValue("quantity")
	cornColor := c.FormValue("corn_color")

	//convert quantity to number from string
	numQuantity, _ := strconv.ParseInt(quantity, 10, 64)

	//time
	currentTime := time.Now()

	database.CreateRecord(productID, cornColor, numQuantity, currentTime)

	//
	path := fmt.Sprintf("/%v", productID)

	return c.Redirect(http.StatusMovedPermanently, path)
}
