package model

import (
	"bytes"

	"github.com/phuwn/tools/errors"
)

// JSON - represent a json data type
type JSON []byte

// Scan - scan a value into a json variable
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON - implementation of Marshaler interface
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON - implementation of Unmarshaler interface
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull - check if json is null
func (j JSON) IsNull() bool {
	return j == nil || len(j) == 0 || string(j) == "null"
}

// Equals - check if 2 json have the same value
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
