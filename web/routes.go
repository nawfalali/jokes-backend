package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/v1/joke/random", app.showJoke)
	mux.HandleFunc("/v1/joke/test", app.randomJoke)

	return (mux)
}
