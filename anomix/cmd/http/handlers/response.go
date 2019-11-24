package handlers

import (
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/metrics"
	"github.com/labstack/echo"
)

// Error defines
func Error(err error) contracts.SingleResponse {
	var response contracts.SingleResponse
	if err == nil {
		response.SetErrorData(nil)
		return response
	}
	switch err := err.(type) {
	case *contracts.Error:
		if err != nil {
			response.SetErrorData(err)
		}
	case contracts.Error:
		response.SetErrorData(&err)
	default:
		er := contracts.ErrInternalServerError("")
		response.SetErrorData(er)
	}
	return response
}

// Response accepts valid response object and sets basic response fields
// httpcode,requestid and method and set
// and it encodes the response in the format that requester `accepts`
// checks the accept header for the same
func Response(c echo.Context, response contracts.Response, httpCode int) error {
	response.SetHTTPCode(httpCode)
	response.SetRequestID(c.Get("RequestID").(string))
	response.SetMethod(c.Get("Method").(string))
	return RawResponse(c, response, httpCode)
}

// RawResponse creates response and responds it
// it encode the respone in the format that requester `accpets`
// checks the accept header for the same
func RawResponse(c echo.Context, response interface{}, httpCode int) error {
	var responseFunc func(int, interface{}) error
	switch c.Request().Header.Get("accept") {
	case "application/json", "text/json", "json":
		responseFunc = c.JSON
	case "text/xml", "application/xml", "xml":
		responseFunc = c.XML
	default:
		responseFunc = c.JSON
	}
	metrics.Request(c.Get("RequestID").(string), c.Request().URL.Path, c.Request().Method, httpCode)
	return responseFunc(httpCode, response)
}
