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

}

//GetByRedis
func (res *UltDeviceType) GetByRedis(dbIndex int) ResultType {
	var resultVal ResultType

	res.DeviceMain.DeviceId = res.DeviceId
	resultVal = res.DeviceMain.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceBase.DeviceId = res.DeviceId
	resultVal = res.DeviceBase.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceGps.DeviceId = res.DeviceId
	resultVal = res.DeviceGps.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceTherm.DeviceId = res.DeviceId
	resultVal = res.DeviceTherm.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceVersion.DeviceId = res.DeviceId
	resultVal = res.DeviceVersion.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceAlarm.DeviceId = res.DeviceId
	resultVal = res.DeviceAlarm.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceStatu.DeviceId = res.DeviceId
	resultVal = res.DeviceStatu.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceSens.DeviceId = res.DeviceId
	resultVal = res.DeviceSens.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceBattery.DeviceId = res.DeviceId
	resultVal = res.DeviceBattery.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisBySerial
func (res *UltDeviceType) GetByRedisBySerial(serial string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_SERIAL_ULT_DEVICE, serial)
	if resultVal.Result == RESULT_OK {
		var deviceId string = resultVal.Retval.(string)
		res.DeviceId = StringIdToFloat64(deviceId)
		resultVal = res.GetByRedis()
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

//Byte To UltDeviceType
func ByteToUltDeviceType(retByte []byte) UltDeviceType {
	var retVal UltDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceType
func StringToUltDeviceType(retStr string) UltDeviceType {
	return ByteToUltDeviceType([]byte(retStr))
}
