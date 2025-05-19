package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/MoonHack2077/Parcial2-SO/config"
	"github.com/MoonHack2077/Parcial2-SO/controllers"
)

func main() {
	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando .env")
	}

	// Conectar a MongoDB
	config.ConectarDB()

	// Crear router
	router := mux.NewRouter()

	// Registrar endpoints
	router.HandleFunc("/tareas", controllers.CrearTarea).Methods("POST")
	router.HandleFunc("/tareas", controllers.ObtenerTareas).Methods("GET")
	router.HandleFunc("/tareas/{id}", controllers.ActualizarTarea).Methods("PUT")
	router.HandleFunc("/tareas/{id}", controllers.EliminarTarea).Methods("DELETE")


	// Ruta raíz de prueba
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Microservicio de Tareas 🚀")
	}).Methods("GET")

	// Correr servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🟢 Servidor corriendo en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
