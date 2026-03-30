package models

import (
	"github.com/google/uuid"
	"time"
)

type BillingLog struct {
	ID            uuid.UUID `db:"id" json:"id"`
	InstitutionID uuid.UUID `db:"institution_id" json:"institution_id"`
	AmountUSD     float64   `db:"amount_usd" json:"amount_usd"`
	TxHash        string    `db:"tx_hash" json:"tx_hash"` // Referencia cruzada con la transacción en Polygon
	Status        string    `db:"status" json:"status"`   // 'pending', 'paid'
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}
