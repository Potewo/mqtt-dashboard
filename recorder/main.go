package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/Potewo/mqtt-dashboard/recorder/db"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var msgCh = make(chan mqtt.Message)
func brokerLoadHandler(client mqtt.Client, msg mqtt.Message) {
	msgCh <- msg
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("mqtt://broker:1883")
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Mqtt error: %s", token.Error())
	}

	if subscribeToken := c.Subscribe("sensor0/#", 0, brokerLoadHandler); subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}
	if subscribeToken := c.Subscribe("sensor1/#", 0, brokerLoadHandler); subscribeToken.Wait() && subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	for {
		select {
		case m := <-msgCh:
			switch m.Topic() {
			case "sensor0/temp":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to parse float")
				}
				value := db.Temperature{V: float32(v), SensorId: 0}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			case "sensor0/humid":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to parse float")
				}
				value := db.Humidity{V: float32(v), SensorId: 0}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			case "sensor0/co2":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
				value := db.Co2{V: float32(v), SensorId: 0}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			case "sensor1/temp":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
				value := db.Temperature{V: float32(v), SensorId: 1}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			case "sensor1/humid":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
				value := db.Humidity{V: float32(v), SensorId: 1}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			case "sensor1/air_pressure":
				fmt.Println("received")
				str := string(m.Payload())
				v, err := strconv.ParseFloat(str, 32)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
				value := db.AirPressure{V: float32(v), SensorId: 1}
				err = db.Set(value)
				if err != nil {
					fmt.Printf("failed to set value to db")
				}
			default:
				fmt.Println(m.Topic())
			}
		fmt.Printf("topic: %v, payload: %v\n", m.Topic(), string(m.Payload()))
		case <-signalCh:
			fmt.Printf("Interrupt detected.\n")
			c.Disconnect(1000)
			return
		}
	}
}
