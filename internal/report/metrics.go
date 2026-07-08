package report

import (
	"fmt"
	"sort"

	"grafana-dashboard-parser/internal/model"
)


func PrintMetricsOverview(catalog *model.Catalog) {

	var metrics []string

	for metric := range catalog.Metrics {
		metrics = append(metrics, metric)
	}

	sort.Strings(metrics)

	for _, metric := range metrics {

		fmt.Printf("\n=================================================\n")
		fmt.Println(metric)
		fmt.Printf("=================================================\n")

		for _, usage := range catalog.Metrics[metric] {

			fmt.Printf("Dashboard : %s\n", usage.Dashboard)
			fmt.Printf("Panel     : %s\n", usage.Panel)
			fmt.Printf("Query     : %s\n", usage.Query)
			fmt.Println()
		}
	}
}
