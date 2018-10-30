package handler

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type Entity struct {
	Value string
}

func addUserData() {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, "my-project")
	if err != nil {
		// Handle error.
	}

	k := datastore.NameKey("Entity", "stringID", nil)
	e := new(Entity)
	if err := dsClient.Get(ctx, k, e); err != nil {
		// Handle error.
	}

	old := e.Value
	e.Value = "Hello World!"

	if _, err := dsClient.Put(ctx, k, e); err != nil {
		// Handle error.
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Value)
}
