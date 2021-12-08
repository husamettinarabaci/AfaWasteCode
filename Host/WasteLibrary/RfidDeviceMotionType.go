package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RfidDeviceMotionType
type RfidDeviceMotionType struct {
	DeviceId float64
	Motions  []GpsMotionType
}

//New
func (res *RfidDeviceMotionType) New() {
	res.DeviceId = 0
	res.Motions = []GpsMotionType{}
}

//GetByRedis
func (res *RfidDeviceMotionType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_MOTION, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceMotionType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_MOTION, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *RfidDeviceMotionType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceMotionType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceMotionType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RfidDeviceMotionType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceMotionType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
