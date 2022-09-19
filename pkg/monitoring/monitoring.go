package monitoring

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Ghazallghe/gofiber-monitoring/pkg/db"
	"github.com/Ghazallghe/gofiber-monitoring/pkg/models"
)

func TestUrls() {
	var urls []models.Url
	result := db.DB.Find(&urls)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}

	for _, url := range urls {
		go testUrl(url)
	}
}

func testUrl(url models.Url) {
	resp, err := http.Get(url.Url)
	if err != nil {
		log.Println(err)
		return
	}

	currentTime := time.Now()
	// MM-DD-YYYY
	date := currentTime.Format("01-02-2006")

	stat := models.Statistics{Date: date, UrlId: url.ID, Url: url.Url}
	status := models.Statistics{}

	findStatus(&status, resp.Status)

	existingStats := new(models.Statistics)

	result := db.DB.Where(stat).Attrs(status).FirstOrCreate(&existingStats)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		existingStats.StatusOk += status.StatusOk
		existingStats.ClientError += status.ClientError
		existingStats.ServerError += status.ServerError
		db.DB.Save(&existingStats)
	}
}

func findStatus(status *models.Statistics, respStatus string) {
	switch {
	case strings.HasPrefix(respStatus, "2"):
		status.StatusOk++
	case strings.HasPrefix(respStatus, "4"):
		status.ClientError++
	case strings.HasPrefix(respStatus, "5"):
	}
}
