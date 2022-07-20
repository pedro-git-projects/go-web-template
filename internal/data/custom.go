package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// custumizing a JSON response for a type

var ErrInvalidCustomFormat = errors.New("invalid custom format")

type Custom int32

// Custom MarshalJSON method for type Custom
func (r *Custom) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r) // Generating the string as needed
	// A JSON string must be quoted so:
	quoted := strconv.Quote(jsonValue)
	return []byte(quoted), nil
}

// Setting up custom unmarshalling for our custom field
func (r *Custom) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidCustomFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidCustomFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidCustomFormat
	}

	*r = Custom(i)
	return nil
}
