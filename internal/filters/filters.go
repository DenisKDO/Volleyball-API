package filters

import (
	"net/http"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/helper"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/jinzhu/gorm"
)

func FiltersByInt(queryStr string, w http.ResponseWriter, db *gorm.DB, parameter string, v *validation.Validator) (bool, *gorm.DB) {
	if strings.Contains(queryStr, ">") || strings.Contains(queryStr, "<") {
		//value that will contain > or <
		tempComparison := ""

		//defining if we have > or <
		if strings.Contains(queryStr, ">") {
			tempComparison = ">"
		} else if strings.Contains(queryStr, "<") {
			tempComparison = "<"
		}

		//deleting first char to convert into int
		queryStr = queryStr[1:]

		//converting str into int
		queryInt := validation.StrToInt(queryStr, w, parameter, v)
		if queryInt == 0 {
			return false, db
		}

		//switch cases to define if we want something > or < than queryInt
		switch tempComparison {
		case ">":
			db = db.Where(parameter+"> ?", queryInt)
		case "<":
			db = db.Where(parameter+"< ?", queryInt)
		}

		//error if no records find
		if helper.NoRecordsFind(db, w, parameter) == 0 {
			return false, db
		}
	} else {
		//if = to some queryParameter
		queryInt := validation.StrToInt(queryStr, w, parameter, v)
		if queryInt == 0 {
			return false, db
		}

		//finding
		db = db.Where(parameter+"= ?", queryInt)

		//error if no records find
		if helper.NoRecordsFind(db, w, parameter) == 0 {
			return false, db
		}
	}
	return true, db
}
