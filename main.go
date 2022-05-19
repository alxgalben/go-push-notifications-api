package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

//TODO make another struct with token(string), username(string) and project(string).
//TODO use it to save the user info and token in your in-memory database

func main() {

	router := httprouter.New()
	//This path will create a firebase service worker that handles the notifications
	router.GET("/api/register", homepage)
	//This path is used for getting the user token after the service worker is created
	//It's called inside firebase-messaging-sw.js after the service worker is up and running
	router.POST("/api/token", getToken)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", router)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.ServeFile(writer, request, "./web/index.html")
	http.ServeFile(writer, request, "./static/firebase-messaging-sw.js")
}

func getToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var request Token
	//Decode the request body to get the token needed for sending notifications
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//This is the token you need to save in your database along with username and project informations
	fmt.Println(request.Token)
}
