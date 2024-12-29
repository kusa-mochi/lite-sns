// DBの操作をラップした関数群。
// DBMS呼出の原子性を考慮した設計にはなっていないため、
// DB操作の競合については関数を呼び出す側で考慮した設計とすること。
package db_utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	log.Println("DB command:", prepare)
	log.Println("DB command params:", placeholders)

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

// whereConditions: プレースホルダ（$1, $2, ...）を含むWHERE句。
// whereParams: プレースホルダに渡すパラメータ。
func SelectFrom(db *sql.DB, keys []string, tableName string, whereConditions string, whereParams ...any) ([]any, error) {
	dbCommand := fmt.Sprintf(
		"SELECT %s FROM %s %s",
		strings.Join(keys, ","),
		tableName,
		whereConditions,
	)
	log.Println("DB command:", dbCommand)
	log.Println("DB command params:", whereParams)

	rows, err := db.Query(
		dbCommand,
		whereParams...,
	)
	if err != nil {
		if rows != nil {
			rows.Close()
		}
		return nil, fmt.Errorf("failed to run SELECT query | %s", err.Error())
	}
	defer rows.Close()

	log.Println("query done")

	ret := make([]any, 0)

	for rows.Next() {
		var p any
		err := rows.Scan(&p)
		if err != nil {
			return nil, fmt.Errorf("failed to scan returned records @ SelectFrom | %s", err.Error())
		}
		ret = append(ret, p)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("something error occured in rows object | %s", err.Error())
	}

	return ret, nil
}

// whereConditions: プレースホルダ（$1, $2, ...）を含むWHERE句。
// whereParams: プレースホルダに渡すパラメータ。
func DeleteFrom(db *sql.DB, tableName string, whereConditions string, whereParams ...any) error {
	dbCommand := fmt.Sprintf(
		"DELETE FROM %s %s",
		tableName,
		whereConditions,
	)
	log.Println("DB command:", dbCommand)
	log.Println("DB command params:", whereParams)

	_, err := db.Exec(
		dbCommand,
		whereParams...,
	)
	if err != nil {
		return fmt.Errorf("failed to execute DELETE command @ DeleteFrom | %s", err.Error())
	}

	return nil
}
