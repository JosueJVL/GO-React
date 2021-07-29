package middlew

import (
	"net/http"

	"github.com/JosueJVL/GO-React/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion  perdida con la Base de Datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}

}
