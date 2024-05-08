package main

import (
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/methods"
	"github.com/gorilla/mux"
)

func main() {
	//connection to database
	database.Database()

	//api routing
	router := mux.NewRouter()

	//teams methods
	router.HandleFunc("/VolleyballAPI/teams", methods.GetTeams).Methods("GET")
	router.HandleFunc("/VolleyballAPI/team/{id}", methods.GetTeam).Methods("GET")
	router.HandleFunc("/VolleyballAPI/team/create", methods.CreateTeam).Methods("POST")
	router.HandleFunc("/VolleyballAPI/team/delete/{id}", methods.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/VolleyballAPI/team/update/{id}", methods.UpdateTeam).Methods("PUT")

	//players methods
	router.HandleFunc("/VolleyballAPI/players", methods.GetPlayers).Methods("GET")
	router.HandleFunc("/VolleyballAPI/player/{id}", methods.GetPlayer).Methods("GET")
	router.HandleFunc("/VolleyballAPI/player/create", methods.CreatePlayer).Methods("POST")
	router.HandleFunc("/VolleyballAPI/player/delete/{id}", methods.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/VolleyballAPI/player/update/{id}", methods.UpdatePlayer).Methods("PUT")

	//users methods
	router.HandleFunc("/VolleyballAPI/user/activate", methods.ActivateUser).Methods("PUT")
	router.HandleFunc("/VolleyballAPI/user/reg", methods.RegisterUser).Methods("POST")
	router.HandleFunc("/VolleyballAPI/user/authentication", methods.Authentication).Methods("POST")

	//start api
	http.ListenAndServe(":8080", router)

	//close connection to database
	defer database.Db.Close()

}
