package types

import (
	"encoding/json"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type AccessToken string
type PasetoKeyHex string
type LogoB64 string
type SiteColor string
type Domains json.RawMessage
type Callbacks json.RawMessage

const (
	SiteColorGreen SiteColor = "green"
	SiteColorBlue  SiteColor = "blue"
)

type Site struct {
	BaseModel
	Name           string       `json:"name"`
	Description    *string      `json:"description,omitempty"`
	Domains        Domains      `gorm:"type:json" json:"domains"`
	Callbacks      *Callbacks   `gorm:"type:json" json:"callbacks"`
	Active         bool         `json:"active"`
	Mitocho        bool         `json:"mitocho"`
	TermsOfService *string      `json:"tos,omitempty"`
	PrivacyPolicy  *string      `json:"privacy_policy,omitempty"`
	Imprint        *string      `json:"imprint,omitempty"`
	Color          *SiteColor   `json:"color,omitempty"`
	Logo           *LogoB64     `json:"logo,omitempty"`
	AccessToken    AccessToken  `json:"access_token"`
	PublicKey      PasetoKeyHex `json:"public_key"`
	PrivateKey     PasetoKeyHex `json:"private_key"`
}

func NewDomain(domain string) Domains {
	domains := []string{domain}
	domainJson, _ := json.Marshal(domains)
	return domainJson
}

func NewDomains(domains []string) Domains {
	domainsJson, _ := json.Marshal(domains)
	return domainsJson
}

func NewCallback(callback string) Callbacks {
	callbacks := []string{callback}
	callbackJson, _ := json.Marshal(callbacks)
	return callbackJson
}

func NewCallbacks(callbacks []string) Callbacks {
	callbacksJson, _ := json.Marshal(callbacks)
	return callbacksJson
}

func NewMitochoSite(name string, domain string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()
	privacyPolicy := "/privacy"
	imprint := "/imprint"
	terms := "/terms"
	color := SiteColorGreen

	return &Site{
		Name:           name,
		Domains:        NewDomain(domain),
		Active:         true,
		Mitocho:        true,
		AccessToken:    AccessToken(accessToken),
		PublicKey:      PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:     PasetoKeyHex(pasetoGen.PrivateKeyHex()),
		PrivacyPolicy:  &privacyPolicy,
		Imprint:        &imprint,
		TermsOfService: &terms,
		Color:          &color,
	}
}

func NewSite(name string, privacyPolicy string, imprint string, termsOfService string, domain string, color SiteColor, callback string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()
	callbackObj := NewCallback(callback)

	return &Site{
		Name:           name,
		Domains:        NewDomain(domain),
		Callbacks:      &callbackObj,
		Active:         true,
		Mitocho:        false,
		PrivacyPolicy:  &privacyPolicy,
		Imprint:        &imprint,
		TermsOfService: &termsOfService,
		Color:          &color,
		AccessToken:    AccessToken(accessToken),
		PublicKey:      PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:     PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}

func NewSiteGroup(name string, privacyPolicy string, imprint string, termsOfService string, color SiteColor, callback string, domains ...string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()
	callbackObj := NewCallback(callback)

	return &Site{
		Name:           name,
		Domains:        NewDomains(domains),
		Callbacks:      &callbackObj,
		Active:         true,
		Mitocho:        false,
		PrivacyPolicy:  &privacyPolicy,
		Imprint:        &imprint,
		TermsOfService: &termsOfService,
		Color:          &color,
		AccessToken:    AccessToken(accessToken),
		PublicKey:      PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:     PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}
