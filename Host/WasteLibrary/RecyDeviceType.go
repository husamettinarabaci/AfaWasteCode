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
}

//GetAll
func (res *RecyDeviceType) GetAll() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RECY_MAIN_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceMain = StringToRecyDeviceMainType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_BASE_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceBase = StringToRecyDeviceBaseType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_GPS_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceGps = StringToRecyDeviceGpsType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_THERM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceTherm = StringToRecyDeviceThermType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_VERSION_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceVersion = StringToRecyDeviceVersionType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_ALARM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceAlarm = StringToRecyDeviceAlarmType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_STATU_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceStatu = StringToRecyDeviceStatuType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RECY_DETAIL_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceDetail = StringToRecyDeviceDetailType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
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

//Byte To RecyDeviceType
func ByteToRecyDeviceType(retByte []byte) RecyDeviceType {
	var retVal RecyDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceType
func StringToRecyDeviceType(retStr string) RecyDeviceType {
	return ByteToRecyDeviceType([]byte(retStr))
}
