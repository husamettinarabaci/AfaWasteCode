package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//UltDeviceType
type UltDeviceType struct {
	DeviceId      float64
	DeviceMain    UltDeviceMainType
	DeviceBase    UltDeviceBaseType
	DeviceStatu   UltDeviceStatuType
	DeviceBattery UltDeviceBatteryType
	DeviceGps     UltDeviceGpsType
	DeviceAlarm   UltDeviceAlarmType
	DeviceTherm   UltDeviceThermType
	DeviceVersion UltDeviceVersionType
	DeviceSens    UltDeviceSensType
	DeviceNote    UltDeviceNoteType
}

//New
func (res *UltDeviceType) New() {
	res.DeviceId = 0
	res.DeviceMain.New()
	res.DeviceBase.New()
	res.DeviceStatu.New()
	res.DeviceBattery.New()
	res.DeviceGps.New()
	res.DeviceAlarm.New()
	res.DeviceTherm.New()
	res.DeviceVersion.New()
	res.DeviceSens.New()
	res.DeviceNote.New()

}

//GetByRedis
func (res *UltDeviceType) GetByRedis(dbIndex string) ResultType {
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
	res.DeviceSens.DeviceId = res.DeviceId
	resultVal = res.DeviceSens.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceBattery.DeviceId = res.DeviceId
	resultVal = res.DeviceBattery.GetByRedis(dbIndex)
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
func (res *UltDeviceType) GetByRedisBySerial(serial string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_SERIAL_ULT_DEVICE, serial)
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
func (res *UltDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
