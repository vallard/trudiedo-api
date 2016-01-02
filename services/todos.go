package services

import (
	"fmt"

	"github.com/vallard/trudiedo-api/models"
)

type todosService struct{}

var Todos todosService

func (todosService) List() models.Todos {
	ts := models.Todos{
		models.Todo{Id: 1, Text: "Surf"},
		models.Todo{Id: 2, Text: "Skate"},
	}

	fmt.Println("%v", ts)
	return ts
}
