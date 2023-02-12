package main

import (
	"net/http"
)

// http库简单使用
func main() {
	http.HandleFunc("/ping", requestHandler)
	http.ListenAndServe(":9999", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	res := "success!"
	w.Write([]byte(res))
}
