package routers

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/telepathy/telepathy-server/docs"
	v1 "github.com/telepathy/telepathy-server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// apiv1 := r.Group("/v1/app")
	// apiv1.Use(app.HeaderValid())
	// {
	// 	apiv1.GET("/api/teamlist", v1.GetTeamList)
	// 	apiv1.GET("/api/playerlist/:team_id", v1.GetPlayerList)
	// }

	r.GET("/api/teamlist", v1.GetTeamList)
	//r.POST("/api/team", api.AddTeamInfo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
