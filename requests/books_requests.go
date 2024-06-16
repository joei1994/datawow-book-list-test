package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BookPayload struct {
	Title  string `json:"title" validate:"required" example:"Harry Potter and The Chamber of Secrets"`
	Author string `json:"author" validate:"required" example:"J.K. Rowling"`
	Genre  string `json:"genre" example:"Fantasy"`
	Year   string `json:"year" example:"1994"`
	Tag    string `json:"tag" example:"Novel"`
}

func (p BookPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.Author, validation.Required),
	)
}
