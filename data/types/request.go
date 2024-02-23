package types

type PostLoginRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,max=120"`
	PasswordRepeat string `json:"password_repeat" validate:"required,max=120"`
}

type PostLoginTOTPRequest struct {
	PostLoginRequest
	TOTP string `json:"totp" validate:"required"`
}
