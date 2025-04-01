package models

import "time"

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Genres          []string  `json:"genres"`
	PublicationDate time.Time `json:"publicationDate"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}