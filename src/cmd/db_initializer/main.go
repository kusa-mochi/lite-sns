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
	ColType_Hash256
	ColType_Integer
	ColType_Interval
	ColType_JWT
	ColType_TimestampWithTimezone
)

// ユーザーアカウントに設定するアイコンの種別
// 列挙型のように定義しているものの、整数値ではなく文字列として定義しているのは、
// 後から種別の追加等の必要が生じた場合に値が変更されアイコンとの対応が
// おかしくなってしまう事態を回避するため。
type IconType string

const (
	IconType_Default = IconType("IconType_Default")
	IconType_Male0   = IconType("IconType_Male0")
	IconType_Male1   = IconType("IconType_Male1")
	IconType_Male2   = IconType("IconType_Male2")
	IconType_Female0 = IconType("IconType_Female0")
	IconType_Female1 = IconType("IconType_Female1")
	IconType_Female2 = IconType("IconType_Female2")
)

// いいね、共感、お気に入りの対象
type ActionTarget int

const (
	ActionTarget_Post = iota
	ActionTarget_Comment
)

const (
	queryLogPrefix string = "[query]"
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
	// テーブルを作成するクエリ。
	// ここではテーブルが既に存在するケースを想定せず、存在する場合はエラー終了して欲しいため、
	// あえて if not exists 指定はしない。
	query := "create table "
	query += tableAttr.Name + " "
	query += "(id serial primary key, "
	for idx, col := range tableAttr.Cols {
		query += col.Name + " "
		switch col.Type {
		case ColType_Boolean:
			query += "boolean"
		case ColType_CharacterVarying:
			query += fmt.Sprintf("varchar(%v)", col.MaxLength)
		case ColType_Hash256:
			query += "varchar(64)" // ハッシュ値の16進数文字列
		case ColType_Integer:
			query += "integer"
		case ColType_Interval:
			query += "interval"
		case ColType_JWT:
			query += "varchar(256)"
		case ColType_TimestampWithTimezone:
			query += "bigint" // unix time (64bit符号付整数) の数値を格納する。
		}

		if col.IsNullable {
			query += " null"
		} else { // lastId, err := res.LastInsertId()
			// if err != nil {
			// 	log.Fatalln(err)
			// }

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

func AddTestRecords(db *sql.DB) {
	// TODO
	query := ""
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
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
							Name:                "icon_type",
							Type:                ColType_CharacterVarying,
							MaxLength:           128,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "icon_background_color",
							Type:                ColType_CharacterVarying,
							MaxLength:           6,
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
							Type:                ColType_Hash256,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "access_token_secret_key",
							Type:                ColType_Hash256,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "post",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
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
				{
					Name: "comment",
					Cols: []ColAttr{
						{
							Name:                "post_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
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
				{
					Name: "good",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "action_target",
							Type:                ColType_Integer, // 0: post, 1: comment
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_id", // postまたはcommentのid。どちらのidなのかはaction_targetで判別する。
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id", // いいねされたpostまたはcommentの作者のuser_id
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "datetime",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "empathy",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "action_target",
							Type:                ColType_Integer, // 0: post, 1: comment
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_id", // postまたはcommentのid。どちらのidなのかはaction_targetで判別する。
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id", // 共感されたpostまたはcommentの作者のuser_id
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "datetime",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "favorite",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "action_target",
							Type:                ColType_Integer, // 0: post, 1: comment
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_id", // postまたはcommentのid。どちらのidなのかはaction_targetで判別する。
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id", // お気に入り登録されたpostまたはcommentの作者のuser_id
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "datetime",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "follow",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "follow_at",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "block",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "block_at",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "mute",
					Cols: []ColAttr{
						{
							Name:                "user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "target_user_id",
							Type:                ColType_Integer,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "mute_at",
							Type:                ColType_TimestampWithTimezone,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
					},
				},
				{
					Name: "signup_access_token",
					Cols: []ColAttr{
						{
							Name:                "access_token",
							Type:                ColType_JWT,
							MaxLength:           0,
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
							Name:                "nickname",
							Type:                ColType_CharacterVarying,
							MaxLength:           20,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "password_hash",
							Type:                ColType_Hash256,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "secret_key",
							Type:                ColType_Hash256,
							MaxLength:           0,
							IsNullable:          false,
							IsAutoIncrementable: false,
						},
						{
							Name:                "expiration_datetime",
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

	// // デバッグ上必要なら、以下のコードでテスト用レコードを各テーブルに追加する。
	// AddTestRecords()
}
