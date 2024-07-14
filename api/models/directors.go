package models

type Director struct{
	ID        string `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
	BirthYear  int `json:"birth_year,omitempty" db:"birth_year"`
	Nationality     string `json:"nationality,omitempty" db:"nationality"`
	MoviesDirected   int `json:"movies_directed,omitempty" db:"movies_directed"`
}