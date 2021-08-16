package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JosueJVL/GO-React/bd"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	UserId := r.URL.Query().Get("id")

	if len(UserId) < 1 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page requerido", http.StatusBadRequest)
		return
	}

	// Atoi, conversion
	pageValue, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Page debe ser mayot a 0 un error", 400)
		return
	}

	pageInt := int64(pageValue)

	response, bool := bd.GetTweets(UserId, pageInt)

	if !bool {
		http.Error(w, "Ocurrio un error al leer los tweets", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
