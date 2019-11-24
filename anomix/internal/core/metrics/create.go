package metrics

import (
	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
	"github.com/govindarajan/anomalydetection/detector"
	"github.com/govindarajan/anomalydetection/model"
)

//Create create metrics
func Create(ctx types.Context, request contracts.CreateMetricsRequest) (*contracts.CreateMetricsResponse, *contracts.Error) {

	if request.Anomaly == nil {
		return nil, contracts.ErrBadRequestParametersMissing()
	}

	if request.Name == "" {
		return nil, contracts.ErrBadRequestInvalidParameter()
	}

	aM := model.NewAnomaly(request.Name)

	if request.FriendlyName != "" {
		aM.FriendlyName = request.FriendlyName
	}

	if request.SampleCount > 0 {
		aM.SampleCount = request.SampleCount
	}

	if request.IntervalInSec > 0 {
		aM.IntervalInSec = request.IntervalInSec
	}

	if request.MinSample > 0 {
		aM.MinSample = request.MinSample
	}

	if request.Tolerance > 0.0 {
		aM.Tolerance = request.Tolerance
	}

	err := detector.InitMetric(aM)
	if err != nil {
		return nil, contracts.ErrInternalServerError()
	}
	res := &contracts.CreateMetricsResponse{}
	sRes := &contracts.SingleMetricsResponse{}
	sRes.ResourceData = request.Anomaly
	res.ResponseData = append(res.ResponseData, sRes)
	return res, nil
}
