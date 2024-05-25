package route

import (
	"database/sql"
	"price-estimation-api/internal/app"
	"price-estimation-api/internal/handlers"
	"price-estimation-api/internal/module"
	"price-estimation-api/internal/persistance"

	"github.com/gin-gonic/gin"
)


func NewEstimateRoute(env *App.Application, router *gin.RouterGroup, db *sql.DB) {
	er := persistance.NewEstimateRepository(db)
	eu := module.NewEstimateUsecase(er)
	
	estimateHandler := rest.NewEstimateHandler(eu)

	estimateGroup := router.Group("/estimations")

	estimateGroup.POST("/", estimateHandler.CreateEstimate)
	estimateGroup.GET("/", estimateHandler.GetEstimates)
	estimateGroup.GET("/:id", estimateHandler.GetEstimate)
	estimateGroup.DELETE("/:id", estimateHandler.DeleteEstimate)
	estimateGroup.PATCH("/:id", estimateHandler.UpdateEstimate)
}