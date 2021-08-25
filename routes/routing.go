package routes

import (
	"github.com/labstack/echo"
	"template-api-gecho/controller"
	"template-api-gecho/utils"
)

type Routing struct {
	example controller.ExampleController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/posts/", Routing.example.GetPostsController)
	e.POST("/posts/", Routing.example.GetPostsController, utils.Middleware)

	return e
}
