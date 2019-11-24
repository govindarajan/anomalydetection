package store

import (
	"strconv"
	"time"

	"github.com/govindarajan/anomalydetection/log"
	"github.com/govindarajan/anomalydetection/model"
)

// StoreHistoricMetric used to store the metric detaiks into a table
// TODO: Store in DB.
func StoreHistoricMetric(name string, m *model.Metric, anom *model.Anomaly) error {

	key := getHistoricMetricKey(name, m.T, anom.IntervalInSec)
	expireAt := getExpiry()

	value, err := m.Encode()
	if err != nil {
		return err
	}

	return GetStore().ExpirableSet(key, value, expireAt)

}

func GetHistoricMetric(name string, t time.Time, anom *model.Anomaly) (*model.Metric, error) {
	key := getHistoricMetricKey(name, t, anom.IntervalInSec)

	val, err := GetStore().ExpirableGet(key)
	if err != nil {
		log.Error("Error while getting historicmetric", name, err)
		return nil, err
	}

	if val == nil {
		return nil, nil
	}

	res, err := model.MetricDecode(val)
	if err != nil {
		log.Error("Error while casting Metric.")
		return res, err
	}

	return res, nil
}

func getHistoricMetricKey(name string, t time.Time, interval int64) string {
	window := (t.Unix() / interval) * interval
	return HIS_METRIC_PREFIX + name + "_" + strconv.Itoa(int(window))
}

func getExpiry() time.Time {
	// Lets expire these after 8 days.
	return time.Now().AddDate(0, 0, 8)
}
