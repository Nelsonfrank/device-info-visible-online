package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/mssola/user_agent"
)

func main() {
	// Template view engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public/static")
	
	app.Get("/", func(c *fiber.Ctx) error {
		ipAddress := c.Context().RemoteAddr().String()
		userAgentString := c.Request().Header.Peek("User-Agent")
		ua := user_agent.New(string(userAgentString))
		name, version := ua.Browser()
		
	    return c.Render("index", fiber.Map{
			"Title": "Device Informations",
			"IpAddress": ipAddress,
			"devicerBrowser": name,
			"devicerBrowserVersion": version,
			"deviceOs": ua.OS(),
			"deviceLocaliation": ua.Localization(),
		}, "layouts/main")
	})

	err := app.Listen(":4400")
	if err != nil {
		panic(err)

	}
}