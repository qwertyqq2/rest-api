package sqlstore

import (
	"database/sql"

	"test_go/internal/app/store"
)

type Store struct {
	db       *sql.DB
	rep      *Repository
	sessions *Sessions
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Repository() store.Repository {
	if s.rep != nil {
		return s.rep
	}

	s.rep = &Repository{
		store: s,
	}

	return s.rep
}

func (s *Store) Sessions() store.Sessions {
	if s.sessions != nil {
		return s.sessions
	}

	s.sessions = &Sessions{
		store: s,
	}

	return s.sessions
}

//Store -> UserRepository -> NewUser()
