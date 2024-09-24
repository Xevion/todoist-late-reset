package main

import (
	"fmt"
	"internal/api"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	client *api.SyncClient
)

func init() {
	// Load .env file if available
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not loaded")
	}
}

func main() {
	// Load timezone
	// tzLocation, err := time.LoadLocation("America/Chicago")
	// if err != nil {
	// 	fmt.Println("Error loading location:", err)
	// 	os.Exit(1)
	// }

	client = api.NewSyncClient(os.Getenv("TODOIST_API_KEY"))
	client.UseResources(api.Items)

	for {
		_, err := client.Synchronize(false)

		fmt.Printf("Items: %d\n", len(client.State.Items))
		for _, item := range client.State.Items {
			fmt.Println(item)
		}

		if err != nil {
			fmt.Println("Error syncing:", err)
			os.Exit(1)
		}

		time.Sleep(30 * time.Second)
	}
}
