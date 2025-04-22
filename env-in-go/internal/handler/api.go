package handler

import "net/http"

func TestApi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}