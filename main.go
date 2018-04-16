package main

import (
	"html/template"
	"log"
	"net/http"

	//3rd party
	"github.com/gorilla/mux"

	//local
	"golangcloudwatch/api"
)

var tpl *template.Template
var router *mux.Router

func main() {
	addControllerRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}

//addControllerRoutes registers api controllers
func addControllerRoutes() {
	router = mux.NewRouter()
	cloudwatchcontroller.AddRoutes(router, tpl)
}
