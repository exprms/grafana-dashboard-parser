package model

type DashboardInfo struct {
	File    string
	Title   string
	Queries []QueryInfo
}

type QueryInfo struct {
	PanelTitle string
	Expr       string
	Metrics    []string
}
