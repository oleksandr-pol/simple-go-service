package handlers

import (
	"html/template"
	"net/http"

	"github.com/oleksandr-pol/simple-go-service/internal/env"
)

type Material struct {
	Url   string
	Title string
}

func HandleMaterials(s *env.Server) (http.HandlerFunc, error) {

	// materials := []Material{
	// 	{"https://golang.org/pkg/net/http/", "HTTP package"},
	// 	{"https://golang.org/doc/articles/wiki/", "Writing Web Applications"},
	// 	{"https://gobyexample.com/http-servers", "Go by Example: HTTP Servers"},
	// 	{"https://yourbasic.org/golang/http-server-example/", "Hello world HTTP server example"},
	// 	{
	// 		"https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831",
	// 		"How I write Go HTTP services after seven years",
	// 	},
	// 	{"https://github.com/gorilla/mux", "Gorilla web toolkit"},
	// }

	tpl, tplErr := template.ParseFiles("./web/templates/materials.html")

	if tplErr != nil {
		return nil, tplErr
	}

	return func(w http.ResponseWriter, r *http.Request) {
		materials, err := s.Db.AllMaterials()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)

			return
		}
		tpl.Execute(w, materials)
	}, nil
}
