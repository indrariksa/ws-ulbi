package main

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	"github.com/indrariksa/ws-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/indrariksa/ws-ulbi/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	site.Use(logger.New())
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
