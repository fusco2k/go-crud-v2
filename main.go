package main

import (
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/patient", func (w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case "GET":
			getData(w, r)
		case "POST":
			postData(w, r)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, r *http.Request){
	//do something
}

func postData(w http.ResponseWriter, r *http.Request){
	//do something
}