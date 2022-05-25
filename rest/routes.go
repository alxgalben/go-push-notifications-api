package rest

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/genproto/googleapis/type/date"
	"net/http"
)

type FirebaseInfo struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	AppName  string `json:"appname"`
}

type Subscriber struct {
	Username     string
	Token        string
	Applications []Application
}

type Application struct {
	AppId          string
	ExpirationDate date.Date
}

func Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.ServeFile(writer, request, "./web/index.html")
	http.ServeFile(writer, request, "./static/firebase-messaging-sw.js")
}

func AddSubscriber(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var requestBody FirebaseInfo
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestBody.Token)
	fmt.Println(requestBody)

	//Here you need to create a subscriber and save it into the database

}
