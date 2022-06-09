package main

import (
	"log"
	"net/http"

	"github.com/SubrotoRoy/wrestlers/handlers"
)

func main() {
	log.Println("Strarting Wrestlers server")

	log.Println("Strarting Wrestlers server")
	http.HandleFunc("/get", handlers.GetAllWrestlers)
	http.HandleFunc("/create", handlers.CreateWrestler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
