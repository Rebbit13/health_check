package domain

type Check struct {
	Title string
}

type SiteToCheck struct {
	Url        string
	Checks     []string
	CheckCount int `json:"min_checks_cnt"`
}

type SiteChecked struct {
	Url          string
	Passed       bool
	FailedChecks []string
}
