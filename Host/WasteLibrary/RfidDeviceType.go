package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RfidDeviceType
type RfidDeviceType struct {
	DeviceId       float64
	DeviceMain     RfidDeviceMainType
	DeviceBase     RfidDeviceBaseType
	DeviceStatu    RfidDeviceStatuType
	DeviceGps      RfidDeviceGpsType
	DeviceAlarm    RfidDeviceAlarmType
	DeviceTherm    RfidDeviceThermType
	DeviceVersion  RfidDeviceVersionType
	DeviceDetail   RfidDeviceDetailType
	DeviceWorkHour RfidDeviceWorkHourType
}

//New
func (res *RfidDeviceType) New() {
	res.DeviceId = 0
	res.DeviceMain.New()
	res.DeviceBase.New()
	res.DeviceGps.New()
	res.DeviceTherm.New()
	res.DeviceVersion.New()
	res.DeviceAlarm.New()
	res.DeviceStatu.New()
	res.DeviceDetail.New()
	res.DeviceWorkHour.New()
}

//GetAll
func (res *RfidDeviceType) GetAll() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RFID_MAIN_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceMain = StringToRfidDeviceMainType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_BASE_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceBase = StringToRfidDeviceBaseType(resultVal.Retval.(string))
		res.DeviceBase.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_GPS_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceGps = StringToRfidDeviceGpsType(resultVal.Retval.(string))
		res.DeviceGps.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_THERM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceTherm = StringToRfidDeviceThermType(resultVal.Retval.(string))
		res.DeviceTherm.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_VERSION_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceVersion = StringToRfidDeviceVersionType(resultVal.Retval.(string))
		res.DeviceVersion.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_ALARM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceAlarm = StringToRfidDeviceAlarmType(resultVal.Retval.(string))
		res.DeviceAlarm.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_STATU_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceStatu = StringToRfidDeviceStatuType(resultVal.Retval.(string))
		res.DeviceStatu.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_DETAIL_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceDetail = StringToRfidDeviceDetailType(resultVal.Retval.(string))
		res.DeviceDetail.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_RFID_WORKHOUR_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.DeviceWorkHour = StringToRfidDeviceWorkHourType(resultVal.Retval.(string))
		res.DeviceWorkHour.NewData = false
	} else {
		return resultVal
	}
	return resultVal
}

//ToId String
func (res *RfidDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceType
func ByteToRfidDeviceType(retByte []byte) RfidDeviceType {
	var retVal RfidDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceType
func StringToRfidDeviceType(retStr string) RfidDeviceType {
	return ByteToRfidDeviceType([]byte(retStr))
}
