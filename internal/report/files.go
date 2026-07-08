package report

import (
	"fmt"

	"grafana-dashboard-parser/internal/model"
)

func PrintDashboardFiles(catalog *model.Catalog) {

	fmt.Printf("Found %d dashboards:\n\n", len(catalog.Dashboards))

	for _, dashboard := range catalog.Dashboards {
		fmt.Printf("%s\n", dashboard.File)
	}
}
