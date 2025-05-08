package database

import (
	"context"
	"log"

	"notification/ent"

	_ "github.com/lib/pq"
)

func DBConnect() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=notification password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
