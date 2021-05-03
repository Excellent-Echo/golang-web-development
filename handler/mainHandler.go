package handler

import (
	"golangWebDev/entity"
	"net/http"
	"path"
	"text/template"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// kembalikan error
		http.Error(w, "404 error page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "root.html"), path.Join("views", "layout.html"))
	// w.Write([]byte("ini adalah root route"))
	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "hello.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

	// w.Write([]byte("hello world"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("ini adalah about"))
	tmpl, err := template.ParseFiles(path.Join("views", "about.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func ListProductHandler(w http.ResponseWriter, r *http.Request) {
	list := []entity.Product{
		{1, "Pajero", 300_000_000, 10},
		{2, "Xenioo", 200_000_000, 20},
		{3, "Daihatz", 250_000_000, 25},
	}

	tmpl, err := template.ParseFiles(path.Join("views", "list.html"), path.Join("views", "layout.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, list)
}
