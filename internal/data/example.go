package data

import (
	"template/internal/validator"
	"time"
)

type Example struct {
	ID        int64     `json:"-"` // omits field
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title,omitempty"` // omits field if empty
	Year      int32     `json:"year,string"`     // forces field to be represented as string
	List      []string  `json:"list"`
	Custom    Custom    `json:"custom"` // custom field
}

func ValidateExample(v validator.Validator, example Example) {
	v.Check(example.Title != "", "title", "must be provided")
	v.Check(len(example.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(example.Year != 0, "year", "must be provided")

	v.Check(example.List != nil, "list", "must be provided")
}
