package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Description string `json:"desc"`
}

func indexHandler(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	indexData := IndexData{"Go Blog Hub", "Demo Project"}
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func main() {
	server := http.Server{
		Addr: "localhost:3000",
	}

	http.HandleFunc("/", indexHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}