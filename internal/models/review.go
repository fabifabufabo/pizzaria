package models

type Review struct {
	Rating  int    `json:"rating"` // 1 to 5
	Comment string `json:"comment"`
}
