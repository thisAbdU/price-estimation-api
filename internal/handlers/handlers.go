package rest

import (
	"net/http"
	"price-estimation-api/internal/domain"
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
    var estimate domain.Estimate

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

    estimate, err := h.EstimateUsecase.GetEstimate(c.Request.Context(), id)
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

    var estimate domain.Estimate

    if err := c.BindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedEstimate, err := h.EstimateUsecase.UpdateEstimate(c.Request.Context(), id)
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

    err = h.EstimateUsecase.DeleteEstimate(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) GetEstimates(c *gin.Context) {
    estimates, err := h.EstimateUsecase.GetEstimates(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, estimates)
}

