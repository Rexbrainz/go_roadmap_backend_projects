package main

import (
	"fmt"
	"net/http"
	"os"

	converterhttp "github.com/Rexbrainz/converters/http"
)

func main() {
	router := converterhttp.NewRouter()

	server := &http.Server{
		Addr:		":4000",
		Handler:	router,
	}

	fmt.Println("Listening on Port 4000")
	err := server.ListenAndServe()
	fmt.Fprintln(os.Stderr, err)
}
