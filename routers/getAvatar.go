package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/JosueJVL/GO-React/bd"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	UserId := r.URL.Query().Get("id")

	if len(UserId) < 1 {
		http.Error(w, "ID requerido", http.StatusBadRequest)
		return
	}

	perfil, err := bd.GetPerfil(UserId)

	if err != nil {
		http.Error(w, "Usuario no encontrado", 404)
		return
	}

	file, _ := os.Open("uploads/avatars/" + perfil.Avatar)

	// if err != nil {
	// 	http.Error(w, "El Avatar no existe", 404)
	// 	return
	// }

	_, err = io.Copy(w, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", 404)
		return
	}
}
