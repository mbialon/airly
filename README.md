# airly [![GoDoc](https://godoc.org/github.com/mbialon/airly?status.svg)](http://godoc.org/github.com/mbialon/airly) [![Go Report Card](https://goreportcard.com/badge/github.com/mbialon/airly)](https://goreportcard.com/report/github.com/mbialon/airly)

[Airly](https://airly.eu) API client for Go.

## Usage

```go
package main

import (
	"os"

	"github.com/mbialon/airly"
)

func main() {
	c := airly.NewClient(airly.WithAPIKey(os.Getenv("AIRLY_APIKEY")))
	v, err := c.NearestSensorMeasurements(&airly.NearestSensorMeasurementParams{
		Latitude:    50.09394,
		Longitude:   18.99622,
		MaxDistance: 1500,
	})
	if err != nil {
		panic(err)
	}
	q := v.AirQualityIndex
	switch {
	case q <= 25:
		println("Very low")
	case q <= 50:
		println("Low")
	case q <= 75:
		println("Medium")
	case q <= 100:
		println("High")
	case q > 100:
		println("Very high")
	}
}
```
