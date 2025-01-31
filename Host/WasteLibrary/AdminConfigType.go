package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//AdminConfigType
type AdminConfigType struct {
	CustomerId      float64
	DeviceBaseWork  string
	WorkStartHour   int
	WorkStartMinute int
	WorkAddMinute   int
	CenterLatitude  float64
	CenterLongitude float64
	SWLatitude      float64
	SWLongitude     float64
	NELatitude      float64
	NELongitude     float64
	DefaultZoom     int
	MinZoom         int
	MaxZoom         int
	ParkLocations   []LocationType
	WebUIPrivate    string
	Active          string
	CreateTime      string
}

//New
func (res *AdminConfigType) New() {
	res.CustomerId = 1
	res.WorkStartHour = 06
	res.WorkStartMinute = 0
	res.WorkAddMinute = 510
	res.SWLatitude = 0
	res.SWLongitude = 0
	res.NELatitude = 0
	res.NELongitude = 0
	res.CenterLatitude = 0
	res.CenterLongitude = 0
	res.DefaultZoom = 11
	res.MinZoom = 8
	res.MaxZoom = 18
	res.DeviceBaseWork = STATU_PASSIVE
	res.ParkLocations = []LocationType{}
	res.WebUIPrivate = STATU_PASSIVE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *AdminConfigType) GetByRedis() ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_CUSTOMER_ADMINCONFIG, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *AdminConfigType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_ADMINCONFIG, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *AdminConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *AdminConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *AdminConfigType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *AdminConfigType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *AdminConfigType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
