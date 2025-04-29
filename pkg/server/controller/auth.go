package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"circle/pkg/server/controller/schema"
	"circle/pkg/server/usecase"
)

type IAuthController interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	SignOut(c echo.Context) error
}

type authController struct {
	au usecase.IAuthUsecase
}

func NewAuthController(au usecase.IAuthUsecase) IAuthController {
	return &authController{
		au: au,
	}
}

func (ac *authController) SignIn(c echo.Context) error {
	req := schema.SignInRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	token, userID, err := ac.au.SignIn(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, schema.SignInResponse{
		Token:  token,
		UserID: userID,
	})
}

func (ac *authController) SignUp(c echo.Context) error {
	req := schema.SignUpRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	token, err := ac.au.SignUp(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, schema.SignUpResponse{
		Token: token,
	})
}

func (ac *authController) SignOut(c echo.Context) error {
	err := ac.au.SignOut()
	if err != nil {
		return err
	}
	return nil
}
