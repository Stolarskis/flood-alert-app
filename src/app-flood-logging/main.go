package log

import (
	"log"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

//Creates the logger
func createLogger(name string) logging.Logger {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "flood-alert-app"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logger := client.Logger(name).StandardLogger(logging.Info)

	logger.Printf("Logger Created")
	return logger
}
