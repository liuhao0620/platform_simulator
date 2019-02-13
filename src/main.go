package main

import (
	"data"
	"htmls"
	"net/http"
)

func main() {
	data.RedisInit()
	defer data.RedisClose()
	http.HandleFunc("/", htmls.IndexHandler)
	http.HandleFunc("/registe", htmls.RegisteHandler)
	http.HandleFunc("/changepassword", htmls.ChangePasswordHandler)
	http.HandleFunc("/checkuser", htmls.CheckUserHandler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
