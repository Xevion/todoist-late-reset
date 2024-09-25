// Uses go cron to schedule a job to run every 3 minutes from 1AM to 3AM.
// Queries the todoist API for task completion
// Tasks that are completed between 12AM and 3AM are considered 'not too late' and will be rescheduled one day earlier (if they are recurring tasks).
// Only tasks with the included label (e.g. well-being) are considered.
// Most tasks are scheduled daily, so this means they'll be rescheduled for the current day. But not all.
package main

import (
	"fmt"
	"internal/api"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

var (
	client     *api.SyncClient
	tzLocation *time.Location
)

func init() {
	// Load .env file if available
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not loaded")
	}
}

// primary is the main function that will be run by the scheduler
func primary() error {
	// Get recently completed tasks

	log, err := client.RecentlyCompleted()
	if err != nil {
		fmt.Println("Error fetching recently completed tasks:", err)
		return err
	}

	for _, event := range log.Events {
		// if event.EventDate.In(tzLocation).Hour() >= 1 && event.EventDate.In(tzLocation).Hour() < 3 {
		fmt.Printf("Task completed: %s at %s\n", event.ExtraData["content"], event.EventDate.In(tzLocation).Format("Monday, January 2, 3:04 PM MST"))
		// }
	}

	return nil
}

func main() {
	// redisURL := os.Getenv("REDIS_URL")
	cronSchedule := os.Getenv("CRON_SCHEDULE")
	// opt, _ := redis.ParseURL(redisURL)
	// client := redis.NewClient(opt)

	// Load timezone
	tzLocation, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading location:", err)
		os.Exit(1)
	}

	client = api.NewSyncClient(os.Getenv("TODOIST_API_KEY"))

	// create a scheduler
	s, err := gocron.NewScheduler(gocron.WithLocation(tzLocation))
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

	// Setup signal handler channel
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)    // Ctrl+C signal
	signal.Notify(stop, syscall.SIGTERM) // Container stop signal

	// Wait for signal (indefinite)
	closingSignal := <-stop

	fmt.Println("Gracefully shutting down, received signal:", closingSignal.String())
	err = s.Shutdown()
	if err != nil {
		fmt.Println("Error shutting down scheduler:", err)
		os.Exit(1)
	}
}
