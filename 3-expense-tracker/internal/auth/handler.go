package auth

import "net/http"


func SignupHandler(w http.ResponseWriter, r *http.Request){
w.Write([] byte("signup handler"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	w.Write([] byte("login handler"))
	}