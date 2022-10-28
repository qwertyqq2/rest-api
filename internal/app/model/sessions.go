package model

type Session struct {
	UserId       int    `json:"userId"`
	Status       string `json:"status"`
	RefreshToken string `json:"refreshToken"`
	TimeClose    int    `json:"timeClose"`
}
