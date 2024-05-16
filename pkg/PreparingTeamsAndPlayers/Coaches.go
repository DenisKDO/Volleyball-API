package PreparingTeamsAndPlayers

import "github.com/DenisKDO/Vollyball-API/pkg/essences"

func RussiaCoach() *essences.Coach {
	var coach essences.Coach

	coach.FirstName = "Konstantin"
	coach.SecondName = "Bryanskiyly"
	coach.Age = 46
	coach.DateOfBirth = "14 June, 1977"
	coach.Position = "L"
	coach.TeamID = 2
	return &coach
}

func JapanCoach() *essences.Coach {
	var coach essences.Coach

	coach.FirstName = "Philippe"
	coach.SecondName = "Blain"
	coach.Age = 63
	coach.DateOfBirth = "20 May, 1960"
	coach.Position = "OP"
	coach.TeamID = 1
	return &coach
}

func BrazilCoach() *essences.Coach {
	var coach essences.Coach

	coach.FirstName = "Bernardo"
	coach.SecondName = "Rezende"
	coach.Age = 64
	coach.DateOfBirth = "25 August, 1959"
	coach.Position = "S"
	coach.TeamID = 3
	return &coach
}

func GermanCoach() *essences.Coach {
	var coach essences.Coach

	coach.FirstName = "Michal"
	coach.SecondName = "Winiarski"
	coach.Age = 40
	coach.DateOfBirth = "28 September, 1983"
	coach.Position = "OH"
	coach.TeamID = 4
	return &coach
}

func PolandCoach() *essences.Coach {
	var coach essences.Coach

	coach.FirstName = "Nikola"
	coach.SecondName = "Grbic"
	coach.Age = 50
	coach.DateOfBirth = "6 September, 1973"
	coach.Position = "S"
	coach.TeamID = 5
	return &coach
}
