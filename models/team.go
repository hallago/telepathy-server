package models

import "log"

type Team struct {
	ID          int64  `json:"id"`
	TeamName    string `json:"team_name"`
	ManagerName string `json:"manager_name"`
	Formation   string `json:"formation"`
}

func GetAllTeamList() ([]*Team, error) {
	var teams []*Team
	//teams := []Team{}
	team := Team{}

	rows, err := db.Query("SELECT *FROM team")
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&team.ID, &team.TeamName, &team.ManagerName, &team.Formation)
		if err != nil {
			log.Fatalf(err.Error())
		}
		teams = append(teams, &team)

	}

	return teams, err
}
