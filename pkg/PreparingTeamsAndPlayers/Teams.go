package PreparingTeamsAndPlayers

//making teams that will be insert in database
import "github.com/DenisKDO/Vollyball-API/pkg/essences"

func JapanNationalVolleyballTeam() *essences.Team {
	var JapanNationalTeam essences.Team

	JapanNationalTeam.Title = "Japan men's national volleyball team"
	JapanNationalTeam.Nickname = "Ryujin Nippon"
	JapanNationalTeam.Association = "Japan Volleyball Association"
	JapanNationalTeam.FIVBRanking = 4
	return &JapanNationalTeam
}

func RussiaNationalVolleyballTeam() *essences.Team {
	var RussiaNationalTeam essences.Team

	RussiaNationalTeam.Title = "Russia men's national volleyball team"
	RussiaNationalTeam.Nickname = "Reds Caesar Land"
	RussiaNationalTeam.Association = "Russian Volleyball Federation"
	return &RussiaNationalTeam
}

func BrazilNationalVolleyballTeam() *essences.Team {
	var BrazilNationalTeam essences.Team

	BrazilNationalTeam.Title = "Brazil men's national volleyball team"
	BrazilNationalTeam.Nickname = "Canarinhos Galactic Best of All Times"
	BrazilNationalTeam.Association = "CBV"
	BrazilNationalTeam.FIVBRanking = 5
	return &BrazilNationalTeam
}

func GermanNationalVolleyballTeam() *essences.Team {
	var BrazilNationalTeam essences.Team

	BrazilNationalTeam.Title = "German men's national volleyball team"
	BrazilNationalTeam.Nickname = "Die Adler (The Eagles)"
	BrazilNationalTeam.Association = "Deutscher Volleyball-Verband"
	BrazilNationalTeam.FIVBRanking = 10
	return &BrazilNationalTeam
}

func PolandNationalVolleyballTeam() *essences.Team {
	var BrazilNationalTeam essences.Team

	BrazilNationalTeam.Title = "Poland men's national volleyball team"
	BrazilNationalTeam.Nickname = "Bialo-Czerwoni"
	BrazilNationalTeam.Association = "Polish Volleyball Federation"
	BrazilNationalTeam.FIVBRanking = 1
	return &BrazilNationalTeam
}
