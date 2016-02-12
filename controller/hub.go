package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Data   string `json:"data"`
}

type Hub struct {
	ws        *websocket.Conn
	send      chan []byte
	broadcast chan []byte
	topics    map[string]func(...string) []byte
}

// Constructor.
func NewHub(ws *websocket.Conn, t map[string]func(...string) []byte) *Hub {
	return &Hub{
		ws:        ws,
		send:      make(chan []byte, 256),
		broadcast: make(chan []byte),
		topics:    t,
	}
}

// Send data.
func (this *Hub) Send(data []byte) {
	log.Println("Sent:", string(data))
	this.send <- data
}

// Connection reader.
func (this *Hub) Reader() {
	defer this.ws.Close()
	log.Println("Reader started.")
	for {
		_, message, err := this.ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		this.broadcast <- message
	}
}

// Connection writer.
func (this *Hub) Writer() {
	defer this.ws.Close()
	log.Println("Writer started.")
	for {
		select {
		case message, ok := <-this.send:
			if ok {
				err := this.ws.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println(err)
					break
				}
			}
		}
	}
}

// Main method.
func (this *Hub) Run() {
	log.Println("Hub started.")
	for {
		select {
		case b := <-this.broadcast:

			log.Println("Received:", string(b))

			// unmarshal message
			var msg WSMessage
			err := json.Unmarshal(b, &msg)
			if err != nil {
				log.Println(err)
				break
			}

			// check action
			switch msg.Action {
			case "PUBLISH":
				// get func by topic name
				if fn, ok := this.topics[msg.Topic]; ok {
					if m := fn(msg.Data); m != nil {
						this.Send(m)
						time.Sleep(time.Millisecond * 500)
					}
				}
			}
		}
	}
}