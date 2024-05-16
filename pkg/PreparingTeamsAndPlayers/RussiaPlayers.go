package PreparingTeamsAndPlayers

import "github.com/DenisKDO/Vollyball-API/pkg/models"

//making Russia Players

func RussiaPlayers() []models.Player {
	var players []models.Player

	var player models.Player

	player.UniformNumber = 1
	player.Position = "OH"
	player.FirstName = "Yaroslav"
	player.SecondName = "Podlesnykh"
	player.Height = 198
	player.Weight = 89
	player.SpikeHeight = 341
	player.BlockHeight = 330
	player.DateOfBirth = "3 September 1994"
	player.Age = 29
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 4
	player.Position = "MB"
	player.FirstName = "Artem"
	player.SecondName = "Volvich"
	player.Height = 212
	player.Weight = 96
	player.SpikeHeight = 350
	player.BlockHeight = 330
	player.DateOfBirth = "22 June 1990"
	player.Age = 34
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 7
	player.Position = "OH"
	player.FirstName = "Dmitry"
	player.SecondName = "Volkov"
	player.Height = 201
	player.Weight = 88
	player.SpikeHeight = 340
	player.BlockHeight = 330
	player.DateOfBirth = "25 May 1995"
	player.Age = 28
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 9
	player.Position = "MB"
	player.FirstName = "Ivan"
	player.SecondName = "Iakovlev"
	player.Height = 207
	player.Weight = 89
	player.SpikeHeight = 360
	player.BlockHeight = 350
	player.DateOfBirth = "17 April 1995"
	player.Age = 28
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 10
	player.Position = "OH"
	player.FirstName = "Denis"
	player.SecondName = "Bogdan"
	player.Height = 200
	player.Weight = 92
	player.SpikeHeight = 350
	player.BlockHeight = 340
	player.DateOfBirth = "13 October 1996"
	player.Age = 27
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 11
	player.Position = "S"
	player.FirstName = "Pavel"
	player.SecondName = "Pankov"
	player.Height = 198
	player.Weight = 90
	player.SpikeHeight = 345
	player.BlockHeight = 330
	player.DateOfBirth = "14 August 1995"
	player.Age = 28
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 15
	player.Position = "OP"
	player.FirstName = "Viktor"
	player.SecondName = "Poletaev"
	player.Height = 196
	player.Weight = 85
	player.SpikeHeight = 360
	player.BlockHeight = 340
	player.DateOfBirth = "27 July 1995"
	player.Age = 28
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 17
	player.Position = "OP"
	player.FirstName = "Maxim"
	player.SecondName = "Mikhaylov"
	player.Height = 202
	player.Weight = 103
	player.SpikeHeight = 360
	player.BlockHeight = 340
	player.DateOfBirth = "19 March 1988"
	player.Age = 35
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 18
	player.Position = "OH"
	player.FirstName = "Egor"
	player.SecondName = "Kliuka"
	player.Height = 209
	player.Weight = 93
	player.SpikeHeight = 370
	player.BlockHeight = 350
	player.DateOfBirth = "15 June 1995"
	player.Age = 28
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 20
	player.Position = "MB"
	player.FirstName = "Ilyas"
	player.SecondName = "Kurkaev"
	player.Height = 207
	player.Weight = 95
	player.SpikeHeight = 355
	player.BlockHeight = 335
	player.DateOfBirth = "18 January 1994"
	player.Age = 30
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 24
	player.Position = "S"
	player.FirstName = "Igor"
	player.SecondName = "Kobzar"
	player.Height = 198
	player.Weight = 86
	player.SpikeHeight = 337
	player.BlockHeight = 315
	player.DateOfBirth = "13 April 1991"
	player.Age = 32
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)

	player.UniformNumber = 27
	player.Position = "L"
	player.FirstName = "Valentin"
	player.SecondName = "Golubev"
	player.Height = 190
	player.Weight = 70
	player.SpikeHeight = 310
	player.BlockHeight = 305
	player.DateOfBirth = "3 May 1992"
	player.Age = 31
	player.TeamID = 2
	player.TeamTitle = "Russia men's national volleyball team"
	players = append(players, player)
	return players
}
