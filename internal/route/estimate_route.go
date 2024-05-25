package route

import (
	App "price-estimation-api/internal/app"
	rest "price-estimation-api/internal/handlers"
	"price-estimation-api/internal/ports"
	"time"

	"github.com/gin-gonic/gin"
)


func NewEstimateRoute(env *App.Application, timeout time.Duration, group *gin.Engine)  {
	er := ports.NewEstimateRepository(env.DB)
	eu := ports.NewEstimateUsecase(er)
	eh := rest.NewEstimateHandler(eu)	
}