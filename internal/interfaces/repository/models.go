package repository

import "gorm.io/gorm"

type SiteFullCheck struct {
	gorm.Model
	Url    string
	Passed bool
}
