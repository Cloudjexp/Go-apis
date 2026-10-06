package main

import (
	"encoding/json"
	"net/http"
)

type Game struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}

var games = []Game{
	{ID: 1, Title: "Horripilant", Genre: "Clicker", Rating: 10},
	{ID: 2, Title: "Hollow Knight", Genre: "Metroidvania", Rating: 10},
}

func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(games)
}

func main() {
	http.HandleFunc("/games", getGames)

	http.ListenAndServe(":8080", nil)
}
