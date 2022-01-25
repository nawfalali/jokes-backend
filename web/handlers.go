package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"nawfalali.com/jokesapp/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Testing home %s\n", " All good")

}

func (app *application) randomJoke(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is a random joke")

}

func (app *application) showJoke(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		app.NotFound(w)
		return
	}

	s, err := app.jokes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}
	fmt.Fprintf(w, "%v", s)
}
