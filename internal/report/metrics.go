package report

import (
	"fmt"
	"sort"

	"grafana-dashboard-parser/internal/model"
)

type MetricUsage struct {
	Dashboard string
	Panel      string
	Query      string
}

func PrintMetricsOverview(dashboards []*model.DashboardInfo) {

	usage := make(map[string][]MetricUsage)

	for _, dashboard := range dashboards {

		for _, query := range dashboard.Queries {

			for _, metric := range query.Metrics {

				usage[metric] = append(usage[metric], MetricUsage{
					Dashboard: dashboard.Title,
					Panel:     query.PanelTitle,
					Query:     query.Expr,
				})
			}
		}
	}

	var metrics []string

	for metric := range usage {
		metrics = append(metrics, metric)
	}

	sort.Strings(metrics)

	for _, metric := range metrics {

		fmt.Printf("\n=================================================\n")
		fmt.Printf("%s\n", metric)
		fmt.Printf("=================================================\n")

		for _, u := range usage[metric] {

			fmt.Printf("Dashboard : %s\n", u.Dashboard)
			fmt.Printf("Panel     : %s\n", u.Panel)
			fmt.Printf("Query     : %s\n", u.Query)
			fmt.Println()
		}
	}
}
