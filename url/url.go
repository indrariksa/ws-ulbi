package url

import (
	"github.com/indrariksa/ws-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/tes", controller.GetAll)
	page.Get("/all", controller.GetAll2)
	page.Post("/ins", controller.InsertData)

}
