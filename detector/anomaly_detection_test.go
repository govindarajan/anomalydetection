package detector_test

import (
	"testing"
	"time"

	ad "github.com/govindarajan/anomalydetection/detector"
	"github.com/govindarajan/anomalydetection/model"
	store "github.com/govindarajan/anomalydetection/store"
)

func TestBasic(t *testing.T) {
	db, err := store.NewSQLite("/tmp/adix1.db")
	if err != nil {
		t.Error(err)
	}
	store.InitStore(db)
	name := "testing1"
	e := ad.InitMetric(model.NewAnomaly(name))
	if e != nil {
		t.Error(e)
	}
	tim := time.Now()
	ot := tim.Add(time.Minute * -50)

	input := []float64{800, 975, 936, 940, 925, 1015, 920, 881, 939, 1078, 997, 867, 903, 1500}
	for i, in := range input {
		now := ot.Add(time.Minute * time.Duration(i))
		isAnom, _, _ := ad.DetectAnomaly(name, now, in)
		if i == 12 && isAnom {
			t.Error("Anomaly detection failed")
		}
	}

}

func BenchmarkTest(b *testing.B) {
	db, _ := store.NewSQLite("/tmp/adix1.db")
	store.InitStore(db)
	name := "testing2"
	ad.InitMetric(model.NewAnomaly(name))
	tim := time.Now()
	ot := tim.Add(time.Minute * -50)
	in := 100.0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		now := ot.Add(time.Minute * time.Duration(i))
		ad.DetectAnomaly(name, now, in)
	}
}

func TestCPU(t *testing.T) {
	db, err := store.NewSQLite("/tmp/adix2.db")
	if err != nil {
		t.Error("Error while initing DB", err)
	}
	store.InitStore(db)

	input := []float64{27, 26, 27, 28, 27, 27, 28, 38, 32, 29, 29, 28, 28, 29, 29, 28, 28, 26, 28, 26, 28, 28, 27, 29, 27, 27, 28, 26, 27, 27, 27, 27, 27, 29, 27, 27, 27, 27, 27, 27, 28, 28, 26, 28, 28, 28, 27, 28, 29, 28, 28, 28, 29, 28, 28, 28, 29, 28, 29, 29, 29, 28, 31, 30, 28, 29, 24, 26, 25, 22, 21, 20, 22, 21, 20, 20, 20, 20, 20, 22, 23, 21, 22, 22, 22, 24, 22, 23, 24, 23, 24, 24, 25, 25, 24, 24, 24, 25, 25, 25, 25, 25, 24, 26, 24, 24, 24, 26, 24, 26, 24, 25, 26, 26, 26, 26, 24, 25, 26, 25, 26, 26, 25, 25, 24, 26, 25, 24, 26, 25, 24, 25, 25, 25, 25, 26, 26, 25, 25, 25, 26, 24, 24, 27, 26, 25, 24, 25, 26, 24, 24, 26, 24, 24, 24, 25, 26, 24, 26, 25, 24, 26, 24, 24, 24, 24, 23, 24, 24, 25, 24, 25, 25, 25, 25, 25, 25, 25, 25, 25, 37.345, 36.721, 36.955, 38.165, 37.433, 37.433, 37.698, 37.616, 38.386, 37.668, 37.989, 37.543, 37.879, 39.634, 39.086, 38.317, 38.957, 38.236, 40.032, 37.696, 38.742, 37.973, 38.457, 38.967, 38.04, 38.13, 38.197, 38.185, 38.676, 38.172, 37.869, 37.771, 37.736, 39.38, 38.244, 38.382, 38.47, 38.321, 38.752, 38.563, 38.026, 38.715, 37.814, 39.144, 38.788, 38.839, 39.076, 39.306, 39.654, 39.098, 39.625, 38.662, 39.697, 40.092, 39.453, 39.55, 39.387, 39.455, 40.301, 40.393, 40.303, 40.024, 40.802, 42.1, 40.763, 40.007, 34.452, 34.415, 35.359, 29.929, 28.908, 28.705, 28.775, 29.706, 29.435, 29.018, 28.969, 29.35, 29.93, 29.159, 30.132, 29.962, 30.282, 31.001, 30.852, 31.181, 30.371, 32.672, 33.422, 32.572, 32.521, 33.125, 33.257, 34.305, 33.994, 34.739, 34.391, 34.565, 35.162, 34.459, 34.445, 34.467, 34.446, 35.003, 34.593, 34.509, 34.572, 34.932, 35.118, 35.061, 34.994, 34.256, 34.624, 35.736, 35.416, 35.112, 35.021, 36.01, 36.966, 35.463, 35.774, 35.222, 35.292, 35.376, 34.891, 35.176, 34.467, 35.458, 35.71, 34.883, 35.248, 34.538, 34.206, 36.32, 34.899, 35.372, 35.276, 35.149, 35.676, 35.28, 34.303, 34.483, 34.803, 35.468, 35.428, 35.067, 34.852, 34.286, 35.272, 34.205, 34.444, 34.887, 34.411, 34.864, 34.602, 34.532, 33.323, 34.517, 35.027, 34.277, 34.538, 33.722, 34.148, 34.176, 34.087, 34.158, 33.31, 33.908, 34.588, 34.119, 33.508, 33.709, 33.286, 33.907, 33.535, 34.181, 35.145, 34.897, 34.445, 35.235}

	name := "testingcpu1"
	e := ad.InitMetric(model.NewAnomaly(name))
	if e != nil {
		t.Error(e)
	}
	tim := time.Now()
	ot := tim.Add(time.Minute * -400)
	for i, in := range input {
		now := ot.Add(time.Minute * time.Duration(i))
		ad.DetectAnomaly(name, now, in)
		//fmt.Println(isAnom, score, now)

	}
	//t.Error("Print log")
}
