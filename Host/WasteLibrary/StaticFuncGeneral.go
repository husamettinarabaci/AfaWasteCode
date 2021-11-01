package WasteLibrary

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

//Debug Mode
var Debug bool = os.Getenv("DEBUG") == "1"
var AllowCors bool = true
var Version string = "1"

//Container
var Container string = os.Getenv("CONTAINER_TYPE")

//GetTime
func GetTime() string {
	currentTime := time.Now()
	return currentTime.Format(time.RFC3339)
}

//GetTimePlus
func GetTimePlus(plus time.Duration) string {
	plusTime := time.Now().Add(plus)
	return plusTime.Format(time.RFC3339)
}

//AddTimePlus
func AddTimeToBase(base time.Time, plus time.Duration) time.Time {
	plusTime := base.Add(plus)
	return plusTime
}

//StringToTime
func StringToTime(timeStr string) time.Time {
	currentTime, _ := time.Parse(time.RFC3339, timeStr)
	return currentTime
}

//TimeToString
func TimeToString(timeVal time.Time) string {
	return timeVal.Format(time.RFC3339)
}

//Float64IdToString String
func Float64IdToString(id float64) string {
	return fmt.Sprintf("%.0f", id)
}

//Float64ToString String
func Float64ToString(floatVal float64) string {
	return fmt.Sprint(floatVal)
}

//StringIdToFloat64
func StringIdToFloat64(id string) float64 {
	floatId, _ := strconv.Atoi(id)
	return float64(floatId)
}

//StringToFloat64
func StringToFloat64(sVal string) float64 {
	var fVal float64 = 0
	if s, err := strconv.ParseFloat(sVal, 32); err == nil {
		fVal = s
	}
	return fVal
}

//GetMD5Hash
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
