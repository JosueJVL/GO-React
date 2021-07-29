package main

import (
	"log"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conecion a la Base de Datos")
		return
	}

	handlers.Manejadores()

}
