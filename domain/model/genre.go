package model

import (
    "time"
)

type Genres []Genre

type Genre struct {
	ID          int64     `json:"id"`
	GenreKey    string    `json:"genre_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyGenre() *Genre {
	return &Genre{}
}

func (genre *Genre) IsEmpty() bool {
	return (
		genre.ID == 0 &&
		genre.GenreKey == "" &&
		genre.Name == "" &&
		genre.Description == "" &&
		genre.Type == "")
}
