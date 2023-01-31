package check_service

import (
	"health_check/internal/domain"
	"net/http"
)

type siteCheck struct {
	title  string
	passed bool
}

type site struct {
	url        string
	needToPass int
	checks     []string
	report     []*siteCheck
}

func newSite(siteToCheck *domain.SiteToCheck) *site {
	return &site{url: siteToCheck.Url, needToPass: siteToCheck.CheckCount, checks: siteToCheck.Checks}
}

func (s site) initCheck(report *chan *domain.SiteChecked) {
	response, err := http.Get(s.url)
	for _, check := range s.checks {
		passed := false
		if err == nil {
			passed = checksMap[check].Pass(response)
		}
		s.report = append(
			s.report,
			&siteCheck{
				title:  check,
				passed: passed,
			},
		)
	}
	*report <- s.getResult()
}

func (s site) getPassedChecks() int {
	passed := 0
	for _, check := range s.report {
		if check.passed {
			passed++
		}
	}
	return passed
}

func (s site) getFailedCheckTitles() []string {
	failed := make([]string, 0)
	for _, check := range s.report {
		if !check.passed {
			failed = append(failed, check.title)
		}
	}
	return failed
}

func (s site) getResult() *domain.SiteChecked {
	siteChecked := &domain.SiteChecked{}
	siteChecked.Url = s.url
	if s.getPassedChecks() >= s.needToPass {
		siteChecked.Passed = true
	} else {
		siteChecked.Passed = false
		siteChecked.FailedChecks = s.getFailedCheckTitles()
	}
	return siteChecked
}
