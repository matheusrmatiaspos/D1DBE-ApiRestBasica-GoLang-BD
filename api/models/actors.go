package models

type Actor struct{
	ID        string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	BirthYear  int `json:"birth_year" db:"birth_year"`
	Nationality     string `json:"nationality" db:"nationality"`
	MoviesStarred   int `json:"movies_starred" db:"movies_starred"`
}