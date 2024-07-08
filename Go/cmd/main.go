package main

import (
	"net/http"

	"github.com/oxxi/jobsity/database"
	"github.com/oxxi/jobsity/routers"
)

func main() {

	database.Connection()
	router := http.NewServeMux()
	routers.RegisterRouter(router, database.GetDB())

	routersCors := routers.EnableCors(router)

	if err := http.ListenAndServe(":8080", routersCors); err != nil {
		panic(err)
	}
}
