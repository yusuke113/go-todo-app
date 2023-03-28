package controllers

import (
	"fmt"
	"go_todo_app/app/models"
	"go_todo_app/config"
	"net/http"
	"regexp"
	"strconv"
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
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)")

// ParseUrlは、HTTPリクエストに対するURLのパースとバリデーションを行い、
// ハンドラ関数をラップして返します。
//
// 引数:
//   - fn: HTTPハンドラ関数。3つの引数を受け取り、何も返しません。
//     第1引数にはResponseWriter、第2引数にはRequest、第3引数にはIDが渡されます。
//
// 戻り値:
//   - http.HandlerFunc: ラップされたHTTPハンドラ関数。
//
// 動作:
//   - リクエストURLを正規表現でバリデーションします。
//   - URLが正しい場合、第3セグメントを数値に変換して第3引数に渡します。
//   - ハンドラ関数を実行します。
//   - URLが不正な場合、404 Not Foundを返します。
func parseUrl(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
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
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseUrl(todoEdit))
	http.HandleFunc("/todos/update/", parseUrl(todoUpdate))
	http.HandleFunc("/todos/delete/", parseUrl(todoDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
