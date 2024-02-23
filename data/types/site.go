package types

import (
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type AccessToken string
type PasetoKeyHex string

// TODO: SiteGroup
type SiteGroup struct {
	BaseModel
}

type Site struct {
	BaseModel
	Name        string       `json:"name"`
	Domain      string       `json:"domain"`
	Active      bool         `json:"active"`
	AccessToken AccessToken  `json:"access_token"`
	PublicKey   PasetoKeyHex `json:"public_key"`
	PrivateKey  PasetoKeyHex `json:"private_key"`
}

func NewSite(name string, domain string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()

	return &Site{
		Name:        name,
		Domain:      domain,
		Active:      true,
		AccessToken: accessToken,
		PublicKey:   PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:  PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}

type SiteRole struct {
	Role
	Site   Site      `json:"site"`
	SiteID uuid.UUID `gorm:"type:char(36)" json:"-"`
}

type SiteUser struct {
	User
	Site   Site      `json:"site"`
	SiteID uuid.UUID `gorm:"type:char(36)" json:"-"`
}
