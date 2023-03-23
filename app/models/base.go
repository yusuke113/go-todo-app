package models

import (
	"crypto/sha1"        // sha1ハッシュ関数を使うためのパッケージ
	"database/sql"       // SQLデータベースを扱うためのパッケージ
	"fmt"                // フォーマット関数を提供するパッケージ
	"go_todo_app/config" // 設定ファイルを読み込むためのパッケージ
	"log"                // ログ出力を行うためのパッケージ

	"github.com/google/uuid" // UUID生成用のパッケージ

	_ "github.com/mattn/go-sqlite3" // SQLite3を使うためのパッケージ
)

var Db *sql.DB // DB接続オブジェクト
var err error

const (
	tableNameUser = "users" // ユーザーテーブルの名前
	tableNameTodo = "todos" // TODOテーブルの名前
	tableNameSession = "sessions" // TODOテーブルの名前
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName) // データベースに接続
	if err != nil {                                                   // エラーがあればログに出力して終了
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser) // 作成日時

	Db.Exec(cmdU) // SQL文を実行
	
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT) // SQL文を実行

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) { // UUIDを生成する関数
	uuidobj, _ = uuid.NewUUID() // UUIDオブジェクトを生成
	return uuidobj              // UUIDを返す
}

func Encrypt(plaintext string) (cryptext string) { // パスワードをハッシュ化する関数
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext))) // パスワードをSHA1でハッシュ化
	return cryptext                                           // ハッシュ値を返す
}
