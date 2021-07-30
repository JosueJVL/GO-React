package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JosueJVL/GO-React/middlew"
	"github.com/JosueJVL/GO-React/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	route := mux.NewRouter()

	route.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")

	// Abrir el Puerto
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// los cors hace que sean accesible desde cualquier lugar
	// el AllowAll -- le permisos por completo
	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
