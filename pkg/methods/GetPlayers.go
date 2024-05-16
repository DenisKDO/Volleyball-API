package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/filters"
	"github.com/DenisKDO/Vollyball-API/internal/pagination"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
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

	//Invalid query parameter
	v := validation.New()
	for i, _ := range query {
		if !validation.In(i, "height", "weight", "spike", "block", "position", "sort_by", "page", "page_size") {
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
		pageSize = 10 // Default page size
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
		if heightStr != "" {
			var errI, errR bool
			errI, errR, db = filters.FiltersByInt(heightStr, w, db, "height")
			if !errR {
				w.WriteHeader(http.StatusNotFound)
				v.AddError("HeightRecords", "Records not found")
			}
			if !errI {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("Height", "Invalid value")
			}
		}

		// Filtration by weight
		if weightStr != "" {
			var errI, errR bool
			errI, errR, db = filters.FiltersByInt(weightStr, w, db, "weight")
			if !errR {
				w.WriteHeader(http.StatusNotFound)
				v.AddError("WeightRecords", "Records not found")
			}
			if !errI {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("Weight", "Invalid value")
			}
		}

		//Filtration by spikeHeight
		if spikeHeightStr != "" {
			var errI, errR bool
			errI, errR, db = filters.FiltersByInt(spikeHeightStr, w, db, "spike_height")
			if !errR {
				w.WriteHeader(http.StatusNotFound)
				v.AddError("SpikeHeightRecords", "Records not found")
			}
			if !errI {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("SpikeHeight", "Invalid value")
			}
		}

		//Filtration by blockHeight
		if blockHeightStr != "" {
			var errI, errR bool
			errI, errR, db = filters.FiltersByInt(blockHeightStr, w, db, "block_height")
			if !errR {
				w.WriteHeader(http.StatusNotFound)
				v.AddError("BlockHeightRecords", "Records not found")
			}
			if !errI {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("BlockHeight", "Invalid value")
			}
		}

		// Filtration by position
		if position != "" {
			if !validation.In(position, "S", "L", "OH", "OP", "MB") {
				w.WriteHeader(http.StatusBadRequest)
				v.AddError("Position", "Invalid position value")
			} else {
				db = db.Where("position = ?", position)
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

		if !v.Valid() {
			for key, message := range v.Errors {
				fmt.Fprintf(w, "-%s: %s\n", key, message)
			}
			return
		}

		// Request to DB
		db.Find(&players)

		// Checking if in db any players
		if len(players) == 0 {
			http.Error(w, "-Page: Not found", http.StatusNotFound)
			return
		}
	} else {
		db = db.Offset(offset).Limit(pageSize).Find(&players)
	}

	info = pagination.InfoStructForPagination(r, pageStr, page, pageSize, info.Count)

	response := map[string]interface{}{
		"info":   info,
		"result": players,
	}

	// Return JSON of players
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
