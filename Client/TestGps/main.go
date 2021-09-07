package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

type NMEA struct {
	fixTimestamp       string
	latitude           string
	latitudeDirection  string
	longitude          string
	longitudeDirection string
	fixQuality         string
	satellites         string
	horizontalDilution string
	antennaAltitude    string
	antennaHeight      string
	updateAge          string
}

//GNGGA,095507.00,3953.65517,N,03248.30082,E,2,07,1.47,935.3,M,36.0,M,,0000
func ParseNMEALine(line string) (NMEA, error) {
	tokens := strings.Split(line, ",")
	if tokens[0] == "$GNGGA" {
		return NMEA{
			fixTimestamp:       tokens[1],
			latitude:           tokens[2],
			latitudeDirection:  tokens[3],
			longitude:          tokens[4],
			longitudeDirection: tokens[5],
			fixQuality:         tokens[6],
			satellites:         tokens[7],
		}, nil
	}
	return NMEA{}, errors.New("unsupported nmea string")
}

func ParseDegrees(value string, direction string) (string, error) {
	if value == "" || direction == "" {
		return "", errors.New("the location and / or direction value does not exist")
	}
	lat, _ := strconv.ParseFloat(value, 64)
	degrees := math.Floor(lat / 100)
	minutes := ((lat / 100) - math.Floor(lat/100)) * 100 / 60
	decimal := degrees + minutes
	if direction == "W" || direction == "S" {
		decimal *= -1
	}
	return fmt.Sprintf("%.6f", decimal), nil
}

func (nmea NMEA) GetLatitude() (string, error) {
	return ParseDegrees(nmea.latitude, nmea.latitudeDirection)
}

func (nmea NMEA) GetLongitude() (string, error) {
	return ParseDegrees(nmea.longitude, nmea.longitudeDirection)
}

func main() {
	fmt.Println("Starting the application...")
	options := serial.OpenOptions{
		PortName:        "/dev/ttyAMA0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	serialPort, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	defer serialPort.Close()
	reader := bufio.NewReader(serialPort)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		gps, err := ParseNMEALine(scanner.Text())
		if err == nil {
			if gps.fixQuality == "1" || gps.fixQuality == "2" {
				latitude, _ := gps.GetLatitude()
				longitude, _ := gps.GetLongitude()
				fmt.Println(latitude + "," + longitude)
			} else {
				fmt.Println("no gps fix available")
			}
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println(err)
		}
	}
}
