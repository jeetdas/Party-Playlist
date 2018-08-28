package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"os"
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
	fs := http.FileServer(http.Dir("src/"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleNewRequests()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	port = ":" + port
	log.Println("Http server started on " + port)
	err := http.ListenAndServe(port, nil)
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