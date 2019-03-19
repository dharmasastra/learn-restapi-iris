package main

import (
	"github.com/dharmasastra/restWithIris/config"
	"github.com/kataras/iris"
	"log"
)

func main() {
	app := config.NewRouter()

	log.Fatal(app.Run(iris.Addr(":8080")))
}
