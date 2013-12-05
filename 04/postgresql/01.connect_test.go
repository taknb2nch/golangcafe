package godbtest

import (
	_ "github.com/lib/pq"
	"database/sql"
	"testing"
)

func Testデータベースへの接続(t *testing.T) {
	
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost dbname=godbtest sslmode=disable")

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。2: %v", err)
	}

	defer db.Close()

	t.Log("成功")

}

func Testデータベースへの接続_URL(t *testing.T) {
	
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/godbtest?sslmode=disable")

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。2: %v", err)
	}

	defer db.Close()

	t.Log("成功")

}

func openConnection() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost/godbtest?sslmode=disable")
}