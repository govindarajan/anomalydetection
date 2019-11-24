package contracts

import (
	"fmt"
	"net/http"
)

// Error defines the  error contract
type Error struct {
	HTTPCode        int           `json:"-"`
	Code            uint64        `json:"code"`
	Description     string        `json:"description"`
	InternalMessage []interface{} `json:"-"`
	Message         string        `json:"message"`
}

// Error returns string form of the error
func (err Error) Error() string {
	return fmt.Sprintf("[%s]:[%s] - %s []", err.Message, err.Description, err.InternalMessage)
}
func errorFunc(httpCode int, code uint64, message string) func(internal ...interface{}) *Error {
	if message == "" {
		message = http.StatusText(httpCode)
	}
	return func(internal ...interface{}) *Error {
		var description string
		var ok bool
		if len(internal) > 0 {
			if description, ok = internal[0].(string); ok {
				internal = internal[1:]
			}
		}
		if description == "" {
			description = message
		}
		return &Error{
			Code:            code,
			Message:         message,
			HTTPCode:        httpCode,
			Description:     description,
			InternalMessage: internal,
		}
	}
}

// GetCode returns numeric code corresponding to the error
func (err Error) GetCode() uint64 { return err.Code }

// AddMsg adds internal message to an error
func (err *Error) AddMsg(msg ...interface{}) error {
	if err == nil {
		*err = *ErrInternalServerError(msg...)
		return err
	}
	err.InternalMessage = append(err.InternalMessage, msg...)
	return err
}

// All the error constants would be here
var (
	ErrNotFound                    = errorFunc(http.StatusNotFound, 1000, "")
	ErrBadRequestParametersMissing = errorFunc(http.StatusBadRequest, 1001, "Mandatory Parameter missing")
	ErrBadRequestInvalidParameter  = errorFunc(http.StatusBadRequest, 1002, "Invalid parameter")
	ErrBadRequestInvalidBody       = errorFunc(http.StatusBadRequest, 1007, "Invalid request body,failed parsing ")
	ErrForbidden                   = errorFunc(http.StatusForbidden, 1003, "Authentication failed")
	ErrInternalServerError         = errorFunc(http.StatusInternalServerError, 1004, "")
	ErrTooManyRequests             = errorFunc(http.StatusTooManyRequests, 1006, "")
	ErrMethodNotDefined            = errorFunc(http.StatusMethodNotAllowed, 1005, "Method not implemented")
	ErrResourceConflict            = errorFunc(http.StatusConflict, 1008, "Duplicate resource")
	ErrUnAuthorized                = errorFunc(http.StatusUnauthorized, 1010, "Authentication failed")
	ErrMetricNotFound              = errorFunc(http.StatusNotFound, 1011, "Metric name not found. Please create first before you send datapoints")
)
