package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/jwt"
	"github.com/JosueJVL/GO-React/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("contect-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contrasena invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del Usuario es requerido", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del Usuario es requerido", 400)
	}

	user, exist := bd.Login(t.Email, t.Password)
	if !exist {
		http.Error(w, "Usuario y/o Contrasena invalidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar general el Token correspondiente"+err.Error(), 400)
		return
	}

	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
