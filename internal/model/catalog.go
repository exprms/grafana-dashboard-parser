package model

type Catalog struct {
	Dashboards []*DashboardInfo
	Metrics    map[string][]MetricUsage
}

type MetricUsage struct {
	Dashboard string
	Panel     string
	Query     string
}
