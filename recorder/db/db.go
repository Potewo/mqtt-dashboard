package db

import (
	"context"
	"strconv"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type SensorType int

const (
	Temperature = iota
	Humidity
	AirPressure
	Co2
)

func (st SensorType) String() string {
	switch st {
	case Temperature:
		return "Temperature"
	case Humidity:
		return "Humidity"
	case AirPressure:
		return "AirPressure"
	case Co2:
		return "Co2"
	default:
		return "Unknown"
	}
}

type DBStruct struct {
	SensorType SensorType
	V          float32
	SensorId   uint
}

// TODO set names
var bucket = "climate"
var org = "my_org"
var token = "myadmintoken"
var url = "http://db:8086"

var client influxdb2.Client

var writeAPI api.WriteAPIBlocking
var err error

func init() {
	client = influxdb2.NewClient(url, token)
	writeAPI = client.WriteAPIBlocking(org, bucket)
}

func Set(value DBStruct) error {
	p := influxdb2.NewPoint(
		value.SensorType.String(),
		map[string]string{"sensor-id": strconv.Itoa(int(value.SensorId))},
		map[string]interface{}{"value": value.V}, time.Now())
	return writeAPI.WritePoint(context.Background(), p)
}

func Close() {
	client.Close()
}
