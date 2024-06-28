package main

import (
	"log"
	"net/http"

	"github.com/nigerdyanes/rpsweb/handlers"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"

	log.Printf("Server on http:localhost%s", port)

	log.Fatal(http.ListenAndServe(port, router))

}
