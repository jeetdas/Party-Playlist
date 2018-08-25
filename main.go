package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	//"github.com/gorilla/mux" // To parse arguments sent in
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan YTSong)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type YTSong struct {
	Image	string	`json:"image"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
}

func main() {
	fs := http.FileServer(http.Dir("../Party-Playlist/src/"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleNewRequests()

	log.Println("Http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var ytsong YTSong
		err := ws.ReadJSON(&ytsong)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- ytsong
	}
}

func handleNewRequests() {
	log.Printf("This bitch got connected")
	for {
		ytsong := <-broadcast
		for client := range clients {
			err := client.WriteJSON(ytsong)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}