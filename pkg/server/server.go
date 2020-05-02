package server

import "github.com/phuwn/smart-schedule/pkg/store"

var srv Server

type Server struct {
	store *store.Store
}

func NewServerCfg(store *store.Store) {
	srv.store = store
}

func GetServerCfg() *Server {
	return &srv
}

func (s *Server) Store() *store.Store {
	return s.store
}
