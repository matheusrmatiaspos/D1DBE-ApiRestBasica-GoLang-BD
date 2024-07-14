package models

type Genre struct{
	ID        string `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
	Description  string `json:"description,omitempty" db:"description"`
	MovieCount     int `json:"movie_count,omitempty" db:"movie_count"`
	CreatedAt   string `json:"created_at,omitempty" db:"created_at"`
}