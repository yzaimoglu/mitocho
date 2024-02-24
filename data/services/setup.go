package services

import (
	"github.com/labstack/gommon/log"
	"github.com/yzaimoglu/mitocho/data/types"
)

func (svc *Service) InitialSetup() error {
	if err := svc.CreateInitialSettings(); err != nil {
		return err
	}

	if err := svc.CreateInitialRoles(); err != nil {
		return err
	}

	return nil
}

func (svc *Service) FinishSetup(name string, domain string, email string, username string, password string) error {
	if finished := svc.IsSetupFinished(); finished {
		return nil
	}

	err := svc.CreateInitialSite(name, domain)
	if err != nil {
		return err
	}

	role, err := svc.GetAdminRole()
	if err != nil {
		return err
	}

	user := types.NewUser(username, password, email, role.ID)
	if err = svc.CreateUser(user); err != nil {
		log.Infof("user creation failed: %v", err)
	}

	if err = svc.SetSetting(types.SettingNameSetupFinished, "1"); err != nil {
		return err
	}
	return nil
}
