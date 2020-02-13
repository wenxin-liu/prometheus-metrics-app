# prometheus-metrics-app
Instruments mock metrics for Prometheus

Instrumented three mock metrics (boring_count, special_count, response_count). 
* boring_count increments by 1 every time any path apart from /special and /metrics is requested. 
* special_count increments by 1 every time /special is requested. 
* request_count increments by 1 every time any path apart from /metrics is requested.

Metrics are exposed via the /metrics endpoint.
