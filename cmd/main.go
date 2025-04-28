package main

import (
	"circle/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
