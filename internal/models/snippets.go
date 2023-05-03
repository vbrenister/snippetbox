package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Snippet struct {
	ID      int       `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Created time.Time `db:"created"`
	Expires time.Time `db:"expires"`
}

type SnippetModel struct {
	DB *sqlx.DB
}

func NewSnipperModel(db *sqlx.DB) *SnippetModel {
	return &SnippetModel{
		DB: db,
	}
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := fmt.Sprintf(`INSERT INTO snippets (title, content, created, expires) 
	VALUES ($1, $2, current_timestamp, current_timestamp + INTERVAL '%d DAY') RETURNING id`, expires)

	var id int

	err := m.DB.Get(&id, stmt, title, content)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > current_timestamp AND id = $1`

	s := &Snippet{}

	err := m.DB.Get(s, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}

		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > current_timestamp ORDER BY created DESC LIMIT 10`

	snippets := []*Snippet{}

	err := m.DB.Select(&snippets, stmt)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}
