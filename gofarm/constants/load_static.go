package constants

import "github.com/labstack/echo/v4"

func LoadStatic(app *echo.Echo) {

	//get app path
	// path, _ := os.Executable()
	// // get file path
	// filePath := filepath.Dir(path)
	// //
	// staticFolder := fmt.Sprintf("%v/repository/assets", filePath)

	app.Static("static", "repository/assets")
}
