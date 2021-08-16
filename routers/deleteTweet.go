package routers

import (
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	IdTweet := r.URL.Query().Get("id")

	if len(IdTweet) < 1 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(IdTweet, UserId)

	if err != nil {
		http.Error(w, "Ocurrio un error al Eliminar el Tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
