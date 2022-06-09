package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SubrotoRoy/wrestlers/datastore"
)

func GetAllWrestlers(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all Wrestlers")
	response, err := json.Marshal(datastore.GetAllWrestlers())
	if err == nil {
		w.Header().Add("Content-Type", "application/json")
		w.Write(response)
	}
}
