package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/peterbourgon/ff/v3"
	"gofer.dev/airly/v2"
)

func main() {
	fs := flag.NewFlagSet("airly-status", flag.ExitOnError)
	var (
		fdownload = fs.Bool("d", false, "Download measurements")
		fstatus   = fs.Bool("s", false, "Print status")
		apikey    = fs.String("apikey", "", "Airly API key")
		lat       = fs.Float64("lat", 0, "Latitude")
		lon       = fs.Float64("lon", 0, "Longitude")
		dist      = fs.Float64("dist", 1.5, "Max distance in kilometers")
	)
	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}
	switch {
	case *fdownload:
		err = download(*apikey, *lat, *lon, *dist)
	case *fstatus:
		err = status()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}
}

const defaultAirlyFile = "/tmp/airly.json"

func download(apikey string, lat, lon, dist float64) error {
	c := airly.NewClient(apikey)
	m, err := c.NearestMeasurements(lat, lon, float32(dist))
	if err != nil {
		return err
	}
	f, err := os.OpenFile(defaultAirlyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(m)
}

func decode() (*airly.Measurements, time.Time, error) {
	f, err := os.Open(defaultAirlyFile)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return nil, time.Time{}, err
	}
	m := new(airly.Measurements)
	if err := json.NewDecoder(f).Decode(m); err != nil {
		return nil, time.Time{}, err
	}
	return m, stat.ModTime(), nil
}

func icon(m *airly.Measurements) string {
	if m == nil {
		return "⚠️"
	}
	var caqi *airly.Index
	for _, idx := range m.Current.Indexes {
		if idx.Name == "AIRLY_CAQI" {
			caqi = idx
		}
	}
	var icon string
	switch caqi.Level {
	case "VERY_LOW":
		icon = "💙"
	case "LOW":
		icon = "💚"
	case "MEDIUM":
		icon = "💛"
	case "HIGH":
		icon = "🧡"
	case "VERY_HIGH":
		icon = "❤️"
	case "EXTREME":
		icon = "☢️"
	case "AIRMAGEDDON":
		icon = "☣️"
	default:
		icon = "⚠️"
	}
	return icon
}

func indexValue(m *airly.Measurements, name string) float64 {
	if m == nil {
		return 0
	}
	var caqi *airly.Index
	for _, idx := range m.Current.Indexes {
		if idx.Name == name {
			caqi = idx
			break
		}
	}
	return caqi.Value
}

func value(m *airly.Measurements, name string) float64 {
	if m == nil {
		return 0
	}
	var val *airly.Value
	for _, v := range m.Current.Values {
		if v.Name == name {
			val = v
			break
		}
	}
	return val.Value
}

func timeliness(mtime time.Time) int64 {
	if mtime.IsZero() {
		return 0
	}
	return time.Now().Unix() - mtime.Unix()
}

func status() error {
	meas, mtime, _ := decode()
	fmt.Printf("%s %0.1f🌡%0.1f°C 💦 %0.f%% ️+%ds\n",
		icon(meas),
		indexValue(meas, "AIRLY_CAQI"),
		value(meas, "TEMPERATURE"),
		value(meas, "HUMIDITY"),
		timeliness(mtime))
	return nil
}
