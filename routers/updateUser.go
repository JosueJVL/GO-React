package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var model models.User

	infRequest := json.NewDecoder(r.Body).Decode(&model)
	if infRequest != nil {
		http.Error(w, "Datos incorrectos"+infRequest.Error(), 404)
		return
	}

	status, error := bd.UpdateUser(model, UserId)
	if error != nil {
		http.Error(w, "Ocurrio un error al actualizar el Usuario", 404)
		return
	}

	if !status {
		http.Error(w, "No se pudo actualizar el Registro", 404)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
