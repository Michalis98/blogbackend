package routes

import (
	"github.com/Michalis98/blogbackend/controller"
	"github.com/Michalis98/blogbackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register*", controller.Register)
	//anything below the middleware.IsAuth will not be executed if they are not auth first
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuth)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPosts)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost", controller.DeletePost)
	app.Post("api/uploads-image", controller.Upload)
	app.Static("/api/uploads", "./uploads")

}
