package services

import (
	"github.com/labstack/gommon/log"
	"github.com/yzaimoglu/mitocho/data/types"
)

func (svc *Service) InitialSetup() error {
	if err := svc.CreateInitialSettings(); err != nil {
		return err
	}
	return nil
}

func (svc *Service) FinishSetup(name string, domain string, email string, username string, password string) error {
	if svc.IsSetupFinished() {
		return nil
	}

	siteId, err := svc.CreateInitialSite(name, domain)
	if err != nil {
		return err
	}

	if err := svc.CreateInitialRoles(siteId); err != nil {
		return err
	}

	role, err := svc.GetAdminRole(siteId)
	if err != nil {
		return err
	}

	user := types.NewUser(siteId, username, password, email, role.ID)
	if err = svc.CreateUser(user); err != nil {
		log.Infof("user creation failed: %v", err)
	}

	if err = svc.SetSetting(types.SettingNameSetupFinished, "1"); err != nil {
		return err
	}
	return nil
}
