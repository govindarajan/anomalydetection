package datapoint

import (
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
	"github.com/govindarajan/anomalydetection/detector"
	"github.com/govindarajan/anomalydetection/log"
)

//Create create metrics
func Create(ctx types.Context, request contracts.CreateDataPointRequest) (
	*contracts.CreateDataPointResponse, *contracts.Error) {

	isAnomaly, score, err := detector.DetectAnomaly(*request.Name, *request.Time, *request.Value)
	if err != nil {
		// TODO: Don't compare like this. Pass error code
		if err.Error() == "METRIC_NOT_FOUND" {
			return nil, contracts.ErrMetricNotFound()
		}

		log.Error("Error Creating Datapoint:", err)
		return nil, contracts.ErrInternalServerError()
	}

	sRes := &contracts.SingleDataPointResponse{}
	sRes.Name = *request.Name
	sRes.Time = *request.Time
	sRes.Value = *request.Value
	sRes.IsAnomaly = isAnomaly
	sRes.Score = score

	res := &contracts.CreateDataPointResponse{}
	res.ResponseData = append(res.ResponseData, sRes)

	return res, nil
}
