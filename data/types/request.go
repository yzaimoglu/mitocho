package types

type PostLoginRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8,max=128"`
	PasswordRepeat string `json:"password_repeat" validate:"required,min=8,max=128"`
}

type PostLoginTOTPRequest struct {
	PostLoginRequest
	TOTP string `json:"totp" validate:"required,max=6"`
}

type PostSitePublicKeyRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type PostInstallFinishRequest struct {
	Name           string `json:"name" validate:"required"`
	Domain         string `json:"domain" validate:"required,hostname"`
	Email          string `json:"email" validate:"required,email"`
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required,min=8,max=128"`
	PasswordRepeat string `json:"password_repeat" validate:"required,min=8,max=128"`
}
