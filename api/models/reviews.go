package models

type Review struct{
	ID        string `json:"id,omitempty" db:"id"`
	Movie string `json:"movie,omitempty" db:"movie"`
	Reviewer  string `json:"reviewer,omitempty" db:"reviewer"`
	Rating     int `json:"rating,omitempty" db:"rating"`
	Comment   string `json:"comment,omitempty" db:"comment"`
}