package config

import (
	"github.com/dharmasastra/restWithIris/app/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func NewRouter() *iris.Application {

	app := iris.New()
	app.Use(recover.New())

	requestLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,

		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	})

	app.Use(requestLogger)

	app.Get("/", controllers.GetAllPerson)
	app.Post("/person", controllers.CreatePerson)
	app.Get("/{name:string}", controllers.GetPerson)
	app.Post("/{name:string}", controllers.UpdatePerson)
	app.Delete("/{name:string}", controllers.DeletePerson)

	return app
}
