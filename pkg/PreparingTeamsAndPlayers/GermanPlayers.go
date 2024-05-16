package PreparingTeamsAndPlayers

//making German Players
import "github.com/DenisKDO/Vollyball-API/pkg/models"

func GermanPlayers() []models.Player {
	var players []models.Player

	var player models.Player

	//German volleyball national team
	player.UniformNumber = 1
	player.Position = "OH"
	player.FirstName = "Cristian"
	player.SecondName = "Fromm"
	player.Height = 204
	player.Weight = 99
	player.SpikeHeight = 324
	player.BlockHeight = 324
	player.DateOfBirth = "August 15, 1990"
	player.Age = 33
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 2
	player.Position = "S"
	player.FirstName = "Johannes"
	player.SecondName = "Tille"
	player.Height = 186
	player.Weight = 73
	player.SpikeHeight = 331
	player.BlockHeight = 330
	player.DateOfBirth = "July 5, 1997"
	player.Age = 26
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 3
	player.Position = "OH"
	player.FirstName = "Denys"
	player.SecondName = "Kaliberda"
	player.Height = 193
	player.Weight = 95
	player.SpikeHeight = 343
	player.BlockHeight = 314
	player.DateOfBirth = "June 24, 1990"
	player.Age = 33
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 5
	player.Position = "OH"
	player.FirstName = "Moritrz"
	player.SecondName = "Reichert"
	player.Height = 194
	player.Weight = 85
	player.SpikeHeight = 336
	player.BlockHeight = 314
	player.DateOfBirth = "March 15, 1995"
	player.Age = 29
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 9
	player.Position = "OP"
	player.FirstName = "Gyorgy"
	player.SecondName = "Grozer"
	player.Height = 200
	player.Weight = 102
	player.SpikeHeight = 374
	player.BlockHeight = 345
	player.DateOfBirth = "November 27, 1984"
	player.Age = 39
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 10
	player.Position = "OP"
	player.FirstName = "Julian"
	player.SecondName = "Zenger"
	player.Height = 190
	player.Weight = 80
	player.SpikeHeight = 330
	player.BlockHeight = 315
	player.DateOfBirth = "August 26, 1997"
	player.Age = 26
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 12
	player.Position = "MB"
	player.FirstName = "Anton"
	player.SecondName = "Brehme"
	player.Height = 206
	player.Weight = 97
	player.SpikeHeight = 350
	player.BlockHeight = 330
	player.DateOfBirth = "10 August, 1999"
	player.Age = 24
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 13
	player.Position = "OH"
	player.FirstName = "Ruben"
	player.SecondName = "Schott"
	player.Height = 193
	player.Weight = 89
	player.SpikeHeight = 326
	player.BlockHeight = 309
	player.DateOfBirth = "July 8, 1994"
	player.Age = 29
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 14
	player.Position = "OH"
	player.FirstName = "Moritz"
	player.SecondName = "Karlitzek"
	player.Height = 191
	player.Weight = 91
	player.SpikeHeight = 335
	player.BlockHeight = 310
	player.DateOfBirth = "August 12, 1996"
	player.Age = 27
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 17
	player.Position = "S"
	player.FirstName = "Jan"
	player.SecondName = "Zimmermenn"
	player.Height = 191
	player.Weight = 82
	player.SpikeHeight = 340
	player.BlockHeight = 312
	player.DateOfBirth = "12 Februery, 1993"
	player.Age = 31
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 18
	player.Position = "MB"
	player.FirstName = "Florian"
	player.SecondName = "Krage"
	player.Height = 203
	player.Weight = 90
	player.SpikeHeight = 343
	player.BlockHeight = 331
	player.DateOfBirth = "January 13, 1993"
	player.Age = 31
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 19
	player.Position = "OH"
	player.FirstName = "Erik"
	player.SecondName = "Rohrs"
	player.Height = 201
	player.Weight = 82
	player.SpikeHeight = 338
	player.BlockHeight = 325
	player.DateOfBirth = "April 24, 2001"
	player.Age = 23
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 21
	player.Position = "MB"
	player.FirstName = "Tobias"
	player.SecondName = "Krick"
	player.Height = 210
	player.Weight = 85
	player.SpikeHeight = 350
	player.BlockHeight = 330
	player.DateOfBirth = "October 22, 1998"
	player.Age = 25
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	player.UniformNumber = 25
	player.Position = "L"
	player.FirstName = "Lukas"
	player.SecondName = "Maase"
	player.Height = 208
	player.Weight = 95
	player.SpikeHeight = 340
	player.BlockHeight = 320
	player.DateOfBirth = "August 28, 1998"
	player.Age = 23
	player.TeamTitle = "German men's national volleyball team"
	player.TeamID = 4
	players = append(players, player)

	return players

}
