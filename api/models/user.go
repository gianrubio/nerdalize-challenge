package models

import (
	"github.com/satori/go.uuid"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token uuid.UUID `json:"token"`
}
