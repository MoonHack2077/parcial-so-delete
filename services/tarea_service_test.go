package services

import (
	// "os"
	"context"
	"testing"

	// "github.com/joho/godotenv"
	"time"

	"github.com/MoonHack2077/parcial-so-delete/config"
	"github.com/MoonHack2077/parcial-so-delete/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestObtenerTareas(t *testing.T) {
	config.ConectarDB()
	tareas, err := ObtenerTareas()

	assert.NoError(t, err)
	assert.NotNil(t, tareas)
}

func TestEliminarTarea(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Crear una tarea para luego eliminar
	tarea := models.Tarea{
		Titulo:      "Tarea a eliminar",
		Descripcion: "Eliminar esta tarea en el test",
		Completado:  false,
	}
	tarea.CreadoEn = time.Now()
	collection := config.GetCollection("tareas")
	result, err := collection.InsertOne(ctx, tarea)
	assert.NoError(t, err)

	// Eliminar la tarea
	err = EliminarTarea(result.InsertedID.(primitive.ObjectID).Hex())
	assert.NoError(t, err)
}
