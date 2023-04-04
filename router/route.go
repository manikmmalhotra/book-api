package router

import (
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", handler.UserHandlerGetAll)
	r.Get("/user/:id", middleware.Auth, handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
}
