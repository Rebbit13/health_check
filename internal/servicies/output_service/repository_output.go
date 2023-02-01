package output_service

import "health_check/internal/domain"

type StatusChanges struct {
	Url            string
	PreviousStatus bool
	CurrentStatus  bool
}

// RepositoryOutput saves current report
// and if there are site status changes sends alert
type RepositoryOutput struct {
	repository    Repository
	alertSender   AlertSender
	currentReport []*domain.SiteChecked
}

func NewStatusAlerts(repository Repository, alertSender AlertSender) *RepositoryOutput {
	return &RepositoryOutput{repository: repository, alertSender: alertSender}
}

func (r RepositoryOutput) saveCurrentReport() error {
	return r.repository.SaveResults(r.currentReport)
}

func (r RepositoryOutput) getSiteUrls() []string {
	siteUrls := make([]string, 0)
	for _, site := range r.currentReport {
		siteUrls = append(siteUrls, site.Url)
	}
	return siteUrls
}

func (r RepositoryOutput) findCurrentCheck(previousCheck *domain.SiteChecked) *domain.SiteChecked {
	for _, currentCheck := range r.currentReport {
		if previousCheck.Url == currentCheck.Url {
			return currentCheck
		}
	}
	return nil
}

func (r RepositoryOutput) getStatusChanges() ([]*StatusChanges, error) {
	siteUrls := r.getSiteUrls()
	previousReport, err := r.repository.GetLastResults(siteUrls)
	if err != nil {
		return nil, err
	}
	alertsReport := make([]*StatusChanges, 0)
	for _, previousCheck := range previousReport {
		currentCheck := r.findCurrentCheck(previousCheck)
		if currentCheck != nil && currentCheck.Passed != previousCheck.Passed {
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
	return alertsReport, nil
}

func (r RepositoryOutput) SendToOutput(report []*domain.SiteChecked) error {
	r.currentReport = report
	alertsReport, err := r.getStatusChanges()
	if err != nil {
		return err
	}
	if len(alertsReport) > 0 {
		err = r.alertSender.SendStatusChanged(alertsReport)
		if err != nil {
			return err
		}
	}
	return r.saveCurrentReport()
}
