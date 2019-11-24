package contracts

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"strconv"

	"github.com/ajg/form"
)

// Decoder defines interface for a function
type Decoder interface {
	Decode(interface{}) error
}

// ContentDecoder returns decoder given content type
// if the content type given does not align to the expectations it will return json decoder
func ContentDecoder(contentType string) func(r io.Reader) Decoder {
	switch contentType {
	case "application/json":
		return func(r io.Reader) Decoder { return json.NewDecoder(r) }
	case "application/xml", "text/xml":
		return func(r io.Reader) Decoder { return xml.NewDecoder(r) }
	case "application/x-www-form-urlencoded", "multipart/form-data":
		return func(r io.Reader) Decoder { return form.NewDecoder(r) }
	default:
		return func(r io.Reader) Decoder { return json.NewDecoder(r) }
	}
}

// ParseToInt parses a given string and returns
// int64 value of the string
func ParseToInt(data string) (int64, error) {
	return strconv.ParseInt(data, 0, 64)
}

// ParseToUint parses a given string and returns
// uint64 value of the string
func ParseToUint(data string) (uint64, error) {
	return strconv.ParseUint(data, 0, 64)
}

// ParseToFloat parses a given string and returns
// float64 value of the string
func ParseToFloat(data string) (float64, error) {
	return strconv.ParseFloat(data, 0)
}

// ParseToBool parses a given string and returns
// bool value of the string
func ParseToBool(data string) (bool, error) {
	return strconv.ParseBool(data)
}
