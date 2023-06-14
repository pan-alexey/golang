package router

import (
	"fmt"
	"net/http"
	serverContract "veva/contract/server"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler!")
}

func Init(m *http.ServeMux) {
	m.HandleFunc("/hello", HelloHandler)
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"response": " ok"}`))
	})

	var empty serverContract.Test

	fmt.Printf("serverContract.Empty: %v\n", empty)
}
