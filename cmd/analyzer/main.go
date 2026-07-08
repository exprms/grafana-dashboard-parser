package main

import (
	"flag"
	// "fmt"
	"log"

	"grafana-dashboard-parser/internal/analyzer"
	"grafana-dashboard-parser/internal/dashboard"
	"grafana-dashboard-parser/internal/model"
	"grafana-dashboard-parser/internal/report"
)

func main() {

	dashboardDir := flag.String(
		"dir",
		"./dashboards",
		"Directory containing Grafana dashboards",
	)

	reportType := flag.String(
		"report",
		"",
		"Report to generate (dashboards|metrics)",
	)

	flag.Parse()

	files, err := dashboard.ReadDashboards(*dashboardDir)
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
	
	catalog := analyzer.BuildCatalog(dashboards)

	switch *reportType {

	case "":
		report.PrintDashboardFiles(catalog)

	case report.ReportDashboards:
		report.PrintDashboardOverview(catalog)

	case report.ReportMetrics:
		report.PrintMetricsOverview(catalog)

	default:
		log.Fatalf("Unknown report type: %s", *reportType)
	}
}
