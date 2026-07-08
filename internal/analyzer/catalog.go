package analyzer

import (
	"grafana-dashboard-parser/internal/model"
)

func BuildCatalog(dashboards []*model.DashboardInfo) *model.Catalog {

	catalog := &model.Catalog{
		Dashboards: dashboards,
		Metrics:    make(map[string][]model.MetricUsage),
	}

	for _, dashboard := range dashboards {

		for _, query := range dashboard.Queries {

			for _, metric := range query.Metrics {

				catalog.Metrics[metric] = append(
					catalog.Metrics[metric],
					model.MetricUsage{
						Dashboard: dashboard.Title,
						Panel:     query.PanelTitle,
						Query:     query.Expr,
					},
				)
			}
		}
	}

	return catalog
}
