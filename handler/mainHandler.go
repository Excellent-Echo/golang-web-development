package handler

import (
	"encoding/json"
	"fmt"
	"golangWebDev/config"
	"golangWebDev/entity"
	"log"
	"net/http"
	"path"
	"text/template"
)

func GetDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	dataSQL, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println("error query SQL")
		return
	}

	var datas []entity.Student

	for dataSQL.Next() {
		var each = entity.Student{}

		if err := dataSQL.Scan(&each.ID, &each.Fullname, &each.Age, &each.Batch); err != nil {
			log.Fatal(err.Error())
			return
		}

		datas = append(datas, each)
	}

	defer db.Close()

	r.Header.Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// excecute yg bener
		byteDatas, err := json.Marshal(datas)

		if err != nil {
			http.Error(w, "error internal server", http.StatusInternalServerError)
			return
		}

		w.Write(byteDatas)
		return
	}

	// kalau bukan method get tampilkan error
	http.Error(w, "error method not GET", http.StatusBadRequest)
}

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

// contoh method GET
func GetPost(w http.ResponseWriter, r *http.Request) {
	// r.Method untuk mengecek method yang dikirim oleh user
	// kita bisa menggunakan method lain, PUT, PATCH, DELETE
	if r.Method == "POST" {
		w.Write([]byte("ini adalah method POST"))
		return
	}

	http.NotFound(w, r)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "form.html"))

	if err != nil {
		http.Error(w, "error rendering html", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func PostFormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "error rendering html", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")
		checkbox := r.Form.Get("checkbox")

		data := map[string]interface{}{
			"name":     name,
			"message":  pesan,
			"checkbox": checkbox,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

		if err != nil {
			http.Error(w, "error rendering html", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)

		// w.Write([]byte())
		return
	}

	http.NotFound(w, r)
}
