package server

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/Stolarskis/flood-alert-app/src/app-flood-server/server/handler"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Init() {
	s.Router = mux.NewRouter()
	s.setRouters()
}

func (s *Server) setRouters() {
	s.Get("/alerts", s.getAlerts)
	s.Get("/", s.home)
	s.Post("/addUser", s.addUser)
}

func (s *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("GET")
}
func (s *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("POST")
}

//Get request handlers
func (s *Server) getAlerts(w http.ResponseWriter, r *http.Request) {
	err := handler.CheckAlerts(w, r)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	handler.Home(w, r)
}

func (s *Server) addUser(w http.ResponseWriter, r *http.Request) {
	handler.AddUser(w, r)
}

func (a *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
