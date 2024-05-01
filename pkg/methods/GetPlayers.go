package methods

import (
	"encoding/json"
	"fmt"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
	"strconv"
	"strings"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	// Setting header JSON
	w.Header().Set("Content-Type", "application/json")

	var players []essences.Player
	var info essences.Info

	// Getting query parameters
	query := r.URL.Query()
	heightStr := query.Get("height")
	weightStr := query.Get("weight")
	spikeHeightStr := query.Get("spike")
	blockHeightStr := query.Get("block")
	position := query.Get("position")
	sortField := query.Get("sort_by")
	pageStr := query.Get("page")
	pageSizeStr := query.Get("page_size")

	//parse page and page size parameters
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size
	}

	//offset
	offset := (page - 1) * pageSize
	// Taking instance of database
	db := database.Db

	if len(query) > 0 {
		// Filtration by height
		if heightStr != "" {
			var err bool
			err, db = filtersByInt(heightStr, w, db, "height")
			if err != true {
				return
			}
		}

		// Filtration by weight
		if weightStr != "" {
			var err bool
			err, db = filtersByInt(weightStr, w, db, "weight")
			if err != true {
				return
			}
		}

		//Filtration by spikeHeight
		if spikeHeightStr != "" {
			var err bool
			err, db = filtersByInt(spikeHeightStr, w, db, "spike_height")
			if err != true {
				return
			}
		}

		//Filtration by blockHeight
		if blockHeightStr != "" {
			var err bool
			err, db = filtersByInt(blockHeightStr, w, db, "block_height")
			if err != true {
				return
			}
		}

		// Filtration by position
		if position != "" {
			db = db.Where("position = ?", position)
			//checking if it has any players in db with certain position
			if NoRecordsFind(db, w, "position") == 0 {
				return
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
		db.Model(essences.Player{}).Count(&info.Count)
		// Pagination
		db = db.Offset(offset).Limit(pageSize)

		// Request to DB
		db.Find(&players)

		// Checking if in db any players
		if len(players) == 0 {
			http.Error(w, "Players not found", http.StatusNotFound)
			return
		}
	} else {
		db = db.Offset(offset).Limit(pageSize).Find(&players)
	}

	if len(query) == 0 {
		info.Count = 55
	}
	info.Pages = roundUp(float64(info.Count) / float64(pageSize))

	currentURL := r.URL.String()
	nextPage := page + 1
	prevPage := page - 1
	if nextPage > info.Pages {
		info.Next = "Last Page"
	}
	if pageStr == "" {
		info.Prev = "First Page"
		if len(query) == 0 {
			info.Next = currentURL + "?page=2"
		} else {
			info.Next = currentURL + "&page=2"
		}
	}
	if pageStr == "1" {
		info.Prev = "First Page"
	}
	if strings.Contains(currentURL, "page=") && info.Next == "" {
		info.Next = strings.Replace(currentURL, "page="+strconv.Itoa(page), "page="+strconv.Itoa(nextPage), -1)
	}
	if strings.Contains(currentURL, "page=") && info.Prev == "" {
		info.Prev = strings.Replace(currentURL, "page="+strconv.Itoa(page), "page="+strconv.Itoa(prevPage), -1)
	}

	// Return JSON of players
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&info)
	fmt.Fprintf(w, "result:\n")
	json.NewEncoder(w).Encode(&players)
}
