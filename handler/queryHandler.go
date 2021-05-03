package handler

import (
	"net/http"
	"strconv"
)

// url parse : www.google.com"/about" : localhost:3000"/"
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// int  1 - sekian / sadad (salah) kdtx-01213 (salah)
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")

	// jika masukkan id 0 atau kurang dari 0 salah
	// masukkan sejenis string / code lain itu salah

	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber <= 0 {
		http.Error(w, "input id number error", http.StatusBadRequest)
		return
	}

	w.Write([]byte("ini adalah product : " + id + " bernama : " + name))
}
