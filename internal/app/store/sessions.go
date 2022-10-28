package store

import "test_go/internal/app/model"

type Sessions interface {
	Create(*model.Session) error
	Delete(string) error
}
