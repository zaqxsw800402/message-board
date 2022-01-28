package main

import (
	"embed"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"net/http"
)

//go:embed templates
var templateFS embed.FS

func (app *application) routes() http.Handler {
	//engine := html.NewFileSystem(http.Dir("./templates"), ".gohtml")
	engine := html.NewFileSystem(http.FS(templateFS), ".gohtml")
	//
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	//

	mux := fiber.New(fiber.Config{
		Views: engine,
	})

	mux.Get("/", app.Home)
	mux.Get("/user", app.CreateUser)
	mux.Get("/message/:msgID", app.ModifyMessage)

	mux.Get("/login", app.Login)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)

	return adaptor.FiberApp(mux)
}
