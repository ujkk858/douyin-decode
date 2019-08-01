package main

import (
	"douyin/web/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.RegisterView(iris.HTML("./web/views", ".html").
		Layout("shared/layout.html").Reload(true))

	//app.HandleDir("/public", "./web/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	mvc.New(app).Handle(new(controllers.HomeController))
	//mvc.New(app.Party("/home")).Handle(new(controllers.HomeController))

	app.Run(iris.Addr(":8080"))
}
