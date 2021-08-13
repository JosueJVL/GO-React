package routers

import (
	"errors"
	"strings"

	"github.com/JosueJVL/GO-React/bd"
	"github.com/JosueJVL/GO-React/models"
	"github.com/dgrijalva/jwt-go"
)

/* */
var Email string
var UserId string

/* Proceso que valida y extrae el Token */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	mykey := []byte("MastersdelDesarrollo_GO")

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de Token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenInf, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return mykey, nil
	})

	if err == nil {
		_, info, _ := bd.ValidateUser(claims.Email)
		if info {
			UserId = claims.ID.Hex()
			Email = claims.Email
		}

		return claims, info, UserId, nil
	}

	if !tokenInf.Valid {
		return claims, false, string(""), errors.New("formato de Token invalido")
	}

	return claims, false, string(""), err
}
