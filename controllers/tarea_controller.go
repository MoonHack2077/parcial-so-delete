package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MoonHack2077/parcial-so-delete/services"
	"github.com/gorilla/mux"
)

func ObtenerTareas(w http.ResponseWriter, r *http.Request) {
	tareas, err := services.ObtenerTareas()
	if err != nil {
		http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tareas)
}

func EliminarTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := services.EliminarTarea(id)
	if err != nil {
		http.Error(w, "No se pudo eliminar la tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Tarea eliminada correctamente"})
}
