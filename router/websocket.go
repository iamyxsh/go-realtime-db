package router

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/iamyxsh/go-realtime-db/handlers"
)

func WsRouter(app *fiber.App) {
	app.Use("/ws", handlers.HandleWsConnection)
	app.Get("/ws/:database/:table", websocket.New(handlers.HandleTableWS))
}
