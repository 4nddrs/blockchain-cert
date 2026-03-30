package repository

import (
	"context"
	"github.com/4nddrs/blockchain-cert/database"
	"github.com/4nddrs/blockchain-cert/models"
)

// Check if the user used credit
func CheckTrialEligibility(ctx context.Context, ip string, fingerprint string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM trial_usage WHERE ip_address = $1 OR fingerprint = $2)`

	err := database.DB.QueryRow(ctx, query, ip, fingerprint).Scan(&exists)
	return !exists, err // Return true if the user is eligible (not found in trial_usage), false otherwise
}

func SaveCertificateAndMarkUsage(ctx context.Context, cert models.Certificate, ip string, fingerprint string, instID string, isTrial bool) error {
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	certQuery := `
		INSERT INTO certificates (institution_id, file_hash, student_name, course_name, tx_hash, metadata)
		VALUES ($1, $2, $3, $4, $5, $6)`

	var dbInstID interface{}
	if !isTrial {
		dbInstID = instID
	} else {
		dbInstID = nil // En Go/pgx, nil se traduce a NULL en Postgres
	}

	_, err = tx.Exec(ctx, certQuery, dbInstID, cert.FileHash, cert.StudentName, cert.CourseName, cert.TxHash, cert.Metadata)
	if err != nil {
		return err
	}

	if isTrial {
		trialQuery := `INSERT INTO trial_usage (ip_address, fingerprint, file_hash_registered) VALUES ($1, $2, $3)`
		_, err = tx.Exec(ctx, trialQuery, ip, fingerprint, cert.FileHash)
	} else {
		creditQuery := `UPDATE institutions SET credits_remaining = credits_remaining - 1 WHERE id = $1`
		_, err = tx.Exec(ctx, creditQuery, instID)
	}

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func GetInstitutionByAPIKey(ctx context.Context, apiKey string) (*models.Institution, error) {
	var inst models.Institution
	query := `SELECT id, name, plan_type, credits_remaining FROM institutions WHERE api_key = $1`

	err := database.DB.QueryRow(ctx, query, apiKey).Scan(
		&inst.ID, &inst.Name, &inst.PlanType, &inst.CreditsRemaining,
	)
	if err != nil {
		return nil, err
	}
	return &inst, nil
}

func GetInstitutionByID(ctx context.Context, id string) (*models.Institution, error) {
	var inst models.Institution
	query := `SELECT id, name, email, plan_type, credits_remaining, created_at FROM institutions WHERE id = $1`

	err := database.DB.QueryRow(ctx, query, id).Scan(
		&inst.ID, &inst.Name, &inst.Email, &inst.PlanType, &inst.CreditsRemaining, &inst.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &inst, nil
}

func GetCertificateByHash(ctx context.Context, fileHash string) (*models.Certificate, error) {
	var cert models.Certificate
	query := `SELECT student_name, course_name, tx_hash, created_at FROM certificates WHERE file_hash = $1`

	err := database.DB.QueryRow(ctx, query, fileHash).Scan(
		&cert.StudentName, &cert.CourseName, &cert.TxHash, &cert.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func CreateInstitution(ctx context.Context, name string, email string, plan string) (*models.Institution, error) {
	var inst models.Institution
	query := `
		INSERT INTO institutions (name, email, plan_type, credits_remaining)
		VALUES ($1, $2, $3, $4)
		RETURNING id, api_key, created_at`

	// Define credits based on the plan type
	initialCredits := 0
	if plan == "starter" {
		initialCredits = 1
	} else if plan == "institution" {
		initialCredits = 500
	}

	err := database.DB.QueryRow(ctx, query, name, email, plan, initialCredits).Scan(
		&inst.ID, &inst.APIKey, &inst.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	inst.Name = name
	inst.Email = email
	inst.PlanType = plan
	inst.CreditsRemaining = initialCredits

	return &inst, nil
}

func GetAllInstitutions(ctx context.Context) ([]models.Institution, error) {
	query := `SELECT id, name, email, plan_type, credits_remaining, created_at FROM institutions ORDER BY created_at DESC`

	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var institutions []models.Institution
	for rows.Next() {
		var i models.Institution
		if err := rows.Scan(&i.ID, &i.Name, &i.Email, &i.PlanType, &i.CreditsRemaining, &i.CreatedAt); err != nil {
			return nil, err
		}
		institutions = append(institutions, i)
	}
	return institutions, nil
}

func UpdateInstitutionCredits(ctx context.Context, id string, additionalCredits int) error {
	query := `UPDATE institutions SET credits_remaining = credits_remaining + $1 WHERE id = $2`
	_, err := database.DB.Exec(ctx, query, additionalCredits, id)
	return err
}

func UpdateInstitutionPlan(ctx context.Context, id string, newPlan string) error {
	query := `UPDATE institutions SET plan_type = $1 WHERE id = $2`
	_, err := database.DB.Exec(ctx, query, newPlan, id)
	return err
}

func DeleteInstitution(ctx context.Context, id string) error {
	query := `DELETE FROM institutions WHERE id = $1`
	_, err := database.DB.Exec(ctx, query, id)
	return err
}
