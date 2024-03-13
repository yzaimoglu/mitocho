package types

import "github.com/google/uuid"

const (
	SettingNameSetupFinished = "setup_finished"
	SettingNameSiteLogo      = "site_logo"
	SettingNameSiteFavicon   = "site_favicon"
)

type Setting struct {
	BaseModel
	Global bool       `json:"global"`
	SiteId *uuid.UUID `gorm:"type:char(36)" json:"-"`
	Site   *Site      `json:"site"`
	Name   string     `json:"name"`
	Value  string     `json:"value"`
}
