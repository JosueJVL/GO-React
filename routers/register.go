package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var model models.User

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, "Error favor de validar los datos"+err.Error(), 400)
		return
	}

	if len(model.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(model.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, validateUser, _ := bd.ValidateUser(model.Email)
	if validateUser {
		http.Error(w, "Ya existe un usuario registrado con ese email", 404)
		return
	}

	_, status, err := bd.InsertUser(model)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
