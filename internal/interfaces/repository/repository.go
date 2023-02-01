package repository

import (
	"gorm.io/gorm"
	"health_check/internal/domain"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (g GormRepository) GetLastResults(siteUrls []string) ([]*domain.SiteChecked, error) {
	savedSites := make([]SiteFullCheck, 0)
	formedSites := make([]*domain.SiteChecked, 0)
	subQuery := g.db.Table("site_full_checks").Where("url in ?", siteUrls).Order("created_at DESC")
	g.db.Table("(?) as s", subQuery).Group("url").Find(&savedSites)
	for _, site := range savedSites {
		formedSites = append(
			formedSites,
			&domain.SiteChecked{
				Url:    site.Url,
				Passed: site.Passed,
			},
		)
	}
	return formedSites, nil
}

func (g GormRepository) SaveResults(report []*domain.SiteChecked) error {
	sitesToSave := make([]SiteFullCheck, 0)
	for _, site := range report {
		sitesToSave = append(
			sitesToSave,
			SiteFullCheck{
				Url:    site.Url,
				Passed: site.Passed,
			},
		)
	}
	g.db.Create(&sitesToSave)
	return nil
}
