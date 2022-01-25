package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Joke struct {
	ID       int
	Title    string
	Category string
	Body     string
	Rating   int
}
