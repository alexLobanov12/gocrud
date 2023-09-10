package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
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