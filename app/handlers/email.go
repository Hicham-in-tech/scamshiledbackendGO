package handlers

import (
	"net/http"
	"scamshield-backend/app/models"
	"scamshield-backend/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmailReviewHandler struct {
	db       *gorm.DB
	reviewer *services.EmailSecurityReviewer
}

func NewEmailReviewHandler(db *gorm.DB, reviewer *services.EmailSecurityReviewer) *EmailReviewHandler {
	return &EmailReviewHandler{db: db, reviewer: reviewer}
}

func (h *EmailReviewHandler) ReviewEmail(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		InputText string `json:"input_text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Review email
	score, riskLevel := h.reviewer.ReviewEmail(req.InputText)

	review := models.EmailReview{
		UserID:    userID,
		InputText: req.InputText,
		Score:     score,
		RiskLevel: riskLevel,
	}

	if err := h.db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *EmailReviewHandler) GetReviews(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var reviews []models.EmailReview
	if err := h.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *EmailReviewHandler) GetReview(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var review models.EmailReview
	if err := h.db.Where("id = ? AND user_id = ?", id, userID).First(&review).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}

	c.JSON(http.StatusOK, review)
}
