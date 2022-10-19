package apiserver

import (
	"database/sql"
	"net/http"

	"test_go/internal/app/store/sqlstore"
	"test_go/pkg/logging"

	_ "github.com/go-sql-driver/mysql"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)

	srv := NewServer(store)

	logging.GetLogger().Info("Listen...")
	return http.ListenAndServe(config.Listen.BindIp+":"+config.Listen.Port, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logging.GetLogger().Info("DB is ready!")
	return db, nil
}
