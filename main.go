package main

import (
	"fmt"
	"golangWebDev/handler"
	"net/http"
)

func main() {
	// web development (server - client)
	// net/http

	// routing group

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	// tidak bisa menghandle params /product/:id_product
	http.HandleFunc("/product", handler.ProductHandler)
	http.HandleFunc("/list_product", handler.ListProductHandler)
	http.HandleFunc("/get_post", handler.GetPost)
	http.HandleFunc("/form_data", handler.FormHandler)
	http.HandleFunc("/process", handler.PostFormHandler)

	// 1 get data di database : query , Query / QueryRow
	// 2 kita simpen data db nya ke struct
	// 3 kita lakukan convert ke json file json.Marshal
	// - json.Unmarshal (mengubah dari json ke data)
	// - json.marshal (dari data ke data json)
	// kita tampilkan menggunakan fungsi handler
	http.HandleFunc("/students", handler.GetDatabaseHandler)

	port := "localhost:3000"
	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
