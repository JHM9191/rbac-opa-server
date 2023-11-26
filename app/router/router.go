package router

import (
	"github.com/gin-gonic/gin"
	"rbac-opa-server-mariadb/config"
)

func Init(init *config.Initialization) *gin.Engine {
	router := gin.Default()
	//router.Use(cors.New(
	//	cors.Config{
	//		AllowOrigins: []string{os.Getenv("WEB_URL")},
	//		AllowMethods: []string{"*"},
	//		AllowHeaders: []string{"*"},
	//		MaxAge:       12 * time.Hour,
	//	}))

	v1 := router.Group("/api/v1")

	eval := v1.Group("eval")
	eval.POST("", init.ApiCtrl.EvaluateRule)

	return router
}
