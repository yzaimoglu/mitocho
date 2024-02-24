package types

const (
	SettingNameSetupFinished = "setup_finished"
	SettingNameSiteLogo      = "site_logo"
	SettingNameSiteFavicon   = "site_favicon"
)

type Setting struct {
	BaseModel
	Name  string `json:"name"`
	Value string `json:"value"`
}
