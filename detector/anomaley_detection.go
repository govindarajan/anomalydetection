package detector

import (
	"errors"
	"math"
	"time"

	"github.com/govindarajan/anomalydetection/log"
	"github.com/govindarajan/anomalydetection/model"
	"github.com/govindarajan/anomalydetection/store"
)

// InitMetric used to tell this service to initialize anomaly detection.
func InitMetric(m *model.Anomaly) error {
	// Store it
	if e := store.UpsertAnomay(m); e != nil {
		return e
	}
	// TODO: If it has sample data, push them
	return nil
}

// DetectAnomaly used to detect the whether given data point is anomaly or not.
// It returns its score as well.
// Anomaly is detected based on tolerance value
func DetectAnomaly(name string, tim time.Time, val float64) (bool, float64, error) {
	anomaly := false
	anomalyScore := 0.0
	// Get the metrics setting
	anom, err := store.GetAnomaly(name)
	if err != nil {
		log.Error(err)
		return anomaly, anomalyScore, err
	}
	if anom == nil {
		// Not found. Please initite the metrics.
		return anomaly, anomalyScore, errors.New("METRIC_NOT_FOUND")
	}

	// Get recent history metrics and calculate mean and SD
	recMetrics, err := store.GetRecentMetrics(name)
	if err != nil {
		log.Error(err)
		return anomaly, anomalyScore, err
	}
	// store it in the current val table
	recMetrics.InjectCurrMetric(tim, val, anom)
	store.AddRecentMetrics(name, recMetrics)

	m := &model.Metric{Val: val, T: tim}
	if recMetrics.Data == nil || len(recMetrics.Data) < int(anom.MinSample) {
		// Too less data point to detect
		return anomaly, anomalyScore, nil
	}
	divSigma(m, recMetrics.Data)

	// Check the current value whether it lies within the acceptable range.(calc anomaly)
	anomalyScore = m.Score
	if m.Score >= anom.Tolerance || m.Score <= (-1*anom.Tolerance) {
		anomaly = true
	}

	// store it in historical val table
	store.StoreHistoricMetric(name, m, anom)

	// If anomaly detected, Check the last week's data.
	// This step is to validate the anomlay.
	if anomaly {
		// See whether same anomaly found in last week as well.
		// If so, lets not call it as anomaly.
		lastWeekTime := getLastWeekTime(tim)
		hisMet, err := store.GetHistoricMetric(name, lastWeekTime, anom)
		if err != nil || hisMet == nil {
			// Error while fetching. Lets go with what we have.
			return anomaly, anomalyScore, nil
		}

		// Lets assume, we can go with 30% growth. More than that will be anomaly
		if m.Score <= 1.3*hisMet.Score {
			// Lets reset out anomaly values. Score remains same.
			anomaly = false
		}
	}

	// return the values
	return anomaly, anomalyScore, nil
}

func getLastWeekTime(t time.Time) time.Time {
	return t.AddDate(0, 0, -7)
}

func divSigma(m *model.Metric, vals map[int64]float64) {

	// Should We include m.Val also?
	avg := Average(vals)
	std := StdDev(vals, avg)
	m.Mean = avg
	m.SD = std

	if std == 0 {
		switch {
		case m.Val == avg:
			m.Score = 0
		case m.Val > avg:
			m.Score = 1
		case m.Val < avg:
			m.Score = -1
		}
		return
	}
	// 3Sigma
	m.Score = (m.Val - avg) / (3 * std)
}

// Average returns the mean value of float64 values.
func Average(vals map[int64]float64) float64 {

	var sum float64
	i := 0
	for _, v := range vals {
		sum += v
		i++
	}
	return sum / float64(i)
}

// StdDev returns the standard deviation of float64 values, with an input
// average.
func StdDev(vals map[int64]float64, avg float64) float64 {
	var sum float64
	var i int
	for _, v := range vals {
		dis := v - avg
		sum += dis * dis
		i++
	}
	return math.Sqrt(sum / float64(i))
}
