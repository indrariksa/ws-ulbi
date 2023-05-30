package main

import (
	"log"

	"github.com/indrariksa/ws-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/indrariksa/ws-ulbi/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/indrariksa/ws-ulbi/docs"
)

// @title TES SWAG
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @host ws-ulbi.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
