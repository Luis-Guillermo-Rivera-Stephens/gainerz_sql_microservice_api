package main

import (
	"flag"
	"fmt"

	"github.com/Luis-Guillermo-Rivera-Stephens/gainerz_sql_microservice_api/app/routes"
)

func main() {

	test := flag.Bool("test", true, "base de datos de testing o la que esta desplegada en la nube")
	automigrate := flag.Bool("automigrate", false, "hacer o no la automigracion")

	flag.Parse()

	fmt.Println("Starting server")
	Api, err := routes.GetAPI(*test, *automigrate)
	if err == nil {
		fmt.Println("Server started")
		//Api.InitRoutes()
		//Api.ViewsRouter()
		Api.Listen()
	}

}
