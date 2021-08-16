package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func AddAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")

	var extention = strings.Split(handler.Filename, ".")[1]

	var fileImage string = "uploads/avatars/" + UserId + "." + extention

	function, err := os.OpenFile(fileImage, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen"+err.Error(), 404)
		return
	}

	_, err = io.Copy(function, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), 404)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserId + "." + extention
	status, err = bd.UpdateUser(user, UserId)

	if err != nil || !status {
		http.Error(w, "Error al grabar el Avatar en la Base de Datos"+err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
