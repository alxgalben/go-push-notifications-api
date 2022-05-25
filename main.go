package main

import (
	"cloudstack_go/rest"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	//Call here Init database func
	router := httprouter.New()
	//This endpoint call should look like : localhost:8080/api/register?username=aaa&appname=bbb
	router.GET("/api/register", rest.Register)
	router.POST("/api/addSubscriber", rest.AddSubscriber)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", router)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
