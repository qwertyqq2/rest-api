package apiserver

import (
	"test_go/internal/app/handlers/users"
	"test_go/internal/app/store"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httprouter.Router
	store  store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		router: httprouter.New(),
		store:  store,
	}

	s.conifigureServer()

	return s
}

func (s *Server) conifigureServer() {
	users.New(s.store).Register(s.router)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
