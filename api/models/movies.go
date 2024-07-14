package models

type Movie struct{
	ID        string `json:"id,omitempty" db:"id"`
	Title string `json:"title,omitempty" db:"title"`
	Director  string `json:"director,omitempty" db:"director"`
	Genre     string `json:"genre,omitempty" db:"genre"`
	ReleaseYear   int `json:"release_year,omitempty" db:"release_year"`
}