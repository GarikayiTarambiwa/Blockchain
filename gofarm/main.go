package main

import (
	"github.com/GarikayiTarambiwa/Blockchain/gofarm/constants"
	"github.com/GarikayiTarambiwa/Blockchain/gofarm/router"
	"github.com/GarikayiTarambiwa/Blockchain/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//init echo
	app := echo.New()

	// load static asset
	constants.LoadStatic(app)
	//render template path
	app.Renderer = constants.LoadTemplates()
	//load router
	router.LoadRouters(app)

	// listen and serve on 0.0.0.0:8080
	server.SetServer(app)
}
