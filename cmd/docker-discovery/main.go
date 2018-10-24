package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Page struct {
	Title string
	Body  string
}

var markdownText string = `##ABC
> 123`

// The template
var templateText string = `
<head>
  <title>{{.Title}}</title>
</head>

<body>
  {{.Body | markDown}}
</body>
`

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}

func markDowner(args ...interface{}) template.HTML {
	return template.HTML(blackfriday.Run([]byte(fmt.Sprintf("%s", args...))))
	// return template.HTML(strings.ToLower(fmt.Sprintf("%s", args...)))
}

func main() {
	r := newRouter()
	p := &Page{Title: "A Test Demo", Body: markdownText}
	tmpl := template.Must(template.New("page.html").Funcs(template.FuncMap{"markDown": markDowner}).Parse(templateText))

	// Execute the template
	err := tmpl.ExecuteTemplate(os.Stdout, "page.html", p)
	if err != nil {
		fmt.Println(err)
	}
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
