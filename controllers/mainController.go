package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"gocrud/models"
	"gocrud/helpers"
	_ "log"
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
			loginError := models.LoginError{
				Errors: []string{
					"empty json data",
				},
				Status: 400,
			}
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(loginError)
			helpers.Log(w, r)
			return
		}
	}

	errLogin := loginData.ValidateLogin()
	var errorsLoginMsg []string
	if errLogin != nil {
		for _, v := range errLogin {
			errorsLoginMsg = append(errorsLoginMsg, v.Error())
		}
	}

	if errorsLoginMsg != nil {
		loginErrorResp := models.LoginError{
			Errors: errorsLoginMsg,
			Status: 400,
		}
		json.NewEncoder(w).Encode(loginErrorResp)
		helpers.Log(w, r)
		return
	}
	fmt.Println(errorsLoginMsg)
	fmt.Println(&loginData)

}
