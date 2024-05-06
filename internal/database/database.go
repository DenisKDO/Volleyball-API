package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var err error
var Db *gorm.DB

func Database() {
	//loading enviroment variables to the system
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error")
	}
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DB_PORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	//Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	//Opening connection to database
	Db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	//migrations to the database
	Db.AutoMigrate(&essences.Player{})
	Db.AutoMigrate(&essences.Team{})
	Db.AutoMigrate(&essences.User{})
	Db.AutoMigrate(&essences.Token{})

	//Once created Teams in database from PreparingTeamsAndPlayers pkg
	// Db.Create(PreparingTeamsAndPlayers.JapanNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.RussiaNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.BrazilNationalVolleyballTeam())
	// for idx := range PreparingTeamsAndPlayers.JapanPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.JapanPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.RussiaPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.RussiaPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.BrazilPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.BrazilPlayers()[idx])
	// }
}
