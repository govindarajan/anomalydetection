package store

import (
	"github.com/govindarajan/anomalydetection/log"
	"github.com/govindarajan/anomalydetection/model"
)

const ANO_PREFIX = "a"
const REC_METRIC_PREFIX = "rm"
const HIS_METRIC_PREFIX = "hm"

func UpsertAnomay(anom *model.Anomaly) error {
	inp, err := anom.Encode()
	if err != nil {
		return err
	}
	err = GetStore().Set(ANO_PREFIX+anom.Name, inp)
	if err != nil {
		log.Error("in insert anomaly", err)
	}
	return err
}

func GetAnomaly(name string) (*model.Anomaly, error) {
	val, err := GetStore().Get(ANO_PREFIX + name)
	if err != nil {
		log.Error("While reaching new metric type")
	}
	if val == nil {
		// Not found.
		return nil, nil
	}
	res, err := model.AnomalyDecode(val)
	if err != nil {
		log.Error("Error while casting Anomaly metric.", string(val))
		return res, err
	}
	return res, err
}

func GetRecentMetrics(name string) (*model.RecentMetrics, error) {
	val, err := GetStore().Get(REC_METRIC_PREFIX + name)
	if err != nil {
		log.Error("Error while reading Recentmetrics.", err)
		return nil, err
	}
	var res *model.RecentMetrics
	if val == nil {
		data := make(map[int64]float64, 20)
		res = &model.RecentMetrics{Name: name, Data: data}
		//res.Data = make(map[int64]float64)
		return res, nil
	}

	res, err = model.RecentMetricsDecode(val)
	if err != nil {
		log.Error("Error while casting Recentmetrics.", string(val))
		return res, err
	}
	return res, err

}

// This method takes care of
func AddRecentMetrics(name string, val *model.RecentMetrics) error {
	in, err := val.Encode()
	if err != nil {
		return err
	}
	err = GetStore().Set(REC_METRIC_PREFIX+name, in)
	if err != nil {
		log.Error("While writing Recentmetrics", name, err)
	}
	return err
}
