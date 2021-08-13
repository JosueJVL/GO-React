package middlew

import (
	"net/http"

	"github.com/JosueJVL/GO-React/routers"
)

/* ValidateJWT permite validar el JWT de entrada por las peticiones */
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token !"+err.Error(), 404)
		}

		next.ServeHTTP(w, r)
	}
}
