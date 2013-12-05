package godbtest

import (
	_ "github.com/lib/pq"
	"database/sql"
	"testing"
)

func Test存在する1件を削除(t *testing.T) {
	
	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 削除前の表示
	t.Log("削除前")
	selectWithQueryRow(t, db, id)

	// 削除
	deleteWithExec(t, db, id)

	// 削除後の表示
	t.Log("削除後")
	selectWithQueryRow(t, db, id)
}

func Test存在しない1件を削除(t *testing.T) {
	
	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	// 削除
	deleteWithExec(t, db, -1)

}

func Test複数件を削除(t *testing.T) {
	
	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 削除前の表示
	t.Log("削除前")
	selectWithQueryRow(t, db, id)
	selectWithQueryRow(t, db, id - 1)

	// 削除
	query := "delete from table1 where id in (?, ?)"

	result, err := db.Exec(query, id, id -1)

	if err != nil {
		t.Fatalf("クエリーの実行に失敗しました。: %v", err)
	}

	if c, err := result.LastInsertId(); err != nil {
		t.Errorf("LastInsertIdを取得できません。: %v", err)
	} else {
		t.Logf("LastInsertId: %v", c)
	}

	if c, err := result.RowsAffected(); err != nil {
		t.Errorf("RowsAffectedを取得できません。: %v", err)
	} else {
		t.Logf("RowsAffected: %v", c)
	}

	// 削除後の表示
	t.Log("削除後")
	selectWithQueryRow(t, db, id)
	selectWithQueryRow(t, db, id - 1)

}

func TestPreparedStatementによる複数件を削除(t *testing.T) {
	
	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 削除前の表示
	t.Log("削除前")
	for i := 0; i < 5; i++ {
		selectWithQueryRow(t, db, id - i)
	}

	// 削除
	query := "delete from table1 where id=?"

    stmt, err := db.Prepare(query)

    if err != nil {
    	t.Fatalf("Prepareに失敗しました。: ", err)
    }

    for i := 0; i < 5; i++ {

    	result, err := stmt.Exec(id - i)

		if err != nil {
			t.Fatalf("クエリーの実行に失敗しました。: %v", err)
		}

		if c, err := result.LastInsertId(); err != nil {
			t.Errorf("LastInsertIdを取得できません。: %v", err)
		} else {
			t.Logf("LastInsertId: %v", c)
		}

		if c, err := result.RowsAffected(); err != nil {
			t.Errorf("RowsAffectedを取得できません。: %v", err)
		} else {
			t.Logf("RowsAffected: %v", c)
		}

    }

    stmt.Close()

	// 削除後の表示
	t.Log("削除後")
	for i := 0; i < 5; i++ {
		selectWithQueryRow(t, db, id - i)
	}

}

func deleteWithExec(t *testing.T, db *sql.DB, id int) {

	query := "delete from table1 where id=?"

	result, err := db.Exec(query, id)

	if err != nil {
		t.Fatalf("クエリーの実行に失敗しました。: %v", err)
	}

	if c, err := result.LastInsertId(); err != nil {
		t.Errorf("LastInsertIdを取得できません。: %v", err)
	} else {
		t.Logf("LastInsertId: %v", c)
	}

	if c, err := result.RowsAffected(); err != nil {
		t.Errorf("RowsAffectedを取得できません。: %v", err)
	} else {
		t.Logf("RowsAffected: %v", c)
	}

}