package main

import (
	"circle/controller"
	"circle/db"
	"circle/infrustructure"
	"circle/router"
	"circle/usecase"
)

func main() {
	client := db.NewClient()
	ai := infrustructure.NewAuthInfrustructure(client)
	au := usecase.NewAuthUsecase(ai)
	ac := controller.NewAuthController(au)

	e := router.NewRouter(ac)

	e.Logger.Fatal(e.Start(":8080"))
}
