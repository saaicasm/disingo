package main

import (
	"log"

	"github.com/saaicasm/disingo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":429")
	log.Fatal(srv.ListenAndServe())
}
