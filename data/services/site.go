package services

import (
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/data/types"
)

func (svc *Service) DoesMitochoSiteExist() bool {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("mitocho = ?", true).First(site)
	if tx.Error != nil {
		return false
	}
	return true
}

func (svc *Service) CreateInitialSite(name string, domain string) error {
	if !svc.IsSetupFinished() && !svc.DoesMitochoSiteExist() {
		site := types.NewMitochoSite(name, domain)
		tx := svc.DB.Gorm.Create(site)
		if tx.Error != nil {
			return tx.Error
		}
		return nil
	}
	return nil
}

func (svc *Service) CreateSiteIfNotExists(name string, domain string) error {
	site := types.NewSite(name, domain)
	tx := svc.DB.Gorm.Where("name = ?", name).First(site)
	if tx.Error != nil {
		tx = svc.DB.Gorm.Create(site)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (svc *Service) GetSiteById(id uuid.UUID) (*types.Site, error) {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("id = ?", id.String()).
		First(site)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return site, nil
}

func (svc *Service) GetSiteByAccessToken(accessToken types.AccessToken) (*types.Site, error) {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("access_token = ?", accessToken).First(site)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return site, nil
}

func (svc *Service) GetSiteByName(name string) (*types.Site, error) {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("name = ?", name).First(site)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return site, nil
}

func (svc *Service) GetSiteByDomain(domain string) (*types.Site, error) {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("JSON_CONTAINS(domains, ?, '$')", `"`+domain+`"`).First(site)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return site, nil
}
