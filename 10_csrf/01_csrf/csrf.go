package main

import (
	"net/http"
	// "html/template"
	"fmt"
	"strconv"
	"text/template" // надо заменить text/template на html/template чтобы по-умоллчанию было правильное экранирование
)

var cnt = 1

type Msg struct {
	ID      int
	Message string
	Rating  int
}

var messages = map[int]*Msg{}

var messagesTmpl = `<html>
<head>
<script>
	function rateComment(id, vote) {
		var request = new XMLHttpRequest();
		request.open('POST', '/rate?id='+id+"&vote="+(vote > 0 ? "up" : "down"), true);

		request.onload = function() {
		    var resp = JSON.parse(request.responseText);
			console.log(resp, resp.id)
			console.log('#rating-'+resp.id)
			console.log(document.querySelector('#rating-'+resp.id))
			document.querySelector('#rating-'+resp.id).innerHTML = resp.rating;
		};
		request.send();
	}
</script>
</head>
<body>
	<br />
	<form action="/comment" method="post">
		<textarea name="comment"></textarea><br />
		<input type="submit" value="Comment">
	</form>
	<br />
	
    {{range $idx, $var := .Messages}}
		<div style="border: 1px solid black; padding: 5px; margin: 5px;">
			<button onclick="rateComment({{$var.ID}}, 1)">&uarr;</button>
			<span id="rating-{{$var.ID}}">{{$var.Rating}}</span>
			<button onclick="rateComment({{$var.ID}}, -1)">&darr;</button>
			&nbsp;
			{{$var.Message}}
		</div>
    {{end}}
</body></html>`

func main() {

	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(messagesTmpl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct {
			Messages map[int]*Msg
		}{
			Messages: messages,
		})
	})

	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		commentText := r.FormValue("comment")
		id := cnt
		messages[id] = &Msg{
			ID:      id,
			Message: commentText,
			Rating:  0,
		}
		cnt++
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/rate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		vote := r.URL.Query().Get("vote")

		if msg, ok := messages[id]; ok {
			if vote == "up" {
				msg.Rating++
			} else if vote == "down" {
				msg.Rating--
			}
			w.Write([]byte(fmt.Sprintf(`{"id":%d, "rating":%d}`, msg.ID, msg.Rating)))
		} else {
			w.Write([]byte(`{"id":0, "rating":0}`))
		}
	})

	http.HandleFunc("/clear_comments", func(w http.ResponseWriter, r *http.Request) {
		messages = map[int]*Msg{}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

// <img src="/rate?id=1&vote=up">
