package models

import (
	"github.com/google/uuid"
	"time"
)

type Institution struct {
	ID               uuid.UUID `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	Email            string    `db:"email" json:"email"`
	APIKey           uuid.UUID `db:"api_key" json:"api_key"`
	PlanType         string    `db:"plan_type" json:"plan_type"`
	CreditsRemaining int       `db:"credits_remaining" json:"credits_remaining"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
