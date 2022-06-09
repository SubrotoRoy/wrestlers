package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SubrotoRoy/wrestlers/datastore"
	"github.com/SubrotoRoy/wrestlers/model"
)

func CreateWrestler(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating Wrestlers")
	var wrestler model.Wrestler
	err := json.NewDecoder(r.Body).Decode(&wrestler)

	if err == nil {
		log.Println(wrestler)
		datastore.AddWrestlers(wrestler)
	}

}
