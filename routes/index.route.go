package routes

import (
	"go-relation/relasi-gorm/controllers"

	"go-relation/relasi-gorm/configs"
	"go-relation/relasi-gorm/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {

	// JWT INIT
	jwt := middlewares.JWTMiddleware(configs.Secret)

	app.Post("/login", controllers.Login)
	app.Get("/protected", jwt, controllers.Protected)

	app.Get("/users", jwt, controllers.UserGetAll)
	app.Post("/users", jwt, controllers.CreateUser)

	app.Get("/lockers", jwt, controllers.LockerGetAll)
	app.Post("/lockers", jwt, controllers.CreateLocker)

	app.Get("/posts", jwt, controllers.PostGetAll)
	app.Post("/posts", jwt, controllers.CreatePost)

	app.Get("/tags", jwt, controllers.TagGetAll)
	app.Post("/tags", jwt, controllers.CreateTag)
}
