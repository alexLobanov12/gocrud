package main 

import (
	"log"
	"net/http"
	"gocrud/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/kylelemons/go-gypsy/yaml"
	_ "os"
	_ "fmt"
	_ "reflect"
)

var (
	apiV1Prefix string
)

func CreateRoutes() {
	apiV1Prefix = "/api/v_1"
	router := httprouter.New()
	log.Println("run server...")
	MController := controllers.MainController{}
	router.GET("/", MController.Home)
	router.GET("/notes", MController.Notes)
	router.GET("/notes/:id", MController.ShowNote)
	router.GET(apiV1Prefix + "/ping", MController.Ping)

	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	port, _ := config.Get("port")
	errorServer := http.ListenAndServe(port, router)
	log.Fatal(errorServer)
}