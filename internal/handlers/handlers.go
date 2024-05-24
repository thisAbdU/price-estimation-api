package handlers

import (
    "net/http"
    "price-estimation-api/internal/domain"
    "price-estimation-api/internal/ports"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Handlers struct {
    estimateRepo ports.EstimateRepository
}

func NewHandlers(repo ports.EstimateRepository) *Handlers {
    return &Handlers{estimateRepo: repo}
}

func (h *Handlers) CreateEstimate(c *gin.Context) {
    var estimate domain.Estimate
    if err := c.ShouldBindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newEstimate, err := h.estimateRepo.CreateEstimate(estimate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newEstimate)
}

func (h *Handlers) GetEstimates(c *gin.Context) {
    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit <= 0 {
        limit = 10
    }

    offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
    if err != nil || offset < 0 {
        offset = 0
    }

    estimates, err := h.estimateRepo.GetEstimates()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimates)
}

func (h *Handlers) GetEstimate(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    estimate, err := h.estimateRepo.GetEstimateByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimate)
}

func (h *Handlers) UpdateEstimate(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var estimate domain.Estimate
    if err := c.ShouldBindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    estimate.ID = id

    updatedEstimate, err := h.estimateRepo.UpdateEstimate(estimate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedEstimate)
}

func (h *Handlers) DeleteEstimate(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = h.estimateRepo.DeleteEstimate(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Estimate deleted successfully"})
}
