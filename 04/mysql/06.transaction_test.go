package godbtest

import (
	_ "github.com/go-sql-driver/mysql"
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
    query += "values (?, ?, ?, ?, ?, ?, ?)"

    var r = createRecord()
    var newId int

    result, err := tx.Exec(query, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)

	if err != nil {
		t.Fatalf("クエリーの実行に失敗しました。: %v", err)
	}

	if c, err := result.LastInsertId(); err != nil {
		t.Errorf("LastInsertIdを取得できません。: %v", err)
	} else {
		t.Logf("LastInsertId: %v", c)
		newId = int(c)
	}

	if c, err := result.RowsAffected(); err != nil {
		t.Errorf("RowsAffectedを取得できません。: %v", err)
	} else {
		t.Logf("RowsAffected: %v", c)
	}

	return newId, nil
}

func insertWithPreparedQueryRowTx(t *testing.T, tx *sql.Tx, count int) ([]int, error) {

	query := "insert into table1 (display_name, sex, birthday, age, married, rate, salary) "
    query += "values (?, ?, ?, ?, ?, ?, ?)"

    stmt, err := tx.Prepare(query)

    if err != nil {
    	t.Fatalf("Prepareに失敗しました。: ", err)
    }

    newIds := make([]int, count)

    for i := 0; i < count; i++ {

    	var r = createRecord()

    	result, err := stmt.Exec(r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)

		if err != nil {
			t.Fatalf("クエリーの実行に失敗しました。: %v", err)
		}

		if c, err := result.LastInsertId(); err != nil {
			t.Errorf("LastInsertIdを取得できません。: %v", err)
		} else {
			t.Logf("LastInsertId: %v", c)
			newIds[i] = int(c)
		}

		if c, err := result.RowsAffected(); err != nil {
			t.Errorf("RowsAffectedを取得できません。: %v", err)
		} else {
			t.Logf("RowsAffected: %v", c)
		}

    }

    stmt.Close()

    return newIds, nil

}