package PreparingTeamsAndPlayers

//making teams that will be insert in database
import "github.com/DenisKDO/Vollyball-API/pkg/essences"

func JapanNationalVolleyballTeam() *essences.Team {
	var JapanNationalTeam essences.Team

	JapanNationalTeam.Title = "Japan men's national volleyball team"
	JapanNationalTeam.Nickname = "Ryujin Nippon"
	JapanNationalTeam.Association = "Japan Volleyball Association"
	JapanNationalTeam.Coach = "Philippe Blain"
	JapanNationalTeam.FIVBRanking = 4
	return &JapanNationalTeam
}

func RussiaNationalVolleyballTeam() *essences.Team {
	var RussiaNationalTeam essences.Team

	RussiaNationalTeam.Title = "Russia men's national volleyball team"
	RussiaNationalTeam.Nickname = "Reds Caesar Land"
	RussiaNationalTeam.Association = "Russian Volleyball Federation"
	RussiaNationalTeam.Coach = "Konstantin Bryanskiy"
	RussiaNationalTeam.FIVBRanking = -1
	return &RussiaNationalTeam
}
