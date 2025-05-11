package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"notification/ent/migrate"
	"notification/ent/retry"

	"notification/ent"

	_ "github.com/lib/pq"
)

func DBConnect() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=notification password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	err = client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	// Run the auto migration tool.
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func QueryRetry(ctx context.Context, client *ent.Client, uuid uuid.UUID) (*ent.Retry, error) {
	u, err := client.Retry.
		Query().
		Where(retry.MessageUUID(uuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying message: %w", err)
	}
	log.Println("message returned: ", u)
	return u, nil
}
