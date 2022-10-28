package sqlstore

import (
	"fmt"
	"log"
	"test_go/internal/app/model"
)

type Repository struct {
	store *Store
}

func (r *Repository) Create(e *model.Employee) error {
	query := fmt.Sprintf("INSERT INTO Employees (name, status, password) VALUES('%s', '%s', '%s') RETURNING id;",
		e.Name,
		e.Status,
		e.Password,
	)
	return r.store.db.QueryRow(query).Scan(&e.ID)
}

func (r *Repository) GetUsers() ([]model.Employee, error) {
	query := "SELECT * FROM Employees"
	resp, err := r.store.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var users []model.Employee
	for resp.Next() {
		var e model.Employee
		if err := resp.Scan(&e.ID, &e.Name, &e.Status, &e.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, e)
	}
	return users, err
}

func (r *Repository) FindUser(req *model.Employee) (*model.Employee, error) {
	query := fmt.Sprintf("SELECT * FROM Employees WHERE name='%s' AND password='%s'", req.Name, req.Password)

	u := &model.Employee{}

	if err := r.store.db.QueryRow(query).Scan(
		&u.ID,
		&u.Name,
		&u.Status,
		&u.Password,
	); err != nil {
		return nil, err
	}

	return u, nil

}
