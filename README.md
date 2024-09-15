# SimpleFIN-bridge exporter

## Overview
SimpleFIN-bridge exporter is a simple application that exports account balance information as a prometheus scrapable endpoint.

SimpleFIN Bridge lets you securely share your financial transaction data with apps; this application behaves as a third-party client to the bridge. For more information, check [SimpleFIN's awesome documentation](https://beta-bridge.simplefin.org/).

Currently, this application only exposes balance, and no transaction information is ever used.

> [!CAUTION]
> This application exposes financial information about your accounts on an insecure network endpoint (`localhost` by default). The user is solely responsible for enforcing proper security measures and policies to prevent this endpoint's information from leaking to anything other than the intended consumer.

## Getting started
### Prequisites

- [SimpleFIN bridge account](https://beta-bridge.simplefin.org/). You will need to be able to issue a `Setup token` or an `Access Url`. Refer to their developer documentation and make sure you understand the security concerns for storing either.

You can launch the application using their ready-to-use demo token:
```sh
~/simplefin-exporter $ go run ./cmd -setupToken aHR0cHM6Ly9iZXRhLWJyaWRnZS5zaW1wbGVmaW4ub3JnL3NpbXBsZWZpbi9jbGFpbS9ERU1P
12:44PM INF cmd/main.go:107 > update interval: 1h0m0s
12:44PM INF cmd/main.go:109 > polling account data
12:44PM INF cmd/main.go:118 > done, took 258.14673ms
```
Validate the application is running by `curl`ing on localhost:

```sh
~/simplefin-exporter $ curl localhost:8000/metrics
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP simplefin_balance Account balance
# TYPE simplefin_balance gauge
simplefin_balance{account_name="SimpleFIN Checking",domain="beta-bridge.simplefin.org"} 25584.44921875
simplefin_balance{account_name="SimpleFIN Savings",domain="beta-bridge.simplefin.org"} 115104.4921875
```
