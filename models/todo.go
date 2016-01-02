package models

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	Tags      Tags      `json:"tags"`
}

type Todos []Todo
