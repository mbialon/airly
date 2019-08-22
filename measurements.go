package airly

import (
	"fmt"
	"net/url"
)

// Measurements holds measurements data.
type Measurements struct {
	Current *AveragedValues `json:"current"`
}

type AveragedValues struct {
	Values  []*Value `json:"values"`
	Indexes []*Index `json:"indexes"`
}

type Value struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Index struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	Level       string  `json:"level"`
	Description string  `json:"description"`
	Advice      string  `json:"advice"`
	Color       string  `json:"color"`
}

// NearestMeasurements returns measurements for an installation closest to a given location.
func (c *Client) NearestMeasurements(lat, lon float64, dist float32) (*Measurements, error) {
	u, _ := url.Parse("v2/measurements/nearest")
	u = c.BaseURL.ResolveReference(u)
	q := u.Query()
	q.Set("lat", fmt.Sprintf("%f", lat))
	q.Set("lng", fmt.Sprintf("%f", lon))
	q.Set("maxDistanceKM", fmt.Sprintf("%f", dist))
	u.RawQuery = q.Encode()
	req, err := c.NewRequest(u.String())
	if err != nil {
		return nil, err
	}
	v := &Measurements{}
	if err := c.Do(req, v); err != nil {
		return nil, err
	}
	return v, nil
}
