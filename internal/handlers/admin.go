package handlers

import (
	"log"
	"net/http"

	"github.com/4nddrs/blockchain-cert/internal/repository"
	"github.com/4nddrs/blockchain-cert/models"
	"github.com/gin-gonic/gin"
)

// CreateInstitutionRequest represents the request body for creating an institution
type CreateInstitutionRequest struct {
	Name  string `json:"name" binding:"required" example:"Tech University"`
	Email string `json:"email" binding:"required,email" example:"contact@techuni.edu"`
	Plan  string `json:"plan" binding:"required" example:"premium" enums:"basic,premium,enterprise"`
}

// AddCreditsRequest represents the request body for adding/removing credits to/from an institution
// Positive values add credits, negative values remove credits
type AddCreditsRequest struct {
	AdditionalCredits int `json:"additional_credits" binding:"required" example:"500"`
}

// UpdatePlanRequest represents the request body for updating an institution's plan
type UpdatePlanRequest struct {
	NewPlan string `json:"new_plan" binding:"required" example:"enterprise" enums:"basic,premium,enterprise"`
}

// CreateInstitution godoc
// @Summary Create a new institution
// @Description Creates a new institution with an API key and initial credits based on the selected plan
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param institution body CreateInstitutionRequest true "Institution details"
// @Success 201 {object} models.Institution "Institution created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "Failed to create institution"
// @Router /admin/institutions [post]
func (h *Handler) CreateInstitution(c *gin.Context) {
	var input CreateInstitutionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inst, err := repository.CreateInstitution(c.Request.Context(), input.Name, input.Email, input.Plan)
	if err != nil {
		log.Printf("Error creating institution: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create institution"})
		return
	}

	c.JSON(http.StatusCreated, inst)
}

// ListInstitutions godoc
// @Summary List all institutions
// @Description Retrieves a list of all registered institutions with their credits and plan information
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Institution "List of institutions"
// @Failure 500 {object} ErrorResponse "Failed to retrieve institutions"
// @Router /admin/institutions [get]
func (h *Handler) ListInstitutions(c *gin.Context) {
	institutions, err := repository.GetAllInstitutions(c.Request.Context())
	if err != nil {
		log.Printf("Error listing institutions: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve institutions"})
		return
	}

	if institutions == nil {
		institutions = []models.Institution{}
	}

	c.JSON(http.StatusOK, institutions)
}

// AddCredits godoc
// @Summary Add or remove credits from an institution
// @Description Increases (positive) or decreases (negative) the credit balance for a specific institution. Example: 50 adds 50 credits, -20 removes 20 credits.
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Institution ID (UUID)"
// @Param credits body AddCreditsRequest true "Credits to add (positive) or remove (negative)"
// @Success 200 {object} map[string]interface{} "Credits updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body or insufficient credits"
// @Failure 500 {object} ErrorResponse "Failed to update credits"
// @Router /admin/institutions/{id}/credits [post]
func (h *Handler) AddCredits(c *gin.Context) {
	institutionID := c.Param("id")

	var input AddCreditsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation: ensure operation won't result in negative balance
	if input.AdditionalCredits < 0 {
		// Fetch current balance to check if removal is possible
		inst, err := repository.GetInstitutionByID(c.Request.Context(), institutionID)
		if err != nil {
			log.Printf("Error fetching institution: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify institution"})
			return
		}
		if inst.CreditsRemaining+input.AdditionalCredits < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":           "Insufficient credits",
				"current_balance": inst.CreditsRemaining,
				"requested":       input.AdditionalCredits,
			})
			return
		}
	}

	err := repository.UpdateInstitutionCredits(c.Request.Context(), institutionID, input.AdditionalCredits)
	if err != nil {
		log.Printf("Error updating credits: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update credits"})
		return
	}

	action := "added"
	if input.AdditionalCredits < 0 {
		action = "removed"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Credits " + action + " successfully",
		"institution_id": institutionID,
		"credits_delta":  input.AdditionalCredits,
	})
}

// UpdatePlan godoc
// @Summary Update institution plan
// @Description Changes the plan type for a specific institution (e.g., basic to premium)
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Institution ID (UUID)"
// @Param plan body UpdatePlanRequest true "New plan details"
// @Success 200 {object} map[string]interface{} "Plan updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "Failed to update plan"
// @Router /admin/institutions/{id}/plan [put]
func (h *Handler) UpdatePlan(c *gin.Context) {
	institutionID := c.Param("id")

	var input UpdatePlanRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.UpdateInstitutionPlan(c.Request.Context(), institutionID, input.NewPlan)
	if err != nil {
		log.Printf("Error updating plan for institution %s: %v", institutionID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Plan updated successfully",
		"institution_id": institutionID,
		"new_plan":       input.NewPlan,
	})
}
