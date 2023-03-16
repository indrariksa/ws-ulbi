package controller

import (
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	bep "github.com/indrariksa/be_presensi"
	"github.com/indrariksa/ws-ulbi/config"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/indrariksa/ws-ulbi",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}

func GetAllData(c *fiber.Ctx) error {
	getip := bep.GetAllPresensiFromStatus("masuk", config.Ulbimongoconn, "presensi")
	return c.JSON(getip)
}
