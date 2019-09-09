# airly [![GoDoc](https://godoc.org/gofer.dev/airly?status.svg)](http://godoc.org/gofer.dev/airly) [![Go Report Card](https://goreportcard.com/badge/gofer.dev/airly)](https://goreportcard.com/report/gofer.dev/airly)

[Airly](https://airly.eu) API client for Go.

## Usage

```shell
$ go get gofer.dev/airly/cmd/airly-status
```

### i3

`airly-status -d` downloads measurements for the given location to `/tmp/airly.json`.

```shell
$ airly-status -d -apikey ${APIKEY} -lat ${LATITUDE} -lon ${LONGITUDE} -dist 0.5
```

`airly-status -s` prints status in the terminal or i3status.

```shell
$ airly-status -s
💚 25.8🌡14.8°C 💦 97% ️+55s
```
