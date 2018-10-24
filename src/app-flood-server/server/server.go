package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Initalize() {
	s.Router = mux.NewRouter()
	s.setRouters()
}

func (s *Server) setRouters() {

}

func (s *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.Router.HandleFunc(path, f).Methods("GET")
}

//Get request handlers
func (s *Server) getAlerts(w http.ResponseWriter, r *http.Request) {
	handler.
}
