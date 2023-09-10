package main

import (
	"log"
	"net/http"
	"gocrud/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	log.Println("run server...")
	MController := controllers.MainController{}
	router.GET("/", MController.Home)
	router.GET("/notes", MController.Notes)
	router.GET("/notes/:id", MController.ShowNote)
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}