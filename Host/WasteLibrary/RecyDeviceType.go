package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RecyDeviceType
type RecyDeviceType struct {
	DeviceId      float64
	DeviceMain    RecyDeviceMainType
	DeviceBase    RecyDeviceBaseType
	DeviceGps     RecyDeviceGpsType
	DeviceTherm   RecyDeviceThermType
	DeviceVersion RecyDeviceVersionType
	DeviceAlarm   RecyDeviceAlarmType
	DeviceStatu   RecyDeviceStatuType
	DeviceDetail  RecyDeviceDetailType
	DeviceNote    RecyDeviceNoteType
}

//New
func (res *RecyDeviceType) New() {
	res.DeviceId = 0
	res.DeviceMain.New()
	res.DeviceBase.New()
	res.DeviceGps.New()
	res.DeviceTherm.New()
	res.DeviceVersion.New()
	res.DeviceAlarm.New()
	res.DeviceStatu.New()
	res.DeviceDetail.New()
	res.DeviceNote.New()
}

//GetByRedis
func (res *RecyDeviceType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType

	res.DeviceMain.DeviceId = res.DeviceId
	resultVal = res.DeviceMain.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceBase.DeviceId = res.DeviceId
	resultVal = res.DeviceBase.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceGps.DeviceId = res.DeviceId
	resultVal = res.DeviceGps.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceTherm.DeviceId = res.DeviceId
	resultVal = res.DeviceTherm.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceVersion.DeviceId = res.DeviceId
	resultVal = res.DeviceVersion.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceAlarm.DeviceId = res.DeviceId
	resultVal = res.DeviceAlarm.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceStatu.DeviceId = res.DeviceId
	resultVal = res.DeviceStatu.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceDetail.DeviceId = res.DeviceId
	resultVal = res.DeviceDetail.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceNote.DeviceId = res.DeviceId
	resultVal = res.DeviceNote.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisBySerial
func (res *RecyDeviceType) GetByRedisBySerial(serial string) ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_SERIAL_RECY_DEVICE, serial)
	if resultVal.Result == RESULT_OK {
		var deviceId string = resultVal.Retval.(string)
		res.DeviceId = StringIdToFloat64(deviceId)
		resultVal = res.GetByRedis("0")
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//ToId String
func (res *RecyDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *RecyDeviceType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
