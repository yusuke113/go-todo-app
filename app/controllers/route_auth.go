// package controllers はHTTPリクエストを処理するための関数を定義するパッケージです。
// 各関数はリクエストに応じた処理を行い、レスポンスを生成します。
package controllers

import (
	"go_todo_app/app/models"
	"log"
	"net/http"
)

// signupはユーザー登録ページのGETリクエストと、ユーザー情報を登録するPOSTリクエストを処理します。
// GETリクエストの場合、セッションが存在しなければsignupページを表示します。
// POSTリクエストの場合、リクエストフォームから取得した情報を元に、新しいユーザーを作成し、ルートページにリダイレクトします。
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

// loginはログインページのGETリクエストを処理します。
// セッションが存在しなければloginページを表示します。
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// authenticateはログインフォームから送信された情報を処理します。
// 入力されたメールアドレスに対応するユーザーが存在すれば、入力されたパスワードとハッシュ化されたパスワードが一致するか確認します。
// 認証に成功した場合、新しいセッションを作成し、cookieに保存します。
// 認証に失敗した場合、ログインページにリダイレクトします。
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// logoutはユーザーのセッションを削除してログアウトするための関数。
// リクエストから_cookieを取得して、そのUUIDを持つセッションを削除します。
// エラーが発生した場合はログに記録しますが、エラーがCookieの無効な場合は何もしません。
// 最後に、ログインページにリダイレクトします。
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}

	http.Redirect(w, r, "/login", 302)
}
