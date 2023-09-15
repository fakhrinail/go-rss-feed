package main

import (
	"time"

	"github.com/fakhrinail/go-rss-feed/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey string `json:"api_key"`
	Name      string `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey: dbUser.ApiKey,
		Name: dbUser.Name,
	}
}
