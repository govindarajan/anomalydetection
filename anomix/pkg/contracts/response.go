package contracts

import (
	"net/http"
)

// BaseResponse defines the common structure of standard response
// This contains all but one field of final response
// All the request response will embed this structure
// and should have "response" key
type BaseResponse struct {
	RequestID string    `json:"request_id"`
	Method    string    `json:"method"`
	HTTPCode  int       `json:"http_code"`
	Metadata  *Metadata `json:"metadata,omitempty"`
}

// Metadata defines the structure for metadata
type Metadata struct {
	Failed   *uint64     `json:"failed,omitempty"`
	PageSize *uint64     `json:"page_size,omitempty"`
	Page     *uint64     `json:"page,omitempty"`
	Total    *uint64     `json:"total,omitempty"`
	Success  *uint64     `json:"success,omitempty"`
	Custom   interface{} `json:"custom,omitempty"`
}

// SingleResponse struct defines response for one resource in the request
//go:generate goscinny validator -t SingleResponse -f $GOFILE
type SingleResponse struct {
	Code      *int    `json:"code"`
	ErrorData *Error  `json:"error_data"`
	Msg       *string `json:"msg,omitempty"`
	Status    *string `json:"status" enum:"failure|success"`
}

// SetMetadata setter for Metadata in Response class
func (res *BaseResponse) SetMetadata(metadata *Metadata) Response {
	if metadata == nil {
		return res
	}
	res.Metadata = metadata
	return res
}

// SetRequestID sets requst id for response
func (res *BaseResponse) SetRequestID(requestID string) Response {
	res.RequestID = requestID
	return res
}

// SetHTTPCode stets the http code
func (res *BaseResponse) SetHTTPCode(code int) Response {
	res.HTTPCode = code
	return res
}

// SetMethod  sets http method
func (res *BaseResponse) SetMethod(method string) Response {
	res.Method = method
	return res
}

// SetErrorData setter for ErrorData in Response class
func (res *SingleResponse) SetErrorData(err *Error) *SingleResponse {
	statusCode := http.StatusOK
	statusString := "success"
	if err != nil {
		statusCode = err.HTTPCode
		statusString = "failure"
	}
	res.ErrorData = err
	res.Status = &statusString
	res.Code = &statusCode
	return res
}

// Response defines interface for a valid response
type Response interface {
	SetHTTPCode(int) Response
	SetMethod(string) Response
	SetRequestID(string) Response
	SetMetadata(*Metadata) Response
}
