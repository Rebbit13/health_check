package output_service

import (
	"fmt"
	"health_check/internal/domain"
	"strings"
)

type StdOutput struct{}

func (s StdOutput) SendToOutput(report []*domain.SiteChecked) error {
	textReport := ""
	for _, rep := range report {
		if rep.Passed {
			textReport += fmt.Sprintf("%s: ok\n", rep.Url)
		} else {
			textReport += fmt.Sprintf("%s: failed (%s)\n", rep.Url, strings.Join(rep.FailedChecks, ", "))
		}

	}
	fmt.Println(textReport)
	return nil
}
