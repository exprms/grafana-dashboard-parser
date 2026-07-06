// internal/model/dashboard.go

package model

type Dashboard struct {
	Title  string  `json:"title"`
	Panels []Panel `json:"panels"`
}

type Panel struct {
	Title   string   `json:"title"`
	Panels  []Panel  `json:"panels"`
	Targets []Target `json:"targets"`
}

type Target struct {
	Expr string `json:"expr"`
}
