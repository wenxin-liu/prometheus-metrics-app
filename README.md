# prometheus-metrics-app
Instruments mock metrics for Prometheus

Instrumented four mock metrics (boring_count, special_count, response_count, incremental_count). 
* boring_count increments by 1 every time any path apart from /special and /metrics is requested. 
* special_count increments by 1 every time /special is requested. 
* request_count increments by 1 every time any path apart from /metrics is requested.
* incremental_count increases by 1 every 2 seconds on its own.

Metrics are exposed via the /metrics endpoint.
