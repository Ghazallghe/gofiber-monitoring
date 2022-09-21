package monitoring

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Schedular(duration string) {
	fmt.Println("I am working properly Dude")
	s := gocron.NewScheduler(time.UTC)
	s.Every(duration).Do(TestUrls)
	s.StartAsync()
}
