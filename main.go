package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	//create a context for comunicate with mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//iniatiate the pointed client and connects to the mongo server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	//pings the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	fmt.Println("Mongo Connected")
	//handling the redirects 
	http.HandleFunc("/patient", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getData(w, r)
		case "POST":
			postData(w, r)
		}
	})
	//creates the tpc server on port 8080 using the default
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
