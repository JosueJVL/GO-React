package jwt

import (
	"log"
	"time"

	"github.com/JosueJVL/GO-React/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(model models.User) (string, error) {

	mykey := []byte("MastersdelDesarrollo_grupodeFacebook")

	payload := jwt.MapClaims{
		"email":     model.Email,
		"nombre":    model.Name,
		"lastName":  model.LastName,
		"birthDay":  model.BirthDay,
		"biography": model.Biography,
		"webSite":   model.WebSite,
		"_id":       model.Id.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(mykey)

	if err != nil {
		log.Println(err)
		return tokenStr, err
	}

	return tokenStr, nil
}
