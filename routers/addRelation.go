package routers

import (
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func AddRelation(w http.ResponseWriter, r *http.Request) {
	UserRelationId := r.URL.Query().Get("id")
	if len(UserId) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var model models.Relation

	model.UserId = UserId
	model.UserRelationId = UserRelationId

	status, err := bd.AddRelation(model)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la Relacion con el Usuario", 404)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relacion", 404)
		return
	}

	w.WriteHeader(http.StatusOK)
}
