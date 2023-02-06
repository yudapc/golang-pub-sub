# Golang Pub/Sub with Redis

## How to Install

* Copy `env.sh.sample` to `env.sh`

* Setup redis password and topic name in file `env.sh`

* Apply env with `source env.sh`

* Install dependencies `go mod tidy`


---

## How to run

Run subscriber: `go run subscriber/subscriber.go`

Run publsiher: `go run publisher/publisher.go`