package monitoring

import (
	"time"

	"github.com/go-co-op/gocron"
)

func Schedular(duration string) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(duration).Do(TestUrls)
	s.StartAsync()
}
