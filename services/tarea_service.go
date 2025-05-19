package services

import (
	"github.com/MoonHack2077/parcial-so-delete/models"
	"github.com/MoonHack2077/parcial-so-delete/repositories"
)

func ObtenerTareas() ([]models.Tarea, error) {
	return repositories.ObtenerTodasTareas()
}

func EliminarTarea(id string) error {
	return repositories.EliminarTarea(id)
}
