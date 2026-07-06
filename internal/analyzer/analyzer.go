package analyzer

import (
	"github.com/prometheus/prometheus/promql/parser"

	"grafana-dashboard-parser/internal/model"
)

func BuildDashboardInfo(file string, dashboard *model.Dashboard) (*model.DashboardInfo, error) {

	info := &model.DashboardInfo{
		File:  file,
		Title: dashboard.Title,
	}

	var walk func([]model.Panel)

	walk = func(panels []model.Panel) {

		for _, panel := range panels {

			for _, target := range panel.Targets {

				if target.Expr == "" {
					continue
				}

				metrics, err := extractMetrics(target.Expr)
				if err != nil {
					continue
				}

				info.Queries = append(info.Queries, model.QueryInfo{
					PanelTitle: panel.Title,
					Expr:       target.Expr,
					Metrics:    metrics,
				})
			}

			walk(panel.Panels)
		}
	}

	walk(dashboard.Panels)

	return info, nil
}

func extractMetrics(expr string) ([]string, error) {

	p := parser.NewParser(parser.Options{})

	ast, err := p.ParseExpr(expr)
	if err != nil {
		return nil, err
	}

	metricSet := make(map[string]struct{})

	parser.Inspect(ast, func(node parser.Node, path []parser.Node) error {

		switch n := node.(type) {

		case *parser.VectorSelector:
			if n.Name != "" {
				metricSet[n.Name] = struct{}{}
			}

		case *parser.MatrixSelector:
			if vs, ok := n.VectorSelector.(*parser.VectorSelector); ok {
				if vs.Name != "" {
					metricSet[vs.Name] = struct{}{}
				}
			}
		}

		return nil
	})

	metrics := make([]string, 0, len(metricSet))

	for metric := range metricSet {
		metrics = append(metrics, metric)
	}

	return metrics, nil
}
