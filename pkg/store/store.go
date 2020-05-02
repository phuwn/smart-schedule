package store

import "github.com/phuwn/smart-schedule/pkg/store/list"

// Store - server store struct
type Store struct {
	List list.Store
}

// New - create new store variable
func New() *Store {
	return &Store{
		List: list.NewStore(),
	}
}
