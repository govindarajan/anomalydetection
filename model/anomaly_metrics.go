package model

import "encoding/json"

// Anomaly - Metric for which we need to detect anomaly
type Anomaly struct {
	Name          string  `json:"name" required:"true"`
	FriendlyName  string  `json:"friendly_name"`
	SampleCount   int64   `json:"sample_count"`
	IntervalInSec int64   `json:"interval_in_sec"` // TODO: Change it to duration.
	MinSample     int64   `json:"min_sample"`
	Tolerance     float64 `json:"tolerance"`
	// TODO: Alerting method
}

// NewAnomaly is used get the new instance to metric for which we need to detect.
func NewAnomaly(name string) *Anomaly {
	return &Anomaly{
		Name:          name,
		Tolerance:     1.0,
		MinSample:     20,
		SampleCount:   30,
		IntervalInSec: 60,
	}
}

// Encode - to encode.
func (anom *Anomaly) Encode() ([]byte, error) {
	return json.Marshal(anom)
}

// AnomalyDecode - to get Anomaly from byte array
func AnomalyDecode(b []byte) (*Anomaly, error) {
	var anom Anomaly
	e := json.Unmarshal(b, &anom)
	return &anom, e
}
