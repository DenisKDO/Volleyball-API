package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DenisKDO/Vollyball-API/pkg/models"
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

	if err := Db.Exec("SET NAMES 'UTF8'").Error; err != nil {
		log.Fatal(err)
	}
	//migrations to the database
	Db.AutoMigrate(&models.Player{})
	Db.AutoMigrate(&models.Team{})
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Token{})
	Db.AutoMigrate(&models.Coach{})

	//Once created Teams in database from PreparingTeamsAndPlayers pkg
	// Db.Create(PreparingTeamsAndPlayers.JapanNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.RussiaNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.BrazilNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.GermanNationalVolleyballTeam())
	// Db.Create(PreparingTeamsAndPlayers.PolandNationalVolleyballTeam())
	// for idx := range PreparingTeamsAndPlayers.JapanPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.JapanPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.RussiaPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.RussiaPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.BrazilPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.BrazilPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.GermanPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.GermanPlayers()[idx])
	// }
	// for idx := range PreparingTeamsAndPlayers.PolandPlayers() {
	// 	Db.Create(&PreparingTeamsAndPlayers.PolandPlayers()[idx])
	// }
	// Db.Create(PreparingTeamsAndPlayers.JapanCoach())
	// Db.Create(PreparingTeamsAndPlayers.RussiaCoach())
	// Db.Create(PreparingTeamsAndPlayers.BrazilCoach())
	// Db.Create(PreparingTeamsAndPlayers.GermanCoach())
	// Db.Create(PreparingTeamsAndPlayers.PolandCoach())
}
