// Uses go cron to schedule a job to run every 3 minutes from 1AM to 3AM.
// Queries the todoist API for task completion
// Tasks that are completed between 12AM and 3AM are considered 'not too late' and will be rescheduled one day earlier (if they are recurring tasks).
// Only tasks with the included label (e.g. well-being) are considered.
// Most tasks are scheduled daily, so this means they'll be rescheduled for the current day. But not all.
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	todoistApiToken = os.Getenv("TODOIST_API_TOKEN")
	redisURL        = os.Getenv("REDIS_URL")
	cronSchedule    = os.Getenv("CRON_SCHEDULE")
	client          = &http.Client{}
)

func init() {
	// Load .env file if available
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not loaded")
	}
}

func main() {
	opt, _ := redis.ParseURL(redisURL)
	client := redis.NewClient(opt)

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// TODO: handle error
	}

	// add a job to the scheduler
	_, err = s.NewJob(
		gocron.CronJob(cronSchedule, false),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println("Current time:", time.Now().Format(time.RFC3339))
			},
			"hello",
			1,
		),
	)
	if err != nil {
		// handle error
	}

	// Start the scheduler
	s.Start()

	// Setup signal handler channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)    // Ctrl+C signal
	signal.Notify(stop, syscall.SIGTERM) // Container stop signal

	// Wait for signal (indefinite)
	closingSignal := <-stop

	fmt.Println("Gracefully shutting down, received signal:", closingSignal.String())
	err = s.Shutdown()
	if err != nil {
		// handle error
	}
}
