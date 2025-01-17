// What is the maximal ride speed in rides.json?
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func maxRideSpeed(r io.Reader) (float64, error) {
	dec := json.NewDecoder(r)
	maxSpeed := -1.0
	for {
		var ride struct {
			StartTime string `json:"start"`
			EndTime   string `json:"end"`
			Distance  float64
		}
		err := dec.Decode(&ride)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		const timeFmt = "2006-01-02T15:04"
		// get the start time
		startTime, err := time.Parse(timeFmt, ride.StartTime)
		if err != nil {
			return 0, err
		}
		// get the end time
		endTime, err := time.Parse(timeFmt, ride.EndTime)
		if err != nil {
			return 0, err
		}
		dt := endTime.Sub(startTime) // find the Duration between the start and the end
		dtHour := float64(dt) / float64(time.Hour) // convert into Hour
		speed := ride.Distance / dtHour // calculate the speed
		if speed > maxSpeed {
			maxSpeed = speed
		}
	}

	return maxSpeed, nil
}

func main() {
	file, err := os.Open("rides.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed) // 40.5
}
