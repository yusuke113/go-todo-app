package models

import (
	"log"  // ログ出力を行うためのパッケージ
	"time" // 時間関連のパッケージ
)

type User struct { // ユーザー情報を保持する構造体
	ID        int       // ID
	UUID      string    // UUID
	Name      string    // 名前
	Email     string    // メールアドレス
	Password  string    // パスワード
	CreatedAt time.Time // 作成日時
}

func (u *User) CreatedUser() (err error) { // ユーザーを作成する関数
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)` // プレースホルダを使って値を設定

	_, err = Db.Exec(cmd, // SQL文を実行
		createUUID(),        // UUIDを生成
		u.Name,              // 名前
		u.Email,             // メールアドレス
		Encrypt(u.Password), // パスワードをハッシュ化
		time.Now())          // 現在時刻

	if err != nil { // エラーがあればログに出力して終了
		log.Fatalln(err)
	}
	return err // エラーを返す
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
			from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan( // 指定されたIDのユーザー情報を取得
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err // ユーザー情報とエラーを返す
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}