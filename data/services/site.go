package services

import "github.com/yzaimoglu/mitocho/data/types"

func (svc *Service) GetSiteByAccessToken(accessToken types.AccessToken) (*types.Site, error) {
	site := &types.Site{}
	tx := svc.DB.Gorm.Where("access_token = ?", accessToken).First(site)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return site, nil
}
