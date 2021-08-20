package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JosueJVL/GO-React/middlew"
	"github.com/JosueJVL/GO-React/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	route := mux.NewRouter()

	//endpoint de User
	route.HandleFunc("/registro", middlew.CheckDB(routers.Register)).Methods("POST")
	route.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	route.HandleFunc("/perfil", middlew.CheckDB(middlew.ValidateJWT(routers.GetPerfil))).Methods("GET")
	route.HandleFunc("/updateUsers", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateUser))).Methods("PUT")

	//endpoints de Tweets
	route.HandleFunc("/addTweets", middlew.CheckDB(middlew.ValidateJWT(routers.AddTweet))).Methods("POST")
	route.HandleFunc("/getTweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	route.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	//endpoint de Banner y Avatar
	route.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.AddAvatar))).Methods("POST")
	route.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(routers.AddBanner))).Methods("POST")
	route.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	route.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	//endPoints de Relation
	route.HandleFunc("/addrelations", middlew.CheckDB(middlew.ValidateJWT(routers.AddRelation))).Methods("POST")
	route.HandleFunc("/deleterelations", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	route.HandleFunc("/getrelations", middlew.CheckDB(middlew.ValidateJWT(routers.GetRelation))).Methods("GET")

	// Abrir el Puerto
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// los cors hace que sean accesible desde cualquier lugar
	// el AllowAll -- le permisos por completo
	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
