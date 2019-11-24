curl  -XPOST -H 'Content-Type: application/json' 'http://localhost:8080/metrics/testmetric' -d '{"Name":"testmetric","friendly_name":"SomeName","sample_count":20,"interval_in_sec":60,"min_sample":15,"tolerance":1}'

curl  -XPOST -H 'Content-Type: application/json' 'http://localhost:8080/metrics/testmetric' --data-raw '{"Name":"testmetric"}'

curl  -XPOST -H 'Content-Type: application/json' 'http://localhost:8080/metrics/testmetric/datapoint' -d '{"time":"2019-10-12T20:00:00+05:30", "value":23}'

