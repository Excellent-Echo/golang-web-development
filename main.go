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

	port := "localhost:3000"
	fmt.Println("running on port", port)

	http.ListenAndServe(port, nil)
}
