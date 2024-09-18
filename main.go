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
	_ "time/tzdata"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

var (
	todoistApiToken string
	redisURL        string
	cronSchedule    string
	client          = &http.Client{}
)

func init() {
	// Load .env file if available
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not loaded")
	}
}

func primary() error {
	log, err := getRecentlyCompleted()
	if err != nil {
		fmt.Println("Error getting recently completed tasks:", err)
		return err
	}

	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return err
	}

	now := time.Now().In(loc)
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endTime := startTime.Add(3 * time.Hour)

	for _, event := range log.Events {
		if event.EventDate.After(startTime) && event.EventDate.Before(endTime) {
			fmt.Printf("%s - %s\n", event.EventDate.In(loc).Format("Monday, January 2, 3:04 PM MST"), event.ExtraData["content"])
		}
	}

	return nil
}

func main() {
	todoistApiToken = os.Getenv("TODOIST_API_KEY")
	redisURL = os.Getenv("REDIS_URL")
	cronSchedule = os.Getenv("CRON_SCHEDULE")

	// opt, _ := redis.ParseURL(redisURL)
	// client := redis.NewClient(opt)

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("Error creating scheduler:", err)
		os.Exit(1)
	}

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.CronJob(cronSchedule, false),
		gocron.NewTask(primary),
	)
	if err != nil {
		fmt.Println("Error adding job to scheduler:", err)
		os.Exit(1)
	}

	// Start the scheduler
	s.Start()

	nextRun, err := j.NextRun()
	if err != nil {
		fmt.Println("Error getting next run time:", err)
		os.Exit(1)
	}
	durationUntilNextRun := time.Until(nextRun).Seconds()
	fmt.Printf("startup: next run in %.2f seconds: %v\n", durationUntilNextRun, nextRun.Format(time.RFC3339))

	if durationUntilNextRun > 1 {
		// Run the job immediately
		err = j.RunNow()
		fmt.Println("startup: running job immediately")
		if err != nil {
			fmt.Println("Error running job immediately:", err)
			os.Exit(1)
		}
	}

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
