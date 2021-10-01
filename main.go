package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Template view engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public/static")

	app.Get("/", func(c *fiber.Ctx) error {
		ipAddress := c.IP()
		header := c.Request().Header.Peek("User-Agent")
		fmt.Println(ipAddress)
		fmt.Printf("%s", header)
		
		// Render index
	   return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"IpAdress": ipAddress,
			"header": string(header),
		}, "layouts/main")
	})

	err := app.Listen(":4400")
	if err != nil {
		panic(err)

	}
}