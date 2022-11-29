package web

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/utils/gbl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

type WebEvent collections.Event

type Client struct{}

type GloomWeb struct {
	listenAddress string

	clients    map[*websocket.Conn]Client
	register   chan *websocket.Conn
	unregister chan *websocket.Conn

	broadcast chan string
	// events      []WebEvent
	queueOutWeb *chan *collections.Event

	app *fiber.App
}

func NewGloomWeb(listenAddress string, queueOutWeb *chan *collections.Event) *GloomWeb {
	engine := html.New("./www", ".html")
	if err := engine.Load(); err != nil {
		gbl.Log.Error(err)
	}

	// t, err := template.ParseFiles(templateFiles...)
	// if err != nil {
	// 	gbl.Log.Error(err)
	// }

	gbl.Log.Infof("gloomWeb| engine.Templates: %+v | t: %+v", engine.Templates.DefinedTemplates(), 0)

	app := fiber.New(fiber.Config{Views: engine, ViewsLayout: "layout"})

	// gloomberg web server/ui based on fibre web framework
	gw := &GloomWeb{
		listenAddress: listenAddress,
		clients:       make(map[*websocket.Conn]Client),
		register:      make(chan *websocket.Conn, 1),
		unregister:    make(chan *websocket.Conn, 1),
		broadcast:     make(chan string, 128),
		queueOutWeb:   queueOutWeb,
		app:           app,
	}

	gw.app.Static("/static", "./www/static")
	gw.app.Static("/", "./www/home.html")

	gbl.Log.Infof("gloomWeb| static routes loaded")

	gw.app.Use(favicon.New())

	gw.app.Get("/rekt", func(c *fiber.Ctx) error {
		gbl.Log.Infof("c.Render(view, fiber.Map{}): %+v", c.Render("view", fiber.Map{}))
		return c.Render("layout", fiber.Map{
			"Title": "Gloomberg",
		})
	})

	gw.app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	// gw.app.Use("/rekt", func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
	// 		return c.Next()
	// 	}
	// 	return c.SendStatus(fiber.StatusUpgradeRequired)
	// })

	gw.app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			gw.unregister <- c
			c.Close()
		}()

		// register the client
		gw.register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					gbl.Log.Errorf("read error: %s", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			gbl.Log.Debugf("ws message received (%s): %s", messageType, message)

			if messageType == websocket.TextMessage {
				// broadcast the received message
				gw.broadcast <- string(message)
			} else {
				gbl.Log.Infof("websocket message received of type: %v", messageType)
			}
		}
	}))

	gbl.Log.Infof("gloomWeb| configuration done") //: %+v", gw.app.GetRoutes())

	gbl.Log.Infof("gloomWeb| routes: %v", app.GetRoutes(false))

	return gw
}

func (gw *GloomWeb) Run() error {
	// start hub for websocket connections/message distribution
	gbl.Log.Info("gloomWeb| starting hub...")

	go gw.runHub()

	gbl.Log.Infof("gloomWeb| starting webserver on %s", gw.listenAddress)

	return gw.app.Listen(gw.listenAddress)
}

func (gw *GloomWeb) runHub() {
	for {
		select {
		case message := <-gw.broadcast:
			gbl.Log.Debugf("gloomWeb| message received: %s", message)

			// send message to all clients
			for connection := range gw.clients {
				gbl.Log.Debugf("gloomWeb| message received: %s | connection: %+v", message, connection)

				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					gbl.Log.Warnf("write error for %v: %s", connection, err)

					err := connection.WriteMessage(websocket.CloseMessage, []byte{})
					if err != nil {
						gbl.Log.Warnf("write close-msg error for %v: %s", connection, err)
					}

					connection.Close()
					delete(gw.clients, connection)
				}
			}

		case connection := <-gw.register:
			// register new client
			gw.clients[connection] = Client{}

			gbl.Log.Info("connection registered")

		case connection := <-gw.unregister:
			// remove client
			delete(gw.clients, connection)

			gbl.Log.Info("connection unregistered")

		default:
			if message := <-*gw.queueOutWeb; message != nil {
				gbl.Log.Debugf("gloomWeb| received message from queue: %+v", message)
				gw.broadcast <- fmt.Sprintf("%+v", message)
			}
		}
	}
}
