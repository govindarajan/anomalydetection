package contracts

import "github.com/govindarajan/anomalydetection/model"

// CreateMetricsRequest defines the structure to store request data
// Throttle : Basic,1000,1*M
// Route : /metrics/:metricName
//go:generate goscinny validator -t CreateMetricsRequest -f $GOFILE
//go:generate goscinny extractor -t CreateMetricsRequest -f $GOFILE
type CreateMetricsRequest struct {
	*Request `json:"-"`
	*model.Anomaly
}

// CreateMetricsResponse defines the structure to store response data for CreateItem
type CreateMetricsResponse struct {
	BaseResponse
	ResponseData []*SingleMetricsResponse `json:"response"`
}

// SingleMetricsResponse defines the structure of response for single item
// this is to include in the response item
type SingleMetricsResponse struct {
	SingleResponse
	ResourceData *model.Anomaly `json:"data" required:"true"`
}
