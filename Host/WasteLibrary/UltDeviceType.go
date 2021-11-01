package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceType
type UltDeviceType struct {
	DeviceId        float64
	CustomerId      float64
	ContainerNo     string
	ContainerType   string
	DeviceType      string
	SerialNumber    string
	StatusTime      string
	AliveStatus     string
	AliveLastOkTime string
	Latitude        float64
	Longitude       float64
	GpsTime         string
	AlarmStatus     string
	AlarmTime       string
	AlarmType       string
	Alarm           string
	Therm           string
	ThermTime       string
	ThermStatus     string
	Active          string
	CreateTime      string
	Battery         string
	BatteryStatus   string
	BatteryTime     string
	UltTime         string
	UltRange        float64
	UltStatus       string
	Imei            string
	Imsi            string
	FirmwareVersion string
	OldLatitude     float64
	OldLongitude    float64
}

//New
func (res *UltDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.ContainerNo = ""
	res.ContainerType = CONTAINER_TYPE_NONE
	res.DeviceType = ULT_DEVICE_TYPE_NONE
	res.SerialNumber = ""
	res.StatusTime = ""
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = ""
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = ""
	res.AlarmStatus = ALARM_STATU_NONE
	res.AlarmTime = ""
	res.AlarmType = ALARM_NONE
	res.Alarm = ""
	res.Therm = "0"
	res.ThermTime = ""
	res.ThermStatus = THERM_STATU_NONE
	res.Active = STATU_ACTIVE
	res.CreateTime = ""
	res.Battery = "0"
	res.BatteryStatus = BATTERY_STATU_NONE
	res.BatteryTime = ""
	res.UltTime = ""
	res.UltRange = 0
	res.UltStatus = ULT_STATU_NONE
	res.Imei = ""
	res.Imsi = ""
	res.FirmwareVersion = "1"
	res.OldLatitude = 0
	res.OldLongitude = 0
}

//ToId String
func (res UltDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res UltDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res UltDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceType) ToString() string {
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

//SelectSQL
func (res UltDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerId,
	ContainerNo,
	ContainerType,
	DeviceType,
	SerialNumber,
	StatusTime,
	AliveStatus,
	AliveLastOkTime,
	Latitude,
	Longitude,
	GpsTime,
	AlarmStatus,
	AlarmTime,
	AlarmType,
	Alarm,
	Therm,
	ThermTime,
	ThermStatus,
	Active,
	CreateTime,
	Battery,
	BatteryStatus,
	BatteryTime,
	UltTime,
	UltRange,
	UltStatus,
	Imei,
	Imsi,
	FirmwareVersion,
	OldLatitude,
	OldLongitude 
	 FROM public.ult_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_devices 
	(SerialNumber,CustomerId,Imei,Imsi) 
	  VALUES ('%s',%f,'%s','%s')   
	  RETURNING DeviceId;`,
		res.SerialNumber, res.CustomerId, res.Imei, res.Imsi)
}

//InsertDeviceDataSQL
func (res UltDeviceType) InsertDeviceDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_devicedata 
	(DeviceId,
		CustomerId,
		ContainerNo,
		ContainerType,
		DeviceType,
		SerialNumber,
		StatusTime,
		AliveStatus,
		AliveLastOkTime,
		Latitude,
		Longitude,
		GpsTime,
		AlarmStatus,
		AlarmTime,
		AlarmType,
		Alarm,
		Therm,
		ThermTime,
		ThermStatus,
		Active,
		CreateTime,
		Battery,
		BatteryStatus,
		BatteryTime,
		UltTime,
		UltRange,
		UltStatus,
		Imei,
		Imsi,
		FirmwareVersion,
		OldLatitude,
		OldLongitude)  
	  VALUES (%f,%f
		,'%s','%s','%s','%s','%s'
		,'%s','%s'
	    ,%f,%f
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s'
		,%f
		,'%s','%s','%s','%s'
		,%f,%f) 
	  RETURNING DataId;`,
		res.DeviceId,
		res.CustomerId,
		res.ContainerNo,
		res.ContainerType,
		res.DeviceType,
		res.SerialNumber,
		res.StatusTime,
		res.AliveStatus,
		res.AliveLastOkTime,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.AlarmStatus,
		res.AlarmTime,
		res.AlarmType,
		res.Alarm,
		res.Therm,
		res.ThermTime,
		res.ThermStatus,
		res.Battery,
		res.BatteryStatus,
		res.BatteryTime,
		res.UltTime,
		res.UltRange,
		res.UltStatus,
		res.Imei,
		res.Imsi,
		res.FirmwareVersion,
		res.OldLatitude,
		res.OldLongitude)
}

//UpdateBasicSQL
func (res UltDeviceType) UpdateBasicSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices 
				SET DeviceType='%s',ContainerNo='%s',ContainerType='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.DeviceType, res.ContainerNo, res.ContainerType, res.DeviceId)
}

//UpdateCustomerSQL
func (res UltDeviceType) UpdateCustomerSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices 
				SET CustomerId=%f 
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.CustomerId, res.DeviceId)
}

//UpdateVersionSQL
func (res UltDeviceType) UpdateVersionSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices 
				SET FirmwareVersion='%s' 
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.FirmwareVersion, res.DeviceId)
}

//UpdateGpsSQL
func (res UltDeviceType) UpdateGpsSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices
			   SET GpsTime='%s',Latitude=%f,Longitude=%f,OldLatitude=%f,OldLongitude=%f
			   WHERE DeviceId=%f AND CustomerId=%f 
			   RETURNING DeviceId;`, res.GpsTime,
		res.Latitude, res.Longitude, res.OldLatitude, res.OldLongitude,
		res.DeviceId, res.CustomerId)
}

//UpdateThermSQL
func (res UltDeviceType) UpdateThermSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_devices
		 SET Therm='%s',ThermTime='%s'
		 WHERE DeviceId=%f AND CustomerId=%f 
		 RETURNING DeviceId;`,
		res.Therm, res.ThermTime, res.DeviceId, res.CustomerId)
}

//SelectWithDb
func (res UltDeviceType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(&res.CustomerId,
		&res.ContainerNo,
		&res.ContainerType,
		&res.DeviceType,
		&res.SerialNumber,
		&res.StatusTime,
		&res.AliveStatus,
		&res.AliveLastOkTime,
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime,
		&res.AlarmStatus,
		&res.AlarmTime,
		&res.AlarmType,
		&res.Alarm,
		&res.Therm,
		&res.ThermTime,
		&res.ThermStatus,
		&res.Active,
		&res.CreateTime,
		&res.Battery,
		&res.BatteryStatus,
		&res.BatteryTime,
		&res.UltTime,
		&res.UltRange,
		&res.UltStatus,
		&res.Imei,
		&res.Imsi,
		&res.FirmwareVersion,
		&res.OldLatitude,
		&res.OldLongitude)
	return errDb
}
