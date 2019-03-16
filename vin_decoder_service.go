package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	. "github.com/daandi/vingo"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/vin/{vin:.{17,18}}", vinHandler)

	timeout := 15 * time.Second

	srv := &http.Server{
		Handler:      r,
		Addr:         "[::1]:8000",
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	log.Fatal(srv.ListenAndServe())

}

func vinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]
	decodedVin := DecodeVIN(vin)
	vinJSON, _ := json.Marshal(decodedVin)

	w.WriteHeader(http.StatusOK)
	w.Write(vinJSON)

}
