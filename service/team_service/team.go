package team_service

import "github.com/telepathy/telepathy-server/models"

func GetAllTeamList() ([]*models.Team, error) {
	return models.GetAllTeamList()
}
