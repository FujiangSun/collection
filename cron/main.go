package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type GreetingJob struct {
	Name string
}

func (g *GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}
func main() {
	nyc, _ := time.LoadLocation("America/New_York")
	c := cron.New(cron.WithLocation(nyc))
	c.AddJob("@every 1s", &GreetingJob{"dj"})
	c.AddFunc("0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at New York")
	})

	c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at Tokyo")
	})
	c.Start()

	time.Sleep(9 * time.Second)
}
