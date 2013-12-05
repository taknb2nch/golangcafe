package godbtest

import (
	_ "github.com/lib/pq"
	"database/sql"
	"testing"
)

func Testトランザクションを使用した1件追加_コミット(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		t.Fatalf("トランザクションの取得に失敗しました。: %v", err)
	}

	newId, err := insertWithQueryRowTx(t, tx)

	// テストなので常にコミットします。
	err = tx.Commit()

	if err != nil {
		t.Fatalf("トランザクションのコミットに失敗しました。: %v", err)
	} else {
		t.Logf("トランザクションをコミットしました。")
	}

	//
	selectWithQueryRow(t, db, newId)

}

func Testトランザクションを使用した1件追加_ロールバック(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		t.Fatalf("トランザクションの取得に失敗しました。: %v", err)
	}

	newId, err := insertWithQueryRowTx(t, tx)

	// テストなので常にロールバックします。
	err = tx.Rollback()

	if err != nil {
		t.Fatalf("トランザクションのロールバックに失敗しました。: %v", err)
	} else {
		t.Logf("トランザクションをロールバックしました。")
	}

	//
	selectWithQueryRow(t, db, newId)

}

func TestトランザクションとPreparedStatementを使用した複数件追加_コミット(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	count := 5

	tx, err := db.Begin()

	if err != nil {
		t.Fatalf("トランザクションの取得に失敗しました。: %v", err)
	}

	newIds, err := insertWithPreparedQueryRowTx(t, tx, count)

	// テストなので常にコミットします。
	err = tx.Commit()

	if err != nil {
		t.Fatalf("トランザクションのコミットに失敗しました。: %v", err)
	} else {
		t.Logf("トランザクションをコミットしました。")
	}

	//
	for _, newId := range newIds {
		selectWithQueryRow(t, db, newId)
	}

}

func TestトランザクションPreparedStatementを使用した複数件追加_ロールバック(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	count := 5

	tx, err := db.Begin()

	if err != nil {
		t.Fatalf("トランザクションの取得に失敗しました。: %v", err)
	}

	newIds, err := insertWithPreparedQueryRowTx(t, tx, count)

	// テストなので常にロールバックします。
	err = tx.Rollback()

	if err != nil {
		t.Fatalf("トランザクションのロールバックに失敗しました。: %v", err)
	} else {
		t.Logf("トランザクションをロールバックしました。")
	}

	//
	for _, newId := range newIds {
		selectWithQueryRow(t, db, newId)
	}

}

func insertWithQueryRowTx(t *testing.T, tx *sql.Tx) (int, error) {

	query := "insert into table1 (display_name, sex, birthday, age, married, rate, salary) "
    query += "values ($1, $2, $3, $4, $5, $6, $7) returning id"

    var r = createRecord()
    var newId int

    err := tx.QueryRow(query, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary).Scan(&newId)

	switch {
	case err == sql.ErrNoRows : 
		// TODO:
		t.Logf("対象のレコードは存在しません。: %v", err)
	case err != nil :
		//t.Fatalf("値の取得に失敗しました。: %v", err)
		return 0, err
	default :
		t.Logf("登録ID=%d", newId)
	}

	return newId, nil
}

func insertWithPreparedQueryRowTx(t *testing.T, tx *sql.Tx, count int) ([]int, error) {

	query := "insert into table1 (display_name, sex, birthday, age, married, rate, salary) "
    query += "values ($1, $2, $3, $4, $5, $6, $7) returning id"

    stmt, err := tx.Prepare(query)

    if err != nil {
    	t.Fatalf("Prepareに失敗しました。: ", err)
    }

    newIds := make([]int, count)

    for i := 0; i < count; i++ {

    	var r = createRecord()
    	var newId int

    	err := stmt.QueryRow(r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary).Scan(&newId)

		switch {
		case err == sql.ErrNoRows : 
			// TODO:
			t.Logf("対象のレコードは存在しません。: %v", err)
		case err != nil :
			return nil, err
		default :
			t.Logf("登録ID=%d", newId)
			newIds[i] = newId
		}

    }

    stmt.Close()

    return newIds, nil

}