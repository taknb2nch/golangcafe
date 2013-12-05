package godbtest

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test1件追加_Execを使用(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	query := "insert into table1 (display_name, sex, birthday, age, married, rate, salary) "
    query += "values (?, ?, ?, ?, ?, ?, ?)"

    var r = createRecord()
    result, err := db.Exec(query, r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)

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

func TestPreparedStatementによる複数件追加(t *testing.T) {

	db, err := openConnection()

	if err != nil {
		t.Fatalf("データベースの接続に失敗しました。: %v", err)
	}

	defer db.Close()

	query := "insert into table1 (display_name, sex, birthday, age, married, rate, salary) "
    query += "values (?, ?, ?, ?, ?, ?, ?)"

    stmt, err := db.Prepare(query)

    if err != nil {
    	t.Fatalf("Prepareに失敗しました。: ", err)
    }

    for i := 0; i < 5; i++ {

		var r = createRecord()

    	result, err := stmt.Exec(r.displayName, r.sex, r.birthday, r.age, r.married, r.rate, r.salary)

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

}

func createRecord() Record {

	rec := Record{
		id : 0,
		displayName : sql.NullString { fmt.Sprintf("なまえ%03d", rand.Int63n(100)), true },
		sex : strconv.FormatInt(rand.Int63n(1) + 1, 10),
		birthday : mysql.NullTime { time.Now(), true },
		age : sql.NullInt64 { rand.Int63n(40) + 10, true },
		rate : sql.NullFloat64 { rand.Float64() / 100, true },
		salary : sql.NullInt64 { (rand.Int63n(90) + 10) * 1000, true },
	}

	if rand.Intn(1) == 0 {
		rec.married = sql.NullBool { false, true }
	} else {
		rec.married = sql.NullBool { true, true }
	}

	return rec
}