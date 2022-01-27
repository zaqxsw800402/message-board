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
	//mux := fiber.New()
	//engine := html.NewFileSystem(http.Dir("./templates"), ".gohtml")
	//engine := html.New("./templates", ".gohtml")
	engine := html.NewFileSystem(http.FS(templateFS), ".gohtml")
	//
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	//
	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	// Layout defines the variable name that is used to yield templates within layouts
	engine.Layout("templates/base.gohtml") // Optional. Default: "embed"
	//engine.Layout("embed") // Optional. Default: "embed"

	mux := fiber.New(fiber.Config{
		Views: engine,
	})

	//mux.Get("/", adaptor.HTTPHandlerFunc(app.Home))
	//mux.Get("/", app.Home)
	mux.Get("/", app.Home)
	mux.Get("/user", app.CreateUser)
	mux.Get("/message/:msgID", app.ModifyMessage)

	mux.Get("/login", app.Login)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	//mux.Get("/ws", app.WsEndPoint)

	//mux.Route("/admin", func(mux chi.Router) {
	//	mux.Use(app.Auth)
	//	mux.Get("/all-customers", app.AllCustomers)
	//})
	//
	//
	//// auth routes
	//mux.Get("/login", app.LoginPage)
	//mux.Post("/login", app.PostLoginPage)
	//mux.Get("/logout", app.Logout)

	return adaptor.FiberApp(mux)
}
