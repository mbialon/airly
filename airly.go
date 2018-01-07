package airly

import (
	"fmt"
	"net/url"
)

// NearestSensorMeasurementParams holds nearest sensor's query parameters.
type NearestSensorMeasurementParams struct {
	// Latitude coordinate
	Latitude float64
	// Longitude coordinate
	Longitude float64
	// Max distance to nearest sensor
	MaxDistance int32
}

// NearestSensorMeasurements holds nearest sensor's current measurements.
type NearestSensorMeasurements struct {
	Address struct {
		Country      string `json:"country"`
		Locality     string `json:"locality"`
		Route        string `json:"route"`
		StreetNumber string `json:"streetNumber"`
	} `json:"address"`
	AirQualityIndex float64 `json:"airQualityIndex"`
	Distance        float64 `json:"distance"`
	ID              int64   `json:"id"`
	Location        struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	MeasurementTime string  `json:"measurementTime"`
	Name            string  `json:"name"`
	PM10            float64 `json:"pm10"`
	PM25            float64 `json:"pm25"`
	Vendor          string  `json:"vendor"`
	PollutionLevel  int     `json:"pollutionLevel"`
}

// NearestSensorMeasurements returns nearest sensor's current detailed measurements.
func (c *Client) NearestSensorMeasurements(params *NearestSensorMeasurementParams) (*NearestSensorMeasurements, error) {
	u, _ := url.Parse("v1/nearestSensor/measurements")
	u = c.BaseURL.ResolveReference(u)
	q := u.Query()
	q.Set("latitude", fmt.Sprintf("%f", params.Latitude))
	q.Set("longitude", fmt.Sprintf("%f", params.Longitude))
	q.Set("maxDistance", fmt.Sprintf("%d", params.MaxDistance))
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(u.String())
	if err != nil {
		return nil, err
	}
	v := &NearestSensorMeasurements{}
	if err := c.Do(req, v); err != nil {
		return nil, err
	}
	return v, nil
}
