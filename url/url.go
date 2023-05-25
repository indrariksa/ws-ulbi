package url

import (
	"github.com/gofiber/swagger"
	"github.com/indrariksa/ws-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)
	page.Get("/presensi", controller.GetAllPresensi)
	page.Get("/presensi/:id", controller.GetPresensiID)
	page.Get("/phone/:id", controller.GetPresensiFromPhoneNumber)
	page.Get("/tes", controller.GetAll)
	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
	page.Get("/docs/*", swagger.HandlerDefault)
}
