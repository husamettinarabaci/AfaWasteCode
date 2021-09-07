package WasteLibrary

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//Debug Mode
var Debug bool = os.Getenv("DEBUG") == "1"

//Container
var Container string = os.Getenv("CONTAINER_TYPE")

//GetTime
func GetTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}

//ToCustomerId String
func Float64IdToString(id float64) string {
	return fmt.Sprintf("%.0f", id)
}

//StringIdToFloat64
func StringIdToFloat64(id string) float64 {
	floatId, _ := strconv.Atoi(id)
	return float64(floatId)
}

//StringToFloat64LatLong
func StringToFloat64(latLong string) float64 {
	var fVal float64 = 0
	if s, err := strconv.ParseFloat(latLong, 32); err == nil {
		fVal = s
	}
	return fVal
}
