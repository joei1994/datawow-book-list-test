package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BookPayload struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
	Tag    string `json:"tag"`
}

func (p BookPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Author, validation.Required),
	)
}
