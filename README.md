# go-prometheus-app

Simple app to create [Prometheus](https://prometheus.io) metrics in go.

## Paths

### /ping

Will increment the metric `ping_request_count` every time this endpoint is called.

### /alert

Will increment the metric `alert_count` every time this endpoint is called.

### /resetalert

Will reset the metric `alert_count` to `0`.

### /version

Will return the version of this app. Should be equal to its docker image.


### /print

Will print the content of body. Used to be a webhook for [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/).

### /histogram

Will randomly sleep for 0 to 5 seconds and put it in a histogram with a label `id` between 0 and 4.
