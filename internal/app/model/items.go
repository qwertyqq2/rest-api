package model

type Item struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	OwnerID int    `json:"ownerId"`
}
