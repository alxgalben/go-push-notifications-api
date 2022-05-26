package main

import (
	"ness_test/rest"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//Call here Init database func
	router := httprouter.New()
	//This endpoint call should look like : localhost:8080/api/register?username=aaa&appname=bbb
	router.GET("/api/register", rest.Register)
	router.POST("/api/addSubscriber", rest.AddSubscriber)
	router.DELETE("/api/removeSubscriber", rest.RemoveSubscriber)
	router.PUT("/api/update", rest.Update)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", router)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
