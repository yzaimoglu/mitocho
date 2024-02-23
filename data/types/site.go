package types

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type AccessToken string
type PasetoKeyHex string
type Domains json.RawMessage

type Site struct {
	BaseModel
	Name        string       `json:"name"`
	Domains     Domains      `json:"domains"`
	Active      bool         `json:"active"`
	AccessToken AccessToken  `json:"access_token"`
	PublicKey   PasetoKeyHex `json:"public_key"`
	PrivateKey  PasetoKeyHex `json:"private_key"`
}

func NewDomain(domain string) Domains {
	domainJson, _ := json.Marshal([]string{domain})
	return Domains(domainJson)
}

func NewDomains(domains []string) Domains {
	domainsJson, _ := json.Marshal(domains)
	return Domains(domainsJson)
}

func NewSite(name string, domain string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()

	return &Site{
		Name:        name,
		Domains:     NewDomain(domain),
		Active:      true,
		AccessToken: AccessToken(accessToken),
		PublicKey:   PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:  PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}

func NewSiteGroup(name string, domains ...string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()

	return &Site{
		Name:        name,
		Domains:     NewDomains(domains),
		Active:      true,
		AccessToken: AccessToken(accessToken),
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
