# go-prometheus-app

Simple app to create [Prometheus](https://prometheus.io) metrics in go.

## Paths

### /ping

Will increment the metric `ping_request_count` every time this endpoint is called.

### /alert

Will increment the metric `alert_count` every time this endpoint is called.