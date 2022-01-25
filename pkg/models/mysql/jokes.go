package mysql

import (
	"database/sql"
	"errors"

	"nawfalali.com/jokesapp/pkg/models"
)

type JokeModel struct {
	DB *sql.DB
}

func (m *JokeModel) Insert(title, category, body string, rating int) (int, error) {
	stmt := `insert into joke (title,category,body,rating) values (?, ?,?,?)`

	result, err := m.DB.Exec(stmt, title, category, body, rating)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *JokeModel) Get(id int) (*models.Joke, error) {

	stmt := `SELECT id,title,body,category,rating FROM jokes where id=?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Joke{}

	err := row.Scan(&s.ID, &s.Title, &s.Body, &s.Category, &s.Rating)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
