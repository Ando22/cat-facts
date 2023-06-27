package models

type Fact struct {
	ID     int    `json:"id" db:"id"`
	Fact   string `json:"fact" db:"fact"`
	Length int    `json:"length" db:"length"`
}
