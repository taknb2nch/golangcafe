package godbtest

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"testing"
)

func Testデータベースへの接続(t *testing.T) {
	
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/godbtest")

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
	return sql.Open("mysql", "root:mysql@tcp(localhost:3306)/godbtest")
}