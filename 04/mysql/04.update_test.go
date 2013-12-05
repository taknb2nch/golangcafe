package godbtest

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"testing"
)

func Test存在する1件を更新(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 更新前の表示
	t.Log("更新前")
	selectWithQueryRow(t, db, id)

	// 更新
	updateWithExec(t, db, id)

	// 更新後の表示
	t.Log("更新後")
	selectWithQueryRow(t, db, id)
}

func Test存在しない1件を更新(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	//updateWithQueryRow(t, db, -1)
	updateWithExec(t, db, -1)
}

func Test複数件を更新(t *testing.T) {
	
	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 更新前の表示
	t.Log("更新前")
	selectWithQueryRow(t, db, id)
	selectWithQueryRow(t, db, id - 1)

	query := "update table1 set display_name=?, sex=?, birthday=?, age=?, married=?, rate=?, salary=? "
	query += "where id in (?, ?)"

	result, err := db.Exec(query, nil, 0, nil, nil, nil, nil, nil, id, id - 1)

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

	// 更新後の表示
	t.Log("更新後")
	selectWithQueryRow(t, db, id)
	selectWithQueryRow(t, db, id - 1)

}

func TestPreparedStatementによる複数件を更新(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	id := getMaxId(t, db)

	// 更新前の表示
	t.Log("更新前")
	for i := 0; i < 5; i++ {
		selectWithQueryRow(t, db, id - i)
	}

	query := "update table1 set display_name=?, sex=?, birthday=?, age=?, married=?, rate=?, salary=? "
	query += "where id=?"

    stmt, err := db.Prepare(query)

    if err != nil {
    	t.Fatalf("Prepareに失敗しました。: ", err)
    }

    for i := 0; i < 5; i++ {

    	result, err := stmt.Exec(nil, 0, nil, nil, nil, nil, nil, id - i)

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

    // 更新後の表示
	t.Log("更新後")
	for i := 0; i < 5; i++ {
		selectWithQueryRow(t, db, id - i)
	}

}

func updateWithExec(t *testing.T, db *sql.DB, id int) {

	query := "update table1 set display_name=?, sex=?, birthday=?, age=?, married=?, rate=?, salary=? "
	query += "where id=?"

	result, err := db.Exec(query, nil, 0, nil, nil, nil, nil, nil, id)

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

func getMaxId(t *testing.T, db *sql.DB) int {

	query := "select max(id) as max_id from table1"

	var maxId int

	err := db.QueryRow(query).Scan(&maxId)

	switch {
	case err == sql.ErrNoRows : 
		return 0
	case err != nil :
		t.Fatalf("値の取得に失敗しました。: %v", err)
	default :
		return maxId
	}

	return 0;
}