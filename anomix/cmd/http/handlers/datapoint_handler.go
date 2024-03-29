// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by goscinny handler -p ../../../pkg/contracts DO NOT EDIT
package handlers

// service handler for accounts
import (
	"net/http"
	"strings"

	"github.com/govindarajan/anomalydetection/anomix/internal/checks/throttle"
	"github.com/govindarajan/anomalydetection/anomix/internal/core/metrics/datapoint"
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
	"github.com/labstack/echo"
)

//DataPointHandler is the implementation object of IHandler forDataPointCRUD requests
type DataPointHandler struct{}

// Any handles any method onDataPoint
func (handler DataPointHandler) Any(c echo.Context) error {
	switch c.Get("Method").(string) {
	case http.MethodPost:
		return handler.Create(c)
	}
	return RawResponse(c, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

// Createis the handler for Createrequests toDataPoint
func (DataPointHandler) Create(c echo.Context) error {
	var response *contracts.CreateDataPointResponse
	var err *contracts.Error
	requestID := c.Get("RequestID").(string)
	method := c.Get("Method").(string)
	ctx := types.NewContext(c.Request().Context(), requestID)
	ctx = ctx.Set("path", "datapoint")
	ctx = ctx.Set("action", strings.ToLower(method))
	ctx = ctx.Set("requestIP", c.RealIP())
	req := new(contracts.CreateDataPointRequest)
	if err = req.ExtractFromHTTP(c); err == nil {
		req.Request = &contracts.Request{RequestID: &requestID, Method: &method}
		err = req.Validate()
	}
	if err != nil {
		response = new(contracts.CreateDataPointResponse)
		responseDataItem := new(contracts.SingleDataPointResponse)
		responseDataItem.SingleResponse.SetErrorData(err)
		response.ResponseData = []*contracts.SingleDataPointResponse{responseDataItem}
		return Response(c, response, *responseDataItem.SingleResponse.Code)
	}
	//set the authenticator in context
	ctx = ctx.Set("auth", nil)
	if ok := throttle.BasicThrottler(1000, "1*M").Throttle(ctx); !ok {
		err = contracts.ErrTooManyRequests()
	}
	if err != nil {
		response = new(contracts.CreateDataPointResponse)
		responseDataItem := new(contracts.SingleDataPointResponse)
		responseDataItem.SingleResponse.SetErrorData(err)
		response.ResponseData = []*contracts.SingleDataPointResponse{responseDataItem}
		return Response(c, response, *responseDataItem.SingleResponse.Code)
	}
	response, err = datapoint.Create(ctx, *req)
	if err != nil {
		response = new(contracts.CreateDataPointResponse)
		responseDataItem := new(contracts.SingleDataPointResponse)
		responseDataItem.SingleResponse.SetErrorData(err)
		response.ResponseData = []*contracts.SingleDataPointResponse{responseDataItem}
		return Response(c, response, *responseDataItem.SingleResponse.Code)
	}
	var status int
	if response.Metadata != nil && response.Metadata.Success != nil && response.Metadata.Failed != nil {
		status = http.StatusMultiStatus
	} else {
		status = http.StatusOK
	}
	return Response(c, response, status)
}
