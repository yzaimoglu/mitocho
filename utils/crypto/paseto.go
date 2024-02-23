package crypto

import (
	"aidanwoods.dev/go-paseto"
	"github.com/yzaimoglu/mitocho/data/types"
	"time"
)

type Subject string

const (
	Access  Subject = "access"
	Refresh Subject = "refresh"
	Issuer  string  = "mitocho"
)

type PasetoGen struct {
	PrivateKey paseto.V4AsymmetricSecretKey
	PublicKey  paseto.V4AsymmetricPublicKey
}

func NewPasetoGen() *PasetoGen {
	p := &PasetoGen{}
	p.Generate()
	return p
}

func FromPrivateKey(privateKeyHex string) *PasetoGen {
	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		return nil
	}
	publicKey := privateKey.Public()

	p := &PasetoGen{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	return p
}

func (p *PasetoGen) Generate() {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	p.PrivateKey = privateKey
	p.PublicKey = publicKey
}

func (p *PasetoGen) PrivateKeyHex() string {
	return p.PrivateKey.ExportHex()
}

func (p *PasetoGen) PublicKeyHex() string {
	return p.PublicKey.ExportHex()
}

func (p *PasetoGen) CreateToken(audience string, issuer string, expiration time.Time, subject Subject, notBefore time.Time, issuedAt time.Time, payload any) (string, error) {
	token := paseto.NewToken()
	token.SetAudience(audience)
	token.SetIssuer(issuer)
	token.SetExpiration(expiration)
	token.SetSubject(string(subject))
	token.SetNotBefore(notBefore)
	token.SetIssuedAt(issuedAt)

	if err := token.Set("payload", payload); err != nil {
		return "", err
	}

	signed := token.V4Sign(p.PrivateKey, nil)
	return signed, nil

}

func (p *PasetoGen) CreateAccessToken(site types.Site, payload any) (string, error) {
	audience := site.Domain
	expiration := time.Now().Add(time.Hour * 1)
	subject := Access
	notBefore := time.Now()
	issuedAt := time.Now()

	return p.CreateToken(audience, Issuer, expiration, subject, notBefore, issuedAt, payload)
}

func (p *PasetoGen) CreateRefreshToken(site types.Site, payload any) (string, error) {
	audience := site.Domain
	expiration := time.Now().Add(time.Hour * 24 * 7)
	subject := Refresh
	notBefore := time.Now()
	issuedAt := time.Now()

	return p.CreateToken(audience, Issuer, expiration, subject, notBefore, issuedAt, payload)

}
