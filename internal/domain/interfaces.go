package domain

type InputService interface {
	GetSitesToCheck() ([]*SiteToCheck, error)
}

type CheckService interface {
	CheckSites([]*SiteToCheck) ([]*SiteChecked, error)
}

type OutputService interface {
	SendToOutput([]*SiteChecked) error
}
