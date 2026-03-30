package models

import "time"

type TrialUsage struct {
	ID          int       `db:"id" json:"id"`
	IPAddress   string    `db:"ip_address" json:"ip_address"`
	FileHash    string    `db:"file_hash_registered" json:"file_hash"`
	FingerPrint string    `db:"fingerprint" json:"fingerprint"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
