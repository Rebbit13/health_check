package output_service

import "health_check/internal/domain"

type RepositoryOutput struct {
	repository Repository
}

func NewRepositoryOutput(repository Repository) *RepositoryOutput {
	return &RepositoryOutput{repository: repository}
}

func (r RepositoryOutput) saveSiteChecks(site *domain.SiteChecked) error {
	return r.repository.SaveResult(site)
}

func (r RepositoryOutput) SendToOutput(report []*domain.SiteChecked) error {
	for _, site := range report {
		err := r.saveSiteChecks(site)
		if err != nil {
			return err
		}
	}
	return nil
}
