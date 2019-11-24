package contracts

import "github.com/labstack/echo"

// Extractor defines the interface to be implemented by all contracts
// it defines how to read values from a http request to its fields
type Extractor interface {
	ExtractFromHTTP(echo.Context) *Error
}
