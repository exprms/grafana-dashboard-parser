package dashboard

import (
	"encoding/json"

	"grafana-dashboard-parser/internal/model"
)

func ParseDashboard(data []byte) (*model.Dashboard, error) {

	var d model.Dashboard

	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
