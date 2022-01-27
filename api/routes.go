package main

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/http"
)

func (app *application) routes() http.Handler {
	route := fiber.New()
	route.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	//route.Get("/", adaptor.HTTPHandlerFunc(app.Home))
	route.Post("/api/message", app.saveMessage)
	route.Get("/api/message", app.getMessages)
	route.Put("/api/message/:msgID", app.updateMessage)
	route.Delete("/api/message/:msgID", app.deleteMessage)
	route.Post("/api/user", app.newUser)
	route.Post("/api/login", app.login)

	//route.Route("/admin", func(route chi.Router) {
	//	route.Use(app.Auth)
	//	route.Get("/all-customers", app.AllCustomers)
	//})
	//
	//
	//// auth routes
	//route.Get("/login", app.LoginPage)
	//route.Post("/login", app.PostLoginPage)
	//route.Get("/logout", app.Logout)

	return adaptor.FiberApp(route)
}
