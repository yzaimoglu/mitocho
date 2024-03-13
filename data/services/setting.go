package services

import (
	"github.com/yzaimoglu/mitocho/data/types"
)

var (
	InitialSettings = []types.Setting{
		{Global: true, Name: types.SettingNameSetupFinished, Value: "0"},
	}
)

func (svc *Service) CreateSettingIfNotExists(settingToCreate *types.Setting) error {
	setting := &types.Setting{}
	tx := svc.DB.Gorm.Where("name = ?", settingToCreate.Name).First(setting)
	if tx.Error != nil {
		setting.Name = settingToCreate.Name
		setting.Value = settingToCreate.Value
		tx = svc.DB.Gorm.Create(setting)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (svc *Service) GetSetting(name string) (*types.Setting, error) {
	setting := &types.Setting{}
	tx := svc.DB.Gorm.Where("name = ?", name).First(setting)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return setting, nil
}

func (svc *Service) SetSetting(name, value string) error {
	setting := &types.Setting{}
	tx := svc.DB.Gorm.Where("name = ?", name).First(setting)
	if tx.Error != nil {
		return tx.Error
	}

	setting.Value = value
	tx = svc.DB.Gorm.Save(setting)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (svc *Service) CreateInitialSettings() error {
	if finished := svc.IsSetupFinished(); finished {
		return nil
	}

	for _, setting := range InitialSettings {
		if err := svc.CreateSettingIfNotExists(&setting); err != nil {
			return err
		}
	}

	return nil
}

func (svc *Service) IsSetupFinished() bool {
	setting, err := svc.GetSetting(types.SettingNameSetupFinished)
	if err != nil {
		return false
	}

	return setting.Value == "1"
}

func (svc *Service) GetSiteLogo() (string, error) {
	setting, err := svc.GetSetting(types.SettingNameSiteLogo)
	if err != nil {
		return "", err
	}

	return setting.Value, nil
}

func (svc *Service) GetSiteFavicon() (string, error) {
	setting, err := svc.GetSetting(types.SettingNameSiteFavicon)
	if err != nil {
		return "", err
	}

	return setting.Value, nil
}

func (svc *Service) SetSetupFinished() error {
	return svc.SetSetting(types.SettingNameSetupFinished, "1")
}

func (svc *Service) SetSiteLogo(logo string) error {
	return svc.SetSetting(types.SettingNameSiteLogo, logo)
}

func (svc *Service) SetSiteFavicon(favicon string) error {
	return svc.SetSetting(types.SettingNameSiteFavicon, favicon)
}
