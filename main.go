package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fusco2k/go-crud-v2/model"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userC *mongo.Collection

func main() {
	//create a context for comunicate with mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
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
	//assign a pointes collection
	userC = client.Database("testdb").Collection("user")

	//handling the redirects
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
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
	var results []model.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	csr, err := userC.Find(ctx, bson.D{})
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		user := model.User{}
		err = csr.Decode(&user)
		if err != nil {
			cancel()
			log.Fatal(err)
		}
		results = append(results, user)
	}
	if err := csr.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(results)
}

func postData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	userC.InsertOne(nil, user)

	getData(w, r)
}
