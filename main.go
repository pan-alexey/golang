package main

import (
	"fmt"
	"net/http"
	"veva/libs/server"
	"veva/router"
	"veva/utils/os"
)

func main() {
	mux := http.NewServeMux()
	router.Init(mux) // Register router

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	addr, err := server.Serve(&srv)
	if err != nil {
		fmt.Println("Error run server", err)
		return
	}

	fmt.Println("Server start at", "http://"+addr)
	os.WaitExit()
}
