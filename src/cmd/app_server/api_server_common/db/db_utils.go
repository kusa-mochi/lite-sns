// DBの操作をラップした関数群。
// DBMS呼出の原子性を考慮した設計にはなっていないため、
// DB操作の競合については関数を呼び出す側で考慮した設計とすること。
package db_utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
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
	log.Println("DB command params:", values)

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
func SelectFrom(db *sql.DB, keys []string, tableName string, whereConditions string, whereParams ...any) ([](map[string]any), error) {
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

	ret := make([](map[string]any), 0)
	nCols := len(keys)

	for rows.Next() {
		// cols := make([]any, nCols)
		// for i := 0; i < nCols; i++ {
		// 	cols[i] =
		// }
		// err := rows.Scan(cols...)
		// log.Println("cols:")
		// log.Println(cols...)

		var er error = nil
		var buf [10]any

		if nCols == 1 {
			er = rows.Scan(&buf[0])
		} else if nCols == 2 {
			er = rows.Scan(&buf[0], &buf[1])
		} else if nCols == 3 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2])
		} else if nCols == 4 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3])
		} else if nCols == 5 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4])
		} else if nCols == 6 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4], &buf[5])
		} else if nCols == 7 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4], &buf[5], &buf[6])
		} else if nCols == 8 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4], &buf[5], &buf[6], &buf[7])
		} else if nCols == 9 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4], &buf[5], &buf[6], &buf[7], &buf[8])
		} else if nCols == 10 {
			er = rows.Scan(&buf[0], &buf[1], &buf[2], &buf[3], &buf[4], &buf[5], &buf[6], &buf[7], &buf[8], &buf[9])
		}

		if er != nil {
			return nil, fmt.Errorf("failed to scan returned records @ SelectFrom | %s", er.Error())
		}

		row := make(map[string]any)
		for i := 0; i < nCols; i++ {
			row[keys[i]] = buf[i]
		}

		ret = append(ret, row)
		// ret = append(ret, cols)
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

func PrepareAndExec(db *sql.DB, prepare string, params ...any) (int64, error) {
	stmt, err := db.Prepare(prepare)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare db execution command: %s (%v) @ PrepareAndExec | %s", prepare, params, err.Error())
	}
	res, err := stmt.Exec(params...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute db command @ PrepareAndExec | %s", err.Error())
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected @ PrepareAndExec | %s", err.Error())
	}

	return rowCnt, nil
}

func Query(db *sql.DB, query string, params ...any) ([][]any, error) {
	rows, err := db.Query(query, params...)
	if err != nil {
		if rows != nil {
			rows.Close()
		}
		return nil, fmt.Errorf("failed to run query @ Query | %s", err.Error())
	}

	log.Println("query done")

	ret := make([][]any, 0)
	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns @ Query | %s", err.Error())
	}

	nCols := len(cols)
	buf := make([]any, nCols)
	pBuf := make([]any, nCols)
	for i := 0; i < nCols; i++ {
		pBuf[i] = &buf[i]
	}

	for rows.Next() {
		err := rows.Scan(pBuf...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan returned records @ Query | %s", err.Error())
		}

		ret = append(ret, buf)
	}

	return ret, nil
}

// UnixTimeToString: Unix時間を文字列に変換する。
// Unix時間は秒単位で表現される。
// 戻り値: "2006-01-02 15:04:05"形式の文字列。
func UnixTimeToString(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2006-01-02 15:04:05")
}
