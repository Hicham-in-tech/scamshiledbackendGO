package handlers

import (
	"net/http"
	"scamshield-backend/app/models"
	"scamshield-backend/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScanHandler struct {
	db      *gorm.DB
	scanner *services.URLRiskAnalyzer
}

func NewScanHandler(db *gorm.DB, scanner *services.URLRiskAnalyzer) *ScanHandler {
	return &ScanHandler{db: db, scanner: scanner}
}

func (h *ScanHandler) CreateScan(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Analyze URL
	score, riskLevel := h.scanner.AnalyzeURL(req.URL)

	scan := models.Scan{
		UserID:    userID,
		URL:       req.URL,
		Score:     score,
		RiskLevel: riskLevel,
	}

	if err := h.db.Create(&scan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create scan"})
		return
	}

	c.JSON(http.StatusCreated, scan)
}

func (h *ScanHandler) GetScans(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var scans []models.Scan
	if err := h.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&scans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch scans"})
		return
	}

	c.JSON(http.StatusOK, scans)
}

func (h *ScanHandler) GetScan(c *gin.Context) {
	userID := c.GetUint64("user_id")
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var scan models.Scan
	if err := h.db.Where("id = ? AND user_id = ?", id, userID).First(&scan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "scan not found"})
		return
	}

	c.JSON(http.StatusOK, scan)
}
