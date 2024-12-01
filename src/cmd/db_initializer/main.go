package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQL DB で用いるデータ型
type ColType int

const (
	ColType_Boolean = iota
	ColType_CharacterVarying
	ColType_Integer
	ColType_Interval
	ColType_TimestampWithTimezone
)

const (
	queryLogPrefix string = "*** query ***"
)

// id 以外の列定義に用いるパラメータを格納するための構造体
// id はすべてのテーブルで固定のパラメータ値を用いるため、この構造体は用いない。
type ColAttr struct {
	Name                string
	Type                ColType
	MaxLength           int
	IsNullable          bool
	IsAutoIncrementable bool
}

type TableAttr struct {
	Name string
	Cols []ColAttr
}

type DbAttr struct {
	Name   string
	Tables []TableAttr
}

func CreateDatabase(db *sql.DB, dbName string) {
	query := fmt.Sprintf("create database %s", dbName)
	log.Println(queryLogPrefix, query)

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateTable(db *sql.DB, tableAttr *TableAttr) {
	query := "create table "
	query += tableAttr.Name + " "
	query += "("
	for idx, col := range tableAttr.Cols {
		query += col.Name + " "
		switch col.Type {
		case ColType_Boolean:
			query += "boolean"
		case ColType_CharacterVarying:
			query += "varchar("
			query += fmt.Sprintf("%v", col.MaxLength)
			query += ")"
		case ColType_Integer:
			query += "integer"
		case ColType_Interval:
			query += "interval"
		case ColType_TimestampWithTimezone:
			query += "timestamptz"
		}

		if col.IsNullable {
			query += " null"
		} else {
			query += " not null"
		}

		if idx != len(tableAttr.Cols)-1 {
			query += ", "
		}
	}
	query += ")"
	log.Println(queryLogPrefix, query)

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var (
		dbPort = flag.Int("db_port", 5432, "port number of the db server")
		dbAttr = DbAttr{
			Name: "lite_sns_db",
			Tables: []TableAttr{
				{
					Name: "sns_user",
					Cols: []ColAttr{
						{
							Name:                "name",
							Type:                ColType_CharacterVarying,
							MaxLength:           20,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "email_address",
							Type:                ColType_CharacterVarying,
							MaxLength:           254,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "password_hash",
							Type:                ColType_CharacterVarying,
							MaxLength:           64, // RS256 ハッシュを16進数の文字列で表現したもの。
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "access_token_secret_key",
							Type:                ColType_CharacterVarying,
							MaxLength:           64, // RS256 ハッシュを16進数の文字列で表現したもの。
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "post",
					Cols: []ColAttr{
						{
							Name:                "text",
							Type:                ColType_CharacterVarying,
							MaxLength:           1000,
							IsNullable:          true,
							IsAutoIncrementable: false,
						},
						{
							Name:                "created_at",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "updated_at",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
			},
		}
	)

	// DB名指定無しで接続する。
	db, err := sql.Open("postgres", fmt.Sprintf("host=lite-sns-db port=%v user=user password=postgres sslmode=disable", *dbPort))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("postgres opened")

	// lite-sns専用DBを作成する。
	CreateDatabase(db, dbAttr.Name)

	// 一旦接続を切る。
	db.Close()

	// DB名を指定して接続し直す。
	db, err = sql.Open("postgres", fmt.Sprintf("host=lite-sns-db port=%v user=user password=postgres dbname=%s sslmode=disable", *dbPort, dbAttr.Name))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Printf("connected to %s", dbAttr.Name)

	// テーブルを作成する。
	for _, tableAttr := range dbAttr.Tables {
		CreateTable(db, &tableAttr)
	}
}
