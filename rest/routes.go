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
var subscribers []string

const schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"subscriber": &memdb.TableSchema{
			Name: "subscriber",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Username"},
				},
				"token": &memdb.IndexSchema{
					Name:    "token",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "Token"},
				},
				"applications": &memdb.IndexSchema{
					Name:    "applications",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Applications"},
				},
			},
		},
	},
}

//Create a new data base
db, err := memdb.NewMemDB(schema)
if err != nil {
	panic(err)
}

func InitDatabase() 

subscribers := []*Subscriber{
	&Subscriber{"ness123", "hhffhfbh464646", []string{}},
	&Subscriber{"codrin", "dbdgdv54647", []string{}]},
	&Subscriber{"ana", "dgddget6464", []string{}]},
	&Subscriber{"alex", "ddhdbdb363636", []string{}},
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

func RemoveSubscriber(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var requestBody FirebaseInfo
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestBody.Token)
	fmt.Println(requestBody)

	

}

func Update(w http.ResponseWriter, r *http.Request, params httprouter.Params){

}
