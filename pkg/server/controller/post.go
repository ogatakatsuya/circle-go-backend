package controller

import (
	"circle/pkg/server/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAll(c echo.Context) error
}

type PostController struct {
	pi usecase.IPostUsecase
}

func NewPostController(pi usecase.IPostUsecase) IPostController {
	return &PostController{
		pi: pi,
	}
}

func (pc *PostController) GetAll(c echo.Context) error {
	posts, err := pc.pi.GetAll()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "failed to get posts"},
		)
	}
	return c.JSON(http.StatusOK, posts)
}
