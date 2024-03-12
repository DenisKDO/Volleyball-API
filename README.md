# Volleyball-API
## Project decripton
Volleyball-API is API based on our favorite sport Volleyball. You will have access to all professional volleyball players,
their national teams and volleyball tournaments. Also in this API you can register and create you own an amateur or even a professional team. 

Key Features:

    1.Team management:
        - View lists of modern teams by id (in future we add more teams)
        - Viewing team lineups
        - Creating team (if desired)
        - Updating the team data
        - Deleting teams
     2.Player Management:
         - View the national teams of players in which they play
         - View player characteristics by id
         - Updating player data
         - Creating your “Ideal” players and adding them to the team
         - Deleting players

#### Use of technology:

  - Programming in Golang,
  - Gorilla Mux,
  - Library for working with the PostgreSQL database: GORM;
  - Used the driver "github.com/jinzhu/gorm"
  
  
### Team members:

```
    Kim Denis, 22B21B118
    Golovinov, Andrey 22B030124
```
## API endpoints:

```
POST /VolleyballAPI/team/create : Crete a new team
GET  /VolleyballAPI/team/{id} : Taking information about team by using id
GET  /VolleyballAPI/teams : Taking information about all existing teams in database
PUT  /VolleyballAPI/team/update/{id} : Update information about team by using id
DELETE /VolleyballAPI/team/delete/{id} : Delete team from database by using id

POST /VolleyballAPI/player/create : Crete a new player
GET  /VolleyballAPI/player/{id} : Taking information about player by using id
GET  /VolleyballAPI/players : Taking information about all existing players in database
PUT  /VolleyballAPI/player/update/{id} : Update information about player by using id
DELETE /VolleyballAPI/player/delete/{id} : Delete player from database by using id

```
## DB sructure:
```
  type Player struct { 
	gorm.Model 
  	FirstName   string
	SecondName  string 
	TeamID      int  
	DateOfBirth string 
	Age         int 
	SpikeHeight int 
  	BlockHeight int 
	Height      int  
}

type Team struct {
	gorm.Model
	Title   string
	Coach   string
	Players []Player
} 
```
