package main

import (
	"fmt"

	"github.com/Luis-Guillermo-Rivera-Stephens/gainerz_sql_microservice_api/app/routes"
)

func main() {
	fmt.Println("Starting server")
	Api, err := routes.GetAPI()
	if err == nil {
		fmt.Println("Server started")
		//Api.InitRoutes()
		//Api.ViewsRouter()
		Api.Listen()
	}

}
