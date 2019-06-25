package main

import (
	"echo-swag/handler"
	"net/http"

	"github.com/labstack/echo"
)

// @title Echo to Swagger Example API
// @version 1.0
// @description This is a sample server echo-to-swagger server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	e := echo.New()
	h := handler.NewHandler()
	pathPrefix := "/users/"

	e.GET(pathPrefix+"all", h.ListAllUsers)
	e.GET(pathPrefix+":id", h.GetUser)
	e.POST(pathPrefix+"add", h.AddUser)
	e.DELETE(pathPrefix+":id", h.DeleteUser)
	e.PATCH(pathPrefix+":id", h.UpdateUser)

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
