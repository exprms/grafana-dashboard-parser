package main

import (
	"fmt"
	"log"
	"sort"

	"grafana-dashboard-parser/internal/analyzer"
	"grafana-dashboard-parser/internal/dashboard"
	"grafana-dashboard-parser/internal/model"
)

func main() {

	files, err := dashboard.ReadDashboards("./dashboards")
	if err != nil {
		log.Fatal(err)
	}

	var dashboards []*model.DashboardInfo

	for file, data := range files {

		d, err := dashboard.ParseDashboard(data)
		if err != nil {
			log.Println(err)
			continue
		}

		info, err := analyzer.BuildDashboardInfo(file, d)
		if err != nil {
			log.Println(err)
			continue
		}

		dashboards = append(dashboards, info)
	}

	printDashboardOverview(dashboards)
}

func printDashboardOverview(dashboards []*model.DashboardInfo) {

	for _, dashboard := range dashboards {

		fmt.Printf("\n%s\n", dashboard.Title)
		fmt.Printf("%s\n", dashboard.File)

		for _, query := range dashboard.Queries {

			fmt.Printf("\nPanel: %s\n", query.PanelTitle)
			fmt.Printf("Query: %s\n", query.Expr)

			sort.Strings(query.Metrics)

			fmt.Println("Metrics:")

			for _, metric := range query.Metrics {
				fmt.Printf("  - %s\n", metric)
			}
		}
	}
}
