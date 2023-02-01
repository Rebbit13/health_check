package output_service

import (
	"fmt"
	"health_check/internal/domain"
	"strings"
)

type StdOutput struct{}

func (s StdOutput) getFailedChecksTitles(site *domain.SiteChecked) string {
	failedCheckTitles := make([]string, 0)
	for _, check := range site.Checks {
		if !check.Pass {
			failedCheckTitles = append(failedCheckTitles, check.Title)
		}
	}
	return strings.Join(failedCheckTitles, ", ")
}

func (s StdOutput) SendToOutput(report []*domain.SiteChecked) error {
	textReport := ""

	for _, site := range report {
		if site.Passed {
			textReport += fmt.Sprintf("%s: ok\n", site.Url)
		} else {

			textReport += fmt.Sprintf("%s: failed (%s)\n", site.Url, s.getFailedChecksTitles(site))
		}

	}
	fmt.Println(textReport)
	return nil
}
