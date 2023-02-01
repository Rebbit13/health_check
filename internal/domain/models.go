package domain

type Check struct {
	Title          string
	Pass           bool
	ExpectedResult []byte
	GotResult      []byte
}

type SiteToCheck struct {
	Url        string
	Checks     []string
	CheckCount int `json:"min_checks_cnt"`
}

type SiteChecked struct {
	Url    string
	Passed bool
	Checks []Check
}
