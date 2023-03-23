package controllers

import (
	"fmt"
	"go_todo_app/app/models"
	"go_todo_app/config"
	"net/http"
	"text/template"
)

// generateHTML はHTMLを生成する関数です。
// 引数として http.ResponseWriter, data interface{}, filenames ...string をとります。
// filenamesの中身はapp/views/templates/以下に存在するHTMLテンプレートのファイル名です。
// templates変数にHTMLテンプレートを読み込み、ExecuteTemplateを使ってHTMLを生成しています。
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// session はセッション情報を取得する関数です。
// 引数として http.ResponseWriter, *http.Request をとり、models.Session型の変数sessと、error型のerrを返します。
// クッキーからセッション情報を取得し、セッション情報が無効でないかをチェックしています。
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

// StartMainServer はメインのWebサーバーを開始する関数です。
// config.Config.Staticに設定された静的ファイルをハンドリングし、
// "/"、"/signup"、"/login"、"/authenticate"の各パスに対する処理を定義しています。
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
