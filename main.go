package main

import (
	"log"
	"net/http"
	"html/template"

	//"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

/*
var clients = make(map[*websocket.Conn] string)
var broadcast = make(chan YTSong)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
*/

type Playlist struct {
	NumberOfSongs	int 		`json"numberOfSongs"`
	Songs			[]YTSong	`json"songs"`
}

type YTSong struct {
	Image	string `json:"image"`
	Title	string `json:"title"`
	Artist	string `json:"artist"`
}

func main() {

	r := mux.NewRouter()

	r.PathPrefix("/css").Handler(http.FileServer(http.Dir("src/css/")))
	r.PathPrefix("/js").Handler(http.FileServer(http.Dir("src/js/")))

	r.HandleFunc("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		log.Printf("idiididid %v", id)
	})

	tmpl := template.Must(template.ParseFiles("../Party-Playlist/src/index.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("We're in hello")
		data := Playlist{
			NumberOfSongs: 10,
			Songs: []YTSong{
				{Image: "Image 1", Title: "Song title 1", Artist: "Artist 1"},
				{Image: "Image 1", Title: "Song title 1", Artist: "Artist 1"},
				{Image: "Image 1", Title: "Song title 1", Artist: "Artist 1"},
			},
		}
		tmpl.Execute(w, data)
	})

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}