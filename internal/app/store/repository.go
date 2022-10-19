package store

import "test_go/internal/app/model"

type Repository interface {
	Create(*model.Employee) error

	GetUsers() ([]model.Employee, error)

	FindUser(*model.Employee) (*model.Employee, error)
}
