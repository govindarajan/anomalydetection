package contracts

import "time"

// CreateDataPointRequest defines the structure to store request data
// Throttle : Basic,1000,1*M
// Route : /metrics/:metricName/datapoint
//go:generate goscinny validator -t CreateDataPointRequest -f $GOFILE
//go:generate goscinny extractor -t CreateDataPointRequest -f $GOFILE
type CreateDataPointRequest struct {
	*Request `json:"-"`
	Name     *string    `json:"-" path:"metricName" required:"true"`
	Time     *time.Time `json:"time" required:"true"`
	Value    *float64   `json:"value" required:"true"`
}

// CreateDataPointResponse defines the structure to store response data for Create DP
type CreateDataPointResponse struct {
	BaseResponse
	ResponseData []*SingleDataPointResponse `json:"response"`
}

type SingleDataPointResponse struct {
	SingleResponse
	Name      string    `json:"name"`
	Time      time.Time `json:"time"`
	Value     float64   `json:"value"`
	IsAnomaly bool      `json:"is_anomaly"`
	Score     float64   `json:"score"`
}
