package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	templates "github.com/ppp3ppj/htmx-basic-go-fiber/templetes"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())


    app.Get("/", func(c *fiber.Ctx) error {
        component := templates.Index()
        return Render(c, component)
    })

    app.Use(NotFoundMiddleware)
    log.Fatal(app.Listen(":1323"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
    return Render(c, templates.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
    componentHandler := templ.Handler(component)
    for _, o := range options {
        o(componentHandler)
    }
    return adaptor.HTTPHandler(componentHandler)(c)
}
