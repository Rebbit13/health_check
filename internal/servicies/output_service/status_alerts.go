package output_service

import "health_check/internal/domain"

type StatusChanges struct {
	Url            string
	PreviousStatus bool
	CurrentStatus  bool
}

type StatusAlerts struct {
	repository  Repository
	alertSender AlertSender
}

func NewStatusAlerts(repository Repository, alertSender AlertSender) *StatusAlerts {
	return &StatusAlerts{repository: repository, alertSender: alertSender}
}

func (s StatusAlerts) SendToOutput(report []*domain.SiteChecked) error {
	// TODO: refactor this mess
	siteUrls := make([]string, 0)
	for _, site := range report {
		siteUrls = append(siteUrls, site.Url)
	}
	previosReport, err := s.repository.GetLastResults(siteUrls)
	if err != nil {
		return err
	}
	alertsReport := make([]*StatusChanges, 0)
	for _, previousCheck := range previosReport {
		for _, currentCheck := range report {
			if previousCheck.Url == currentCheck.Url && previousCheck.Passed != currentCheck.Passed {
				alertsReport = append(
					alertsReport,
					&StatusChanges{
						Url:            currentCheck.Url,
						PreviousStatus: previousCheck.Passed,
						CurrentStatus:  currentCheck.Passed,
					},
				)
			}
		}
	}
	if len(alertsReport) > 0 {
		return s.alertSender.SendStatusChanged(alertsReport)
	}
	return nil
}
