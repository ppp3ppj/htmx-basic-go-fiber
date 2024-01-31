package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ppp3ppj/htmx-basic-go-fiber/models"
	templates "github.com/ppp3ppj/htmx-basic-go-fiber/templetes"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    run_number := 0
    mock_user := []*models.User { }
    user1 := models.User{Name: "pppAgen1"}
    user2 := models.User{Name: "pppAgen2"}

    run_number ++
    user1.Id = run_number
    mock_user = append(mock_user, &user1)

    run_number ++
    user2.Id = run_number
    mock_user = append(mock_user, &user2)

    app.Get("/", func(c *fiber.Ctx) error {
        component := templates.Index(mock_user)
        return Render(c, component)
    })

    app.Post("/foo", func(c *fiber.Ctx) error {

        user_add := models.User{Id: run_number, Name: "test"}
        name := c.FormValue("nameStr")

        if len(name) != 0 {
            user_add.Name = name
        }

        mock_user = append(mock_user, &user_add)
        fmt.Println(name)
        component := templates.UserDataDiv(&user_add)
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
