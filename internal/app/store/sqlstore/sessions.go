package sqlstore

import (
	"fmt"
	"test_go/internal/app/model"
)

type Sessions struct {
	store *Store
}

func (session *Sessions) Create(s *model.Session) error {
	query := fmt.Sprintf("INSERT INTO Sessions (userId, status, refreshToken, timeClose) VALUES('%d', '%s', '%s', '%d')",
		s.UserId,
		s.Status,
		s.RefreshToken,
		s.TimeClose,
	)
	_, err := session.store.db.Exec(query)
	return err
}

func (session *Sessions) Delete(refreshToken string) error {
	query := fmt.Sprintf("DELETE FROM Sessions WHERE refreshToken='%s';",
		refreshToken,
	)
	_, err := session.store.db.Exec(query)
	return err
}
