package godbtest

import (
	"github.com/lib/pq"
	"database/sql"
	"testing"
)

type Record struct {
	id int
	displayName sql.NullString
	sex string
	birthday pq.NullTime
	age sql.NullInt64
	married sql.NullBool
	rate sql.NullFloat64
	salary sql.NullInt64
}

func Test複数件を取得(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	selectWithQuery(t, db, 1)
}

func Test複数件を取得_結果0件(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	selectWithQuery(t, db, 100)
}

func Test存在する1件を取得(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	selectWithQueryRow(t, db, 1)
}

func Test存在しない1件を取得(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	selectWithQueryRow(t, db, -1)
}

func Test複数件ヒットする条件でQueryRowを使用した場合(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	query := "select id, display_name, sex, birthday, age, married, rate, salary from table1"

	var r Record
	
	err = db.QueryRow(query).Scan(&r.id, &r.displayName, &r.sex, &r.birthday, &r.age, &r.married, &r.rate, &r.salary)

	switch {
	case err == sql.ErrNoRows : 
		t.Logf("対象のレコードは存在しません。%v", err)
	case err != nil :
		t.Fatalf("値の取得に失敗しました。: %v", err)
	default :
		t.Logf("%d, %v, %v, %v, %v, %v, %v, %v", 
			r.id, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)
	}

}

func selectWithQuery(t *testing.T, db *sql.DB, key int) {

	query := "select id, display_name, sex, birthday, age, married, rate, salary from table1 where id>=$1"

	rows, err := db.Query(query, key)

	if err != nil {
		t.Fatalf("クエリーの実行に失敗しました。: %v", err)
	}

	for rows.Next() {

		var r Record

		if err = rows.Scan(&r.id, &r.displayName, &r.sex, &r.birthday, &r.age, &r.married, &r.rate, &r.salary); err != nil {
			t.Errorf("値の取得に失敗しました。: %v", err)
		}

		t.Logf("%d, %v, %v, %v, %v, %v, %v, %v", 
			r.id, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)
	}

	rows.Close()

}

func selectWithQueryRow(t *testing.T, db *sql.DB, key int) {

	query := "select id, display_name, sex, birthday, age, married, rate, salary from table1 where id=$1"

	var r Record
	
	err := db.QueryRow(query, key).Scan(&r.id, &r.displayName, &r.sex, &r.birthday, &r.age, &r.married, &r.rate, &r.salary)

	switch {
	case err == sql.ErrNoRows : 
		t.Logf("対象のレコードは存在しません。id=%d : %v", key, err)
	case err != nil :
		t.Fatalf("値の取得に失敗しました。: %v", err)
	default :
		t.Logf("%d, %v, %v, %v, %v, %v, %v, %v", 
			r.id, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)
	}

}

func Testカラム名の列挙(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	query := "select id, display_name, sex, birthday, age, married, rate, salary from table1 where id>=$1"

	rows, err := db.Query(query, 1)

	if err != nil {
		t.Fatalf("クエリーの実行に失敗しました。: %v", err)
	}

	names, err := rows.Columns()

	if err != nil {
		t.Fatalf("カラム名の取得に失敗しました。: %v", err)
	}

	for name := range names {
		t.Logf("%v", name)
	}

	rows.Close()

}

