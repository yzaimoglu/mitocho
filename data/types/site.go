package types

import (
	"encoding/json"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type AccessToken string
type PasetoKeyHex string
type Domains json.RawMessage
type Callbacks json.RawMessage

type Site struct {
	BaseModel
	Name          string       `json:"name"`
	Domains       Domains      `gorm:"type:json" json:"domains"`
	Callbacks     *Callbacks   `gorm:"type:json" json:"callbacks"`
	Active        bool         `json:"active"`
	Mitocho       bool         `json:"mitocho"`
	PrivacyPolicy string       `json:"privacy_policy"`
	Imprint       string       `json:"imprint"`
	AccessToken   AccessToken  `json:"access_token"`
	PublicKey     PasetoKeyHex `json:"public_key"`
	PrivateKey    PasetoKeyHex `json:"private_key"`
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

	return &Site{
		Name:          name,
		Domains:       NewDomain(domain),
		Active:        true,
		Mitocho:       true,
		AccessToken:   AccessToken(accessToken),
		PublicKey:     PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:    PasetoKeyHex(pasetoGen.PrivateKeyHex()),
		PrivacyPolicy: "/privacy",
		Imprint:       "/imprint",
	}
}

func NewSite(name string, privacyPolicy string, imprint string, domain string, callback string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()
	callbackObj := NewCallback(callback)

	return &Site{
		Name:          name,
		Domains:       NewDomain(domain),
		Callbacks:     &callbackObj,
		Active:        true,
		Mitocho:       false,
		PrivacyPolicy: privacyPolicy,
		Imprint:       imprint,
		AccessToken:   AccessToken(accessToken),
		PublicKey:     PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:    PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}

func NewSiteGroup(name string, privacyPolicy string, imprint string, callback string, domains ...string) *Site {
	accessToken, err := crypto.GenerateAccessToken()
	if err != nil {
		return nil
	}

	pasetoGen := crypto.NewPasetoGen()
	callbackObj := NewCallback(callback)

	return &Site{
		Name:          name,
		Domains:       NewDomains(domains),
		Callbacks:     &callbackObj,
		Active:        true,
		Mitocho:       false,
		PrivacyPolicy: privacyPolicy,
		Imprint:       imprint,
		AccessToken:   AccessToken(accessToken),
		PublicKey:     PasetoKeyHex(pasetoGen.PublicKeyHex()),
		PrivateKey:    PasetoKeyHex(pasetoGen.PrivateKeyHex()),
	}
}
