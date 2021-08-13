package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
)

func GetPerfil(w http.ResponseWriter, r *http.Request) {
	UserId := r.URL.Query().Get("id")

	if len(UserId) < 1 {
		http.Error(w, "Es necesario el parametro Id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.GetPerfil(UserId)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el Registro", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)

}
