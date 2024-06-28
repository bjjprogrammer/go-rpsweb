package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Game")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Game")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Play")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
