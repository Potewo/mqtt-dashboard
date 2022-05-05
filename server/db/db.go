package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Temperature struct {
	gorm.Model
	V float32
	SensorId uint
}

type Humidity struct {
	gorm.Model
	V float32
	SensorId uint
}

type AirPressure struct {
	gorm.Model
	V float32
	SensorId uint
}

var Atemperature Temperature
var DB *gorm.DB
var err error

func init() {
	connect(30)
	DB.AutoMigrate(&Temperature{}, &Humidity{}, &AirPressure{})
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
