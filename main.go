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
	Src     string
	HTMLSrc template.HTML
	Time    time.Time
}

func readSource(filename string) string {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(source)
}

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/template000.html"))

	srcInfo := SourceHTML{}
	srcInfo.Title = "main.go"
	srcInfo.Src = readSource(srcInfo.Title)
	//srcInfo.Src = template.HTMLEscapeString(srcInfo.Src)
	//srcInfo.Src = strings.Replace(srcInfo.Src, "\n", "<br>", -1)
	srcInfo.HTMLSrc = template.HTML(srcInfo.Src)
	srcInfo.Time = time.Now()
	fmt.Print(srcInfo.Src)

	err := t.Execute(w, srcInfo)
	if err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func main() {
	http.HandleFunc("/index", htmlHandler0)

	// サーバーを起動
	http.ListenAndServe(":8989", nil)
}
