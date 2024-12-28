// DBの操作をラップした関数群。
// DBMS呼出の原子性を考慮した設計にはなっていないため、
// DB操作の競合については関数を呼び出す側で考慮した設計とすること。
package db_utils

import (
	"database/sql"
	"fmt"
	"log"
)

type KeyValuePair struct {
	Key   string
	Value any
}

func InsertInto(db *sql.DB, tableName string, pairs ...KeyValuePair) (int64, error) {
	if tableName == "" {
		return 0, fmt.Errorf("tableName must be specified")
	}
	if db == nil {
		return 0, fmt.Errorf("db must not be nil")
	}

	nPairs := len(pairs)
	if nPairs == 0 {
		log.Println("there are no data pairs @ InsertInto")
		return 0, nil
	}

	keys := ""
	placeholders := ""
	values := make([]any, 0)

	for i, pair := range pairs {
		keys += pair.Key
		placeholders += fmt.Sprintf("$%v", i+1)
		if i < nPairs-1 {
			keys += ", "
			placeholders += ", "
		}

		switch pair.Value.(type) {
		case string:
		case int64:
		case int:
		case bool:
		default:
			return 0, fmt.Errorf("invalid value type @ InsertInto")
		}

		values = append(values, pair.Value)
	}

	prepare := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tableName, keys, placeholders)

	stmt, err := db.Prepare(prepare)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare db execution @ InsertInto | %s", err.Error())
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		return 0, fmt.Errorf("failed to db execution @ InsertInfo | %s", err.Error())
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected @ InsertInfo | %s", err.Error())
	}

	return rowCnt, nil
}
