package handler

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Entity struct {
	Value string
}

func (u *user) addUserData() error {
	ctx := context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	dsClient, err := datastore.NewClient(ctx, "my-project")
	if err != nil {
		return errors.Wrap(err, "Failed to create datastore client")
	}

	k := datastore.NameKey("users", uuid.NewV1().String(), nil)

	if _, err := dsClient.Put(ctx, k, u); err != nil {
		return errors.Wrap(err, "Failed to upload data to datastore")
	}
	return nil

}
