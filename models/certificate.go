package models

import (
	"github.com/google/uuid"
	"time"
)

type Certificate struct {
	ID                uuid.UUID              `db:"id" json:"id"`
	InstitutionID     uuid.UUID              `db:"institution_id" json:"institution_id"`
	FileHash          string                 `db:"file_hash" json:"file_hash"`
	StudentName       string                 `db:"student_name" json:"student_name"`
	CourseName        string                 `db:"course_name" json:"course_name"`
	TxHash            string                 `db:"tx_hash" json:"tx_hash"`
	BlockchainNetwork string                 `db:"blockchain_network" json:"blockchain_network"`
	Metadata          map[string]interface{} `db:"metadata" json:"metadata"`
	CreatedAt         time.Time              `db:"created_at" json:"created_at"`
}
