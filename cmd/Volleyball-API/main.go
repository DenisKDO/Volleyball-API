package main

import (
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/methods"
	"github.com/gorilla/mux"
	"net/http"
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

	//start api
	http.ListenAndServe(":8080", router)

	//close connection to database
	defer database.Db.Close()

}
