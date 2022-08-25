package main

import (
	"doit/db"
	"doit/web/controllers"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func init() {
	db.LoadTheEnv()
	db.CreateDBInstance()
	db.NewClient()
}

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.AllowMethods(iris.MethodOptions)

	{
		app.Post("/login", controllers.Login)
		app.Post("/register", controllers.Regiter)
		app.Post("/logout", j.Serve, controllers.Logout)

		app.Get("/tasks", j.Serve, controllers.GetTasks)
		app.Get("/tasks/{id}", j.Serve, controllers.SearchTasks)
		app.Get("/tasks/{type}", j.Serve, controllers.SearchTypeTasks)

		app.Post("/tasks/create", j.Serve, controllers.CreateTasks)
		app.Post("/tasks/{id}", j.Serve, controllers.UpdoTasks)
		app.Post("/tasks/complete/{id}", j.Serve, controllers.CompleteTask)

		app.Delete("/tasks/{id}", j.Serve, controllers.DeleteOneTasks)
		app.Delete("/tasks", j.Serve, controllers.DeleteAllTasks)

		app.Get("/daily", controllers.GetDaily)
		app.Post("/daily", controllers.CreateDaily)
		app.Delete("/daily/{id}", controllers.DeleteOneDaily)

	}

	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.HandleDir("/public", iris.Dir("./web/public"))

	app.Listen(":8080", iris.WithOptimizations)

}

var mySecret = []byte("secret")
var j = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: true,
	// Extract by the "token" url.
	// There are plenty of options.
	// The default jwt's behavior to extract a token value is by
	// the `Authorization: Bearer $TOKEN` header.
	Extractor: jwt.FromAuthHeader,
	// When set, the middleware verifies that tokens are
	// signed with the specific signing algorithm
	// If the signing method is not constant the `jwt.Config.ValidationKeyGetter` callback
	// can be used to implement additional checks
	// Important to avoid security issues described here:
	// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
})
