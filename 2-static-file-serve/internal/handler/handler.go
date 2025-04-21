package handler

import (
	"net/http"
)



func HTMLFileSender(w http.ResponseWriter, r *http.Request,filepath string)  {
	http.ServeFile(w,r,filepath)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))
}
