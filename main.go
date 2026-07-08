package main

import (
	"log"

	"grafana-dashboard-parser/cmd"
)

func main() {
	if err := cmd.Analyze(); err != nil {
		log.Fatal(err)
	}
}
