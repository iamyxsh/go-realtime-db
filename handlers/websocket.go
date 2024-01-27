package handlers

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var WSConnections map[string]map[string]*websocket.Conn

func HandleWsConnection(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func HandleTableWS(c *websocket.Conn) {
	database := c.Params("database")
	table := c.Params("table")

	WSConnections = make(map[string]map[string]*websocket.Conn)
	WSConnections[database] = make(map[string]*websocket.Conn)
	WSConnections[database][table] = c

	var (
		err error
	)

	for {
		if _, _, err = c.ReadMessage(); err != nil {
			WSConnections[database][table] = nil
			break
		}
	}

}
