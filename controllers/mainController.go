package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"gocrud/models"
	"gocrud/helpers"
	"log"
	_ "io/ioutil"
	"fmt"
)

type MainController struct {
	HomeMessage string
	Note string
}

func (m *MainController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m.HomeMessage = "home endpoint from struct"
	w.Write([]byte(m.HomeMessage))
}

func (m *MainController) Notes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m.Note = "get all notes"
	w.Write([]byte(m.Note))
}

func (m *MainController) ShowNote(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	m.Note = string(params.ByName("id"))
	w.Write([]byte(m.Note))
}

func (m *MainController) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var message = models.Message{
		Message: "pong",
	}
	json.NewEncoder(w).Encode(message)
	helpers.Log(w, r)
}


func (m *MainController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var loginData models.Login
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("empty data")
			return
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(&loginData)

}
