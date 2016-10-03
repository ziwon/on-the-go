package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("list").Parse(`
<html>
<head>
	<title>Hello</title>
</head>
<body>
	<h1>List</h1>
		<ul>
		{{ range $k, $v := . }}
			<li><strong>{{ $k }}</strong>: {{ $v }}</li>
		{{ end }}
		</ul>
</body>
</html>
`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := tmpl.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to execute template: %q\n", err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
