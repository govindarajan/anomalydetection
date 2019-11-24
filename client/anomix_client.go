package anomix

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/govindarajan/anomalydetection/anomix/pkg/contracts"
	"github.com/govindarajan/anomalydetection/model"
	"github.com/valyala/fasthttp"
)

type client struct {
	url string
}

var c *client
var ERR_METRIC_NOT_FOUND = errors.New("METRIC_NOT_FOUND")

// Init used to initialize anomic client
func Init(u string) error {
	// TODO: Validate url
	c = &client{url: u}
	return nil
}

// CreateMetric used to initialize metric in the service which will start detecting anomalies
// for data points. Without this, SendMetric will not work.
func CreateMetric(name string, sampleCount int64, intervalInSec int64, minSample int64, tolerance float64) error {
	am := model.NewAnomaly(name)
	am.SampleCount = sampleCount
	am.IntervalInSec = intervalInSec
	am.MinSample = minSample
	am.Tolerance = tolerance

	data, err := am.Encode()
	if err != nil {
		return err
	}
	code, body, err := doPost(data, c.url+"metrics/"+name)
	if err != nil {
		return err
	}
	if code != fasthttp.StatusOK {
		return errors.New(strconv.Itoa(code) + string(body))
	}
	return nil
}

// SendMetric is used to send the datapoints for the given metric. It returns whether the given point
// is anomaly or not.
func SendMetric(name string, t time.Time, value float64) (isAnomaly bool, score float64, err error) {
	req := &contracts.CreateDataPointRequest{Name: &name, Time: &t, Value: &value}

	data, err := json.Marshal(req)
	if err != nil {
		return isAnomaly, score, err
	}

	code, body, err := doPost(data, c.url+"metrics/"+name+"/datapoint")
	if err != nil {
		return isAnomaly, score, err
	}

	var res contracts.CreateDataPointResponse

	// Decode the  body
	err = json.Unmarshal(body, &res)
	if err != nil {
		return isAnomaly, score, err
	}
	if res.ResponseData == nil || len(res.ResponseData) <= 0 {
		return isAnomaly, score, errors.New("Unknown error. Empty response received")
	}

	// If 404 Not Found
	if code == http.StatusNotFound {
		// check if code is 1011 then send return ERR_NOT_FOUND
		if res.ResponseData[0].ErrorData.Code == 1011 {
			return isAnomaly, score, ERR_METRIC_NOT_FOUND
		}
		return isAnomaly, score, err
	}

	// IF 200 OK
	if code != fasthttp.StatusOK {
		return isAnomaly, score, errors.New(strconv.Itoa(code) + " " + string(body))
	}

	// We are here means, all good.
	isAnomaly = res.ResponseData[0].IsAnomaly
	score = res.ResponseData[0].Score

	return isAnomaly, score, err
}

var defaultClient fasthttp.Client

const contentType = "application/json"

func doPost(dst []byte, url string) (statusCode int, body []byte, err error) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()
	// TODO: Fix this
	req.Header.SetMethod("POST")
	req.Header.SetContentType(contentType)
	req.SetRequestURI(url)
	req.SetBody(dst)

	err = fasthttp.Do(req, res)
	statusCode = res.Header.StatusCode()
	body = res.Body()

	return statusCode, body, err
}
