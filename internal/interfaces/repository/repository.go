package repository

import "health_check/internal/domain"

type GormRepository struct {
}

func (g GormRepository) GetLastResults(siteUrls []string) ([]*domain.SiteChecked, error) {
	//TODO implement me
	panic("implement me")
}

func (g GormRepository) SaveResults(report []*domain.SiteChecked) error {
	//TODO implement me
	panic("implement me")
}
