# Volleyball-API
## Project decripton
Volleyball-API is API based on our favorite sport Volleyball. You will have access to all professional volleyball players,
their national teams and volleyball tournaments. Also in this API you can register and create you own an amateur or even a professional team. 

Key Features:

    1.Team management:
        - View lists of modern teams by index
        - Viewing team lineups
        - Creating new commands (if desired)
        - Updating the Command Database
        - Deleting teams
     2.Player Management:
         - View the national teams of players in which they play
         - View player characteristics by index
         - Updating player data
         - Creating your “Ideal” players and adding them to the team
         - Deleting players

#### Use of technology:

  - Programming in Golang,
  - Gorilla Mux,
  - Library for working with the PostgreSQL database: gorm/sql;
  - Used the driver gorm/mySQL
  
  
### Team members:

```
    Kim Denis, 22B21B118
    Golovinov, Andrey 22B030124
```
## API endpoints:

```
POST /team/create : Crete a new team
GET  /team/{id} : Taking information about team by using id
PUT  /team/update/{id} : Update information about team by using id
DELETE /team/delete/{id} : Delete team and memebers by using id

POST /player/create : Crete a new player
GET  /player/{id} : Taking information about player by using id
PUT  /player/update/{id} : Update information about player by using id
DELETE /player/delete/{id} : Delete player from team by using id

```
## DB sructure:
```
package pkg 
import "github.com/jinzhu/gorm"

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
