package output_service

import "health_check/internal/domain"

type Repository interface {
	GetLastResults(siteUrls []string) ([]*domain.SiteChecked, error)
	SaveResults(report []*domain.SiteChecked) error
}

type AlertSender interface {
	SendStatusChanged([]*StatusChanges) error
}
