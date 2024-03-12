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
	router.HandleFunc("/teams", methods.GetTeams).Methods("GET")
	router.HandleFunc("/team/{id}", methods.GetTeam).Methods("GET")
	router.HandleFunc("/team/create", methods.CreateTeam).Methods("POST")
	router.HandleFunc("/team/delete/{id}", methods.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/team/update/{id}", methods.UpdateTeam).Methods("PUT")

	//players methods
	//router.HandleFunc("/players", methods.GetPlayers).Methods("GET")

	//start api
	http.ListenAndServe(":8080", router)

	//close connection to database
	defer database.Db.Close()

}
