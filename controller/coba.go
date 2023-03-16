package controller

import (
	"github.com/aiteung/musik"
	"github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	"github.com/indrariksa/ws-ulbi/config"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := presensi.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}
