package pagination

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/helper"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
)

func InfoStructForPagination(r *http.Request, pageStr string, page int, pageSize int, count int) essences.Info {
	var info essences.Info
	info.Count = count
	query := r.URL.Query()

	//if no query parameters count in db equal to players in db
	if len(query) == 0 {
		info.Count = 55
	}

	//taking pages
	info.Pages = helper.RoundUp(float64(info.Count) / float64(pageSize))

	//current url string, next and prev pages
	currentURL := r.URL.String()
	nextPage := page + 1
	prevPage := page - 1

	//deticating last page
	if nextPage > info.Pages {
		info.Next = "Last Page"
	}

	//deticating first page
	if pageStr == "" {
		info.Prev = "First Page"
		if len(query) == 0 {
			info.Next = currentURL + "?page=2"
		} else {
			if nextPage <= info.Pages {
				info.Next = currentURL + "&page=2"
			}
		}
	}
	if pageStr == "1" {
		info.Prev = "First Page"
	}

	//string to prev and next pages
	if strings.Contains(currentURL, "page=") && info.Next == "" {
		info.Next = strings.Replace(currentURL, "page="+strconv.Itoa(page), "page="+strconv.Itoa(nextPage), -1)
	}
	if strings.Contains(currentURL, "page=") && info.Prev == "" {
		info.Prev = strings.Replace(currentURL, "page="+strconv.Itoa(page), "page="+strconv.Itoa(prevPage), -1)
	}
	return info
}
