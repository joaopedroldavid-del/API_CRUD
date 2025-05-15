package models

import "github.com/google/uuid"

type ID uuid.UUID

type User struct {
	ID uuid.UUID
	FirstName string
	LastName  string
	Biography string 
}

type Application struct {
	Data map[ID]User
}
