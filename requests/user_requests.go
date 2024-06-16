package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type BasicAuth struct {
	Email    string `json:"email" validate:"required" example:"someone@example.com"`
	Password string `json:"password" validate:"required" example:"1234567890"`
}

func (ba BasicAuth) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Email, is.Email),
		validation.Field(&ba.Password, validation.Length(8, 0)),
	)
}

type RegisterRequest struct {
	BasicAuth
	Name string `json:"name" validate:"required" example:"someone"`
}

func (rr *RegisterRequest) Validate() error {
	err := rr.BasicAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(rr,
		validation.Field(&rr.Name, validation.Required),
	)
}

type LoginRequest struct {
	BasicAuth
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required"`
}
