package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/filters"
	"github.com/DenisKDO/Vollyball-API/internal/helper"
	"github.com/DenisKDO/Vollyball-API/internal/pagination"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	// Setting header JSON
	w.Header().Set("Content-Type", "application/json")

	var teams []models.Team
	var players []models.Player
	var coach []models.Coach
	var info models.Info

	// Getting query parameters
	query := r.URL.Query()
	fivbStr := query.Get("fivb")
	country := query.Get("country")
	sortField := query.Get("sort_by")
	pageStr := query.Get("page")
	pageSizeStr := query.Get("page_size")

	//Invalid query parameter
	v := validation.New()
	for i, _ := range query {
		if !validation.In(i, "fivb", "country", "sort_by", "page", "page_size") {
			w.WriteHeader(http.StatusBadRequest)
			v.AddError("ValidQueryParameter", "Invalid query parameter")
		}
	}

	//parse page and page size parameters
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	//error if page value invalid
	if pageStr != "" {
		_, erro := strconv.Atoi(pageStr)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			v.AddError("PageValue", "Invalid page value")
		}
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 1 // Default page size
	}

	//error if page size value invalid
	if pageSizeStr != "" {
		_, erro := strconv.Atoi(pageSizeStr)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			v.AddError("PageSize", "Invalid page size value")
		}
	}

	//offset
	offset := (page - 1) * pageSize
	// Taking instance of database
	db := database.Db

	if len(query) > 0 {
		// Filtration by height
		if fivbStr != "" {
			var errI, errR bool
			errI, errR, db = filters.FiltersByInt(fivbStr, w, db, "fivb_ranking", "team")
			if !errR {
				w.WriteHeader(http.StatusNotFound)
				v.AddError("FIVBRecords", "Records not found")
			}
			if !errI {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("FIVB", "Invalid value")
			}
		}

		// Filtration by position
		if country != "" {
			db = db.Where("country = ?", country)
			if helper.NoRecordsFindTeam(db, w, country) == 0 {
				v.AddError("CountryRecords", "Not found")
			}
		}

		//sorting
		if sortField != "" {
			if strings.Contains(sortField, "-") {
				sortField = sortField[1:]
				db = db.Order(sortField + " DESC")
			} else {
				db = db.Order(sortField + " ASC")

			}
		}
		db.Model(models.Team{}).Count(&info.Count)
		// Pagination
		db = db.Offset(offset).Limit(pageSize)

		if !v.Valid() {
			for key, message := range v.Errors {
				fmt.Fprintf(w, "-%s: %s\n", key, message)
			}
			return
		}

		// Request to DB
		db.Find(&teams)

		// Checking if in db any players
		if len(teams) == 0 {
			http.Error(w, "-Page: Not found", http.StatusNotFound)
			return
		}
	} else {
		db = db.Offset(offset).Limit(pageSize).Find(&teams)
	}

	var count int
	db.Model(&models.Team{}).Count(&count)
	info = pagination.InfoStructForPagination(r, pageStr, page, pageSize, info.Count, count)

	for index := range teams {
		database.Db.Model(&teams[index]).Related(&coach)
		database.Db.Model(&teams[index]).Related(&players)
		teams[index].Players = players
		teams[index].Coach = coach
	}

	response := map[string]interface{}{
		"info":   info,
		"result": teams,
	}

	// Return JSON of players
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
