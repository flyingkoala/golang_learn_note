/*
@Time : 2020/6/4 13:56
@Author : wkang
@File : http
@Description:
*/
package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)

}
