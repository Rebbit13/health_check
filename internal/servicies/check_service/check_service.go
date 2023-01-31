package check_service

import (
	"health_check/internal/domain"
)

type CheckService struct{}

func (c CheckService) CheckSites(sitesToCheck []*domain.SiteToCheck) ([]*domain.SiteChecked, error) {
	formedReport := make([]*domain.SiteChecked, 0)
	reportBuffer := make(chan *domain.SiteChecked, len(sitesToCheck))
	for _, siteToCheck := range sitesToCheck {
		s := newSite(siteToCheck)
		go s.initCheck(&reportBuffer)
	}
	for i := 0; i < len(sitesToCheck); i++ {
		result := <-reportBuffer
		formedReport = append(formedReport, result)
	}
	return formedReport, nil
}
