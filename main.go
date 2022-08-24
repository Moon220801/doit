package main

import (
	"doit/db"
	"doit/web/controllers"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

/*func init() {
	db.LoadTheEnv()
	db.CreateDBInstance()
}*/

func init() {
	db.LoadTheEnv()
	db.CreateDBInstance()
	db.NewClient()
}

func main() {
	app := iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	app.Use(crs)
	app.Logger().SetLevel("debug")

	//app.Use(recover.New())
	//app.Use(logger.New())

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Get("/login", func(ctx iris.Context) {
		ctx.View("login.html")
	})

	app.Get("/register", func(ctx iris.Context) {
		ctx.View("register.html")
	})

	app.AllowMethods(iris.MethodOptions)

	{
		app.Post("/login", controllers.Login)
		app.Post("/register", controllers.Regiter)

		app.Get("/tasks", controllers.GetTasks)
		app.Get("/tasks/{id}", controllers.SearchTasks)
		app.Post("/tasks/create", controllers.CreateTasks)
		app.Post("/tasks/{id}", controllers.UpdoTasks)
		app.Delete("/tasks/{id}", controllers.DeleteOneTasks)
		app.Delete("/tasks", controllers.DeleteAllTasks)

	}

	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.HandleDir("/public", iris.Dir("./web/public"))

	app.Listen(":8080", iris.WithOptimizations)

}
