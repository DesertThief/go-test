package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize the HTML template engine
	engine := html.New("./", ".html")

	// You can add functions to your templates
	engine.AddFunc("greet", func(name string) string {
		return "Hello, " + name + "!"
	})

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Define a route to render the HTML template
	app.Get("/", func(c *fiber.Ctx) error {
		// c.Render accepts the template name and a fiber.Map for data
		return c.Render("index", fiber.Map{
			"Title":   "Welcome to my Fiber App!",
			"Content": "This is some dynamic content rendered from the server.",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
