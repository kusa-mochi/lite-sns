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

func main() {
	var (
		dbPort = flag.Int("db_port", 5432, "port number of the db server")
		dbAttr = DbAttr{
			Name: "lite_sns_db",
			Tables: []TableAttr{
				{
					Name: "user",
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

	db, err := sql.Open("postgres", fmt.Sprintf("host=lite-sns-db port=%v user=user password=postgres sslmode=disable", *dbPort))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	log.Println("postgres opened")

	// DBを作成する。
	stmt, err := db.Prepare(fmt.Sprintf("create database %s", dbAttr.Name))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}
