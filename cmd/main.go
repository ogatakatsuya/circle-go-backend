package main

import (
	"circle/pkg/server/controller"
	"circle/pkg/server/db"
	"circle/pkg/server/infrustructure"
	"circle/pkg/server/router"
	"circle/pkg/server/usecase"
)

func main() {
	client := db.NewClient()
	ai := infrustructure.NewAuthInfrustructure(client)
	au := usecase.NewAuthUsecase(ai)
	ac := controller.NewAuthController(au)

	e := router.NewRouter(ac)

	e.Logger.Fatal(e.Start(":8080"))
}
