package database

import (
	"fmt"
	"github.com/DenisKDO/Vollyball-API/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var err error
var db *gorm.DB

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
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	//Closing connection to database
	defer db.Close()

	//migrations to the database
	db.AutoMigrate(&pkg.Player{})
	db.AutoMigrate(&pkg.Team{})
}
