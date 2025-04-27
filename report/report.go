package report

import (
	"encoding/json"
	"fmt"
	"os"
)
// dns resolving + reports
type DNSResult struct {
	Server string   `json:"server"`
	IPs    []string `json:"ips"`
	Status string   `json:"status"`
}

func GenerateReport(results []DNSResult) {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON report:", err)
		return
	}

	err = os.WriteFile("report.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON report:", err)
		return
	}

	fmt.Println("Generated JSON report: report.json")
}
