package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"template-api-gecho/constant"
	"template-api-gecho/model"
	"template-api-gecho/responsegraph"
)

type ExampleController struct {
	model model.ExampleModel
}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
	secret := os.Getenv(constant.Authorization)

	posts := ExampleController.model.GetPosts(secret)
	res := responsegraph.ResponseGenericGet{
		Status:  constant.StatusSuccess,
		Message: "Berhasil Insert Data",
		Data:    posts,
	}

	return c.JSON(http.StatusOK, res)
}
