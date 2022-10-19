package model

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Password string `json:"password"`
}
