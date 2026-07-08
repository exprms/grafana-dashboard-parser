package report

import (
	"fmt"
	"sort"

	"grafana-dashboard-parser/internal/model"
)

func PrintDashboardOverview(catalog *model.Catalog) {

	for _, dashboard := range catalog.Dashboards {

		fmt.Printf("\n=== %s ===\n", dashboard.Title)
		fmt.Printf("File: %s\n", dashboard.File)

		for _, query := range dashboard.Queries {

			fmt.Printf("\nPanel: %s\n", query.PanelTitle)
			fmt.Printf("Query: %s\n", query.Expr)

			metrics := append([]string(nil), query.Metrics...)
			sort.Strings(metrics)

			fmt.Println("Metrics:")

			for _, metric := range metrics {
				fmt.Printf("  - %s\n", metric)
			}
		}

		fmt.Println()
	}
}
