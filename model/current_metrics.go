package model

import (
	"encoding/json"
	"sync"
	"time"
)

type Metric struct {
	Val   float64   `json:"val"`
	T     time.Time `json:"time"`
	SD    float64   `json:"sd"`
	Mean  float64   `json:"mean"`
	Score float64   `json:"score"`
}

func (m *Metric) Encode() ([]byte, error) {
	return json.Marshal(m)
}

func MetricDecode(b []byte) (m *Metric, e error) {
	e = json.Unmarshal(b, m)
	return
}

type RecentMetrics struct {
	Name string            `json:"name"`
	Data map[int64]float64 `json:"data"`
	lock sync.Mutex        `json:"-"`
}

func (m *RecentMetrics) Encode() ([]byte, error) {
	return json.Marshal(m)
}

func RecentMetricsDecode(b []byte) (m *RecentMetrics, e error) {
	e = json.Unmarshal(b, &m)
	return
}

// InjectCurrMetric will insert the metrics to current list and delete the oldest one
func (m *RecentMetrics) InjectCurrMetric(tim time.Time, val float64, aM *Anomaly) {
	// Calculate the metric time window
	window := (tim.Unix() / aM.IntervalInSec) * aM.IntervalInSec
	oldWindow := window - (aM.SampleCount * window)
	m.lock.Lock()
	defer m.lock.Unlock()
	// Add to the list
	// Delete the older metrics based on the interval and samples.
	m.Data[window] = val
	delete(m.Data, oldWindow)
}
