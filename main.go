package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// SourceHTML ソースをHTML表示するためのもの
type SourceHTML struct {
	Title   string
	Param   string
	HTMLSrc template.HTML
	Time    time.Time
}

// readSource ソース読み込み
func readSource(filename string) string {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(source)
}

// ServerHandler サーバーハンドラ
type ServerHandler struct{}

// ServerHTTP 受け口
func (h *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/template000.html"))

	// ソースを読み込み表示する
	srcInfo := SourceHTML{}
	srcInfo.Title = "main.go"
	src := readSource(srcInfo.Title)

	srcInfo.HTMLSrc = template.HTML(src)
	srcInfo.Time = time.Now()
	srcInfo.Param = fmt.Sprintf("w = %v r = %v", w, r)

	//fmt.Print(src)

	err := t.Execute(w, srcInfo)
	if err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func main() {
	// 自前のHTTPハンドラ作成
	handler := ServerHandler{}

	// サーバー作成
	server := http.Server{
		Addr:    "localhost:8989",
		Handler: &handler,
	}

	server.ListenAndServe()
}
