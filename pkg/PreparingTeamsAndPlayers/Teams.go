package PreparingTeamsAndPlayers

//making teams that will be insert in database
import "github.com/DenisKDO/Vollyball-API/pkg/models"

func JapanNationalVolleyballTeam() *models.Team {
	var JapanNationalTeam models.Team

	JapanNationalTeam.Title = "Japan men's national volleyball team"
	JapanNationalTeam.Nickname = "Ryujin Nippon"
	JapanNationalTeam.Association = "Japan Volleyball Association"
	JapanNationalTeam.Country = "Japan"
	JapanNationalTeam.FIVBRanking = 4
	return &JapanNationalTeam
}

func RussiaNationalVolleyballTeam() *models.Team {
	var RussiaNationalTeam models.Team

	RussiaNationalTeam.Title = "Russia men's national volleyball team"
	RussiaNationalTeam.Nickname = "Reds Caesar Land"
	RussiaNationalTeam.Association = "Russian Volleyball Federation"
	RussiaNationalTeam.Country = "Russia"
	return &RussiaNationalTeam
}

func BrazilNationalVolleyballTeam() *models.Team {
	var BrazilNationalTeam models.Team

	BrazilNationalTeam.Title = "Brazil men's national volleyball team"
	BrazilNationalTeam.Nickname = "Canarinhos Galactic Best of All Times"
	BrazilNationalTeam.Association = "CBV"
	BrazilNationalTeam.Country = "Brazil"
	BrazilNationalTeam.FIVBRanking = 5
	return &BrazilNationalTeam
}

func GermanNationalVolleyballTeam() *models.Team {
	var GermanNationalTeam models.Team

	GermanNationalTeam.Title = "German men's national volleyball team"
	GermanNationalTeam.Nickname = "Die Adler (The Eagles)"
	GermanNationalTeam.Association = "Deutscher Volleyball-Verband"
	GermanNationalTeam.Country = "German"
	GermanNationalTeam.FIVBRanking = 10
	return &GermanNationalTeam
}

func PolandNationalVolleyballTeam() *models.Team {
	var PolandNationalTeam models.Team

	PolandNationalTeam.Title = "Poland men's national volleyball team"
	PolandNationalTeam.Nickname = "Bialo-Czerwoni"
	PolandNationalTeam.Association = "Polish Volleyball Federation"
	PolandNationalTeam.Country = "Poland"
	PolandNationalTeam.FIVBRanking = 1
	return &PolandNationalTeam
}
