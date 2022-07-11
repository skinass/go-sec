package main

import (
	"fmt"
	// "html/template"
	"math/rand"
	"net/http"
	"text/template"
)

var messages = []string{"Hello World"}

var messagesTmpl = `<html><body>
	<br />
	<br />

	<form action="/add_comment" method="post">
		<textarea name="comment"></textarea><br />
		<input type="submit" value="Comment">
	</form>

	<br />

    {{range .Messages}}
		<div style="border: 1px solid black; padding: 5px; margin: 5px;">
			<!-- text/template по-умолч ничего не экранируется, надо указать | html --> 
			<!-- html/template по-умолч будет экранировать --> 

			{{.}}
		</div>
    {{end}}
</body></html>`

func main() {
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(messagesTmpl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct {
			Messages []string
		}{
			Messages: messages,
		})
	})

	http.HandleFunc("/add_comment", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		commentText := r.FormValue("comment")
		messages = append(messages, commentText)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/clear_comments", func(w http.ResponseWriter, r *http.Request) {
		messages = []string{}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// <h1 style="color:blue;">This is a heading</h1>
