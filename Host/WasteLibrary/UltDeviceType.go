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

//GetAll
func (res *UltDeviceType) GetAll() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_ULT_MAIN_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceMain = StringToUltDeviceMainType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_BASE_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceBase = StringToUltDeviceBaseType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_GPS_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceGps = StringToUltDeviceGpsType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_THERM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceTherm = StringToUltDeviceThermType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_VERSION_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceVersion = StringToUltDeviceVersionType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_ALARM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceAlarm = StringToUltDeviceAlarmType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_STATU_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceStatu = StringToUltDeviceStatuType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_SENS_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceSens = StringToUltDeviceSensType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_ULT_BATTERY_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceBattery = StringToUltDeviceBatteryType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
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
