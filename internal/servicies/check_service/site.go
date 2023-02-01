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
	report     []domain.Check
}

func newSite(siteToCheck *domain.SiteToCheck) *site {
	return &site{url: siteToCheck.Url, needToPass: siteToCheck.CheckCount, checks: siteToCheck.Checks}
}

func (s site) initCheck(report *chan *domain.SiteChecked) {
	response, _ := http.Get(s.url)
	defaultCheck := connection{}.Pass(response)
	if defaultCheck.Pass {
		for _, check := range s.checks {
			s.report = append(s.report, checksMap[check].Pass(response))
		}
	} else {
		s.report = append(s.report, defaultCheck)
	}
	*report <- s.getResult()
}

func (s site) getPassedChecks() int {
	passed := 0
	for _, check := range s.report {
		if check.Pass {
			passed++
		}
	}
	return passed
}

func (s site) getResult() *domain.SiteChecked {
	siteChecked := &domain.SiteChecked{}
	siteChecked.Url = s.url
	siteChecked.Checks = s.report
	if s.getPassedChecks() >= s.needToPass {
		siteChecked.Passed = true
	} else {
		siteChecked.Passed = false
	}
	return siteChecked
}
