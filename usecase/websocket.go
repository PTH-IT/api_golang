package usecase

import (
	"PTH-IT/api_golang/utils"
	"log"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type client struct {
	conn *websocket.Conn
	id   string
}

type clientsMap struct {
	sync.Mutex
	clients map[*websocket.Conn]string
}

func (cm *clientsMap) addClient(c *client) {
	cm.Lock()
	defer cm.Unlock()
	cm.clients[c.conn] = c.id
}

func (cm *clientsMap) removeClient(Conn *websocket.Conn) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.clients, Conn)
}

func (cm *clientsMap) getClientByID(Conn string) (*client, bool) {
	cm.Lock()
	defer cm.Unlock()
	for conn, connID := range cm.clients {
		if connID == Conn {
			return &client{conn: conn, id: connID}, true
		}
	}
	return nil, false
}

var (
	upgrader = websocket.Upgrader{}
)

func (i *Interactor) GetMessage(c echo.Context) error {
	authercations := c.QueryParams().Get("token")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return c.String(http.StatusForbidden, "token awrong")
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	cm := &clientsMap{
		clients: make(map[*websocket.Conn]string),
	}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)

	}
	if err != nil {
		return err
	}
	// Generate a unique connection ID
	connID := uuid.New().String()

	// Add the client to the clients map
	cm.addClient(&client{conn: conn, id: connID})

	i.referrance.UpdateConnectionID(userID, connID)
	// Start a goroutine to read messages from the client
	go func() {
		defer func() {
			// Remove the client from the clients map when the connection is closed
			cm.removeClient(conn)
		}()
		for {

			_, msg, err := conn.ReadMessage()

			if err != nil {
				log.Println("Read error:", err)
				return
			} else {
				result, err := i.referrance.GetConnectionID(userID)

				if err != nil {
					return
				}
				if result == nil {
					return
				} else {
					i.handleMsg(cm, result.ConnectionId, string(msg))
				}

			}

		}
	}()
	return nil
}

func (i *Interactor) handleMsg(cm *clientsMap, connection string, msg string) {
	// Get the client by connection ID
	client, ok := cm.getClientByID(connection)
	if !ok {
		return
	} else if client != nil {

		// Send a message to the client
		err := client.conn.WriteMessage(websocket.TextMessage, []byte("Received message: "+msg))
		if err != nil {
			log.Println("Write error:", err)
			return
		}

	}

}
