package main

import (
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//connection to database
	database.Database()

	//api routing
	router := mux.NewRouter()

	//healthcheck
	router.HandleFunc("/VolleyballAPI/healthcheck", handlers.HealthCheck).Methods("GET")

	//teams methods
	router.HandleFunc("/VolleyballAPI/teams", handlers.GetTeams).Methods("GET")
	router.HandleFunc("/VolleyballAPI/team/{id}", handlers.GetTeam).Methods("GET")
	router.HandleFunc("/VolleyballAPI/team/create", handlers.CreateTeam).Methods("POST")
	router.HandleFunc("/VolleyballAPI/team/delete/{id}", handlers.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/VolleyballAPI/team/update/{id}", handlers.UpdateTeam).Methods("PUT")

	//players methods
	router.HandleFunc("/VolleyballAPI/players", handlers.GetPlayers).Methods("GET")
	router.HandleFunc("/VolleyballAPI/player/{id}", handlers.GetPlayer).Methods("GET")
	router.HandleFunc("/VolleyballAPI/player/create", handlers.CreatePlayer).Methods("POST")
	router.HandleFunc("/VolleyballAPI/player/delete/{id}", handlers.DeletePlayer).Methods("DELETE")
	router.HandleFunc("/VolleyballAPI/player/update/{id}", handlers.UpdatePlayer).Methods("PUT")

	//coach methods
	router.HandleFunc("/VolleyballAPI/coach/{id}", handlers.GetCoach).Methods("GET")
	router.HandleFunc("/VolleyballAPI/coaches", handlers.GetCoaches).Methods("GET")
	router.HandleFunc("/VolleyballAPI/coach/create", handlers.CreateCoach).Methods("POST")
	router.HandleFunc("/VolleyballAPI/coach/delete/{id}", handlers.DeleteCoach).Methods("DELETE")
	router.HandleFunc("/VolleyballAPI/coach/update/{id}", handlers.UpdateCoach).Methods("PUT")

	//users methods
	router.HandleFunc("/VolleyballAPI/user/activate", handlers.ActivateUser).Methods("PUT")
	router.HandleFunc("/VolleyballAPI/user/reg", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/VolleyballAPI/user/authentication", handlers.Authentication).Methods("POST")

	//start api
	http.ListenAndServe(":8080", router)

	//close connection to database
	defer database.Db.Close()

}
