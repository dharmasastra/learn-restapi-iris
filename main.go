package main

import (
	"github.com/dharmasastra/restWithIris/controllers"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	app.Get("/", controllers.GetAllPerson)
	app.Post("/person", controllers.CreatePerson)
	app.Get("/{name:string}", controllers.GetPerson)
	app.Post("/{name:string}", controllers.UpdatePerson)

	_ = app.Run(iris.Addr(":8080"))
}
