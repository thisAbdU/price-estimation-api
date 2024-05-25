package rest

import (
	"net/http"
	"price-estimation-api/internal/db"
	"price-estimation-api/internal/module"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
    EstimateUsecase module.EstimateUsecase
}

func NewEstimateHandler(estimateUsecase module.EstimateUsecase) *Handlers {
    return &Handlers{EstimateUsecase: estimateUsecase}
}


func (h *Handlers) CreateEstimate(c *gin.Context) {
    var estimate db.Estimate

    if err := c.BindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdEstimate, err := h.EstimateUsecase.CreateEstimate(c.Request.Context(), estimate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdEstimate)
}

func (h *Handlers) GetEstimate(c *gin.Context) {
    idStr := c.Param("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
        return
    }

    estimate, err := h.EstimateUsecase.GetEstimate(c.Request.Context(), int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, estimate)
}

func (h *Handlers) UpdateEstimate(c *gin.Context) {
    idStr := c.Param("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
        return
    }

    var estimate db.Estimate

    if err := c.BindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedEstimate, err := h.EstimateUsecase.UpdateEstimate(c.Request.Context(), int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedEstimate)
}

func (h *Handlers) DeleteEstimate(c *gin.Context) {
    idStr := c.Param("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
        return
    }

    err = h.EstimateUsecase.DeleteEstimate(c.Request.Context(), int32(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}

// GetEstimates handles the request to get estimates with pagination.
func (h *Handlers) GetEstimates(c *gin.Context) {
    limitStr := c.DefaultQuery("limit", "10")   // Default limit to 10 if not provided
    offsetStr := c.DefaultQuery("offset", "0") // Default offset to 0 if not provided

    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
        return
    }

    offset, err := strconv.Atoi(offsetStr)
    if err != nil || offset < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
        return
    }

    estimates, err := h.EstimateUsecase.GetEstimates(c.Request.Context(), int32(limit), int32(offset))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, estimates)
}


