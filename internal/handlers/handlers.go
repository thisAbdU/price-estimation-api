package handlers

import (
    "net/http"
    "price-estimation-api/internal/domain"
    "price-estimation-api/internal/adapters/db"
    "price-estimation-api/internal/database"
    "github.com/gin-gonic/gin"
)

var repo = db.NewPostgresEstimateRepository(database.DB)

func CreateEstimate(c *gin.Context) {
    var estimate domain.Estimate
    if err := c.ShouldBindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    estimate, err := repo.CreateEstimate(estimate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, estimate)
}

func GetEstimates(c *gin.Context) {
    estimates, err := repo.GetEstimates()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimates)
}

func GetEstimate(c *gin.Context) {
    id := c.Param("id")
    estimate, err := repo.GetEstimateByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimate)
}

func UpdateEstimate(c *gin.Context) {
    id := c.Param("id")
    var estimate domain.Estimate
    if err := c.ShouldBindJSON(&estimate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    estimate.ProductID = id
    updatedEstimate, err := repo.UpdateEstimate(estimate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedEstimate)
}

func DeleteEstimate(c *gin.Context) {
    id := c.Param("id")
    err := repo.DeleteEstimate(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Estimate deleted successfully"})
}
