package routers

import (
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	UserRelationId := r.URL.Query().Get("id")

	if len(UserId) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}
	var model models.Relation

	model.UserId = UserId
	model.UserRelationId = UserRelationId

	status, err := bd.DeleteRelation(model)
	if err != nil {
		http.Error(w, "Ocurrio un error al Eliminar la Relacion", 404)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado Eliminar la Relacion", 404)
		return
	}

	w.WriteHeader(http.StatusOK)
}
