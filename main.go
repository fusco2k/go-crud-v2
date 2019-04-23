package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/patient", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getData(w, r)
		case "POST":
			postData(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonOk, err := json.Marshal("ok")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonOk)
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonOk, err := json.Marshal("ok")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonOk)
}
