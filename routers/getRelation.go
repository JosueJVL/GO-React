package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func GetRelation(w http.ResponseWriter, r *http.Request) {
	UserRelationId := r.URL.Query().Get("id")

	if len(UserId) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var model models.Relation
	model.UserId = UserId
	model.UserRelationId = UserRelationId

	var response models.ResponseGetRelation

	status, err := bd.GetRelation(model)
	if err != nil || !status {
		response.Status = status
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
