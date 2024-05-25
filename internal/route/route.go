package route

import (
	"database/sql"
	"price-estimation-api/internal/app"

	"github.com/gin-gonic/gin"
)

func Setup(env *App.Application, group *gin.Engine, db *sql.DB)  {
	publicRouter := group.Group("")

	NewEstimateRoute(env, publicRouter, db)
}

