package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/telepathy/telepathy-server/pkg/app"
	"github.com/telepathy/telepathy-server/pkg/e"
	"github.com/telepathy/telepathy-server/service/team_service"
)

type TeamInfo struct {
	TeamName    string `json:"team_name" binding:"required"`
	ManagerName string `json:"manager_name" binding:"required"`
	Formation   string `json:"formation" binding:"required"`
}

// @Summary Get team list
// @Description 구단정보 가져오기
// @Tags 팀관련 API
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/teamlist [get]
func GetTeamList(c *gin.Context) {
	appG := app.Gin{C: c}

	teamList, err := team_service.GetAllTeamList()
	if err != nil {
		appG.Response(e.ERROR_DB, nil)
		return
	}
	data := make(map[string]interface{})
	data["list"] = teamList
	appG.Response(e.SUCCESS, data)
}
