package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
)

func AddTweet(w http.ResponseWriter, r *http.Request) {
	var model models.RequestTweet

	infRequest := json.NewDecoder(r.Body).Decode(&model)

	if infRequest != nil {
		http.Error(w, "Datos incorrectos"+infRequest.Error(), 404)
		return
	}

	registro := models.Tweet{
		UserId:  UserId,
		Message: model.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.AddTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el Registro"+err.Error(), 404)
		return
	}

	if !status {
		http.Error(w, "No se pudo registrar el Tweet", 404)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
