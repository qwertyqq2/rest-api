package sqlstore

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRepository(t *testing.T) {
	db, err := sql.Open("mysql", "root:aaa@tcp(127.0.0.1:3306)/MarketDB")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	store := New(db)

	store.Repository()

	t.Log("All right")

}
