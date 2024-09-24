package develop

import (
	"fmt"
	"internal/api"
	"os"

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
	client.sync(true)

	for {
		_, changes, err := client.sync(false)

		if err != nil {
			fmt.Println("Error syncing:", err)
			os.Exit(1)
		}
	}
}
