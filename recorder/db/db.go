package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Temperature struct {
	gorm.Model
	V        float32
	SensorId uint
}

type Humidity struct {
	gorm.Model
	V        float32
	SensorId uint
}

type AirPressure struct {
	gorm.Model
	V        float32
	SensorId uint
}

type Co2 struct {
	gorm.Model
	V        float32
	SensorId uint
}

type DBStruct interface {
	Temperature | Humidity | AirPressure | Co2
}

var Atemperature Temperature
var DB *gorm.DB
var err error

func init() {
	connect(30)
	DB.AutoMigrate(&Temperature{}, &Humidity{}, &AirPressure{}, &Co2{})
	DB.Create(&Temperature{V: 25, SensorId: 1})
	DB.First(&Atemperature, 1)
}

func connect(count uint) {
	dsn := "root:password@tcp(db)/mqtt?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		time.Sleep(time.Second)
		count--
		if count == 0 {
			panic(err)
		} else {
			connect(count)
		}
	}
}

func Get[T DBStruct](timeRange *TimeRange, sensorIds []uint) []T {
	db := DB
	if timeRange != nil {
		db = DB.Where("created_at BETWEEN ? AND ?", timeRange.Start, timeRange.End)
	}
	if sensorIds != nil {
		db = db.Where("sensor_id IN ?", sensorIds)
	}
	values := []T{}
	db.Find(&values)
	return values
}

func Set[T DBStruct](value T) error {
	ctx := DB.Create(&value)
	return ctx.Error
}

type TimeRange struct {
	Start time.Time
	End   time.Time
}
