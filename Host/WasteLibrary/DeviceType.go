package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//DeviceType
type DeviceType struct {
	DeviceId              float64
	CustomerId            float64
	DeviceName            string
	ContainerNo           string
	ContainerType         string
	DeviceType            string
	SerialNumber          string
	DeviceStatus          string
	StatusTime            string
	AliveStatus           string
	AliveLastOkTime       string
	Latitude              float64
	Longitude             float64
	GpsTime               string
	AlarmStatus           string
	AlarmTime             string
	AlarmType             string
	Alarm                 string
	Therm                 string
	ThermTime             string
	ThermStatus           string
	Active                string
	CreateTime            string
	ReaderAppStatus       string
	ReaderAppLastOkTime   string
	ReaderConnStatus      string
	ReaderConnLastOkTime  string
	ReaderStatus          string
	ReaderLastOkTime      string
	CamAppStatus          string
	CamAppLastOkTime      string
	CamConnStatus         string
	CamConnLastOkTime     string
	CamStatus             string
	CamLastOkTime         string
	GpsAppStatus          string
	GpsAppLastOkTime      string
	GpsConnStatus         string
	GpsConnLastOkTime     string
	GpsStatus             string
	GpsLastOkTime         string
	ThermAppStatus        string
	ThermAppLastOkTime    string
	TransferAppStatus     string
	TransferAppLastOkTime string
	SystemAppStatus       string
	SystemAppLastOkTime   string
	UpdaterAppStatus      string
	UpdaterAppLastOkTime  string
	ContactStatus         string
	ContactLastOkTime     string
	Speed                 float64
	Battery               string
	BatteryStatus         string
	BatteryTime           string
	UltTime               string
	UltRange              float64
	UltStatus             string
	TotalGlassCount       float64
	TotalPlasticCount     float64
	TotalMetalCount       float64
	DailyGlassCount       float64
	DailyPlasticCount     float64
	DailyMetalCount       float64
	RecyTime              string
	MotorAppStatus        string
	MotorAppLastOkTime    string
	MotorConnStatus       string
	MotorConnLastOkTime   string
	MotorStatus           string
}

//New
func (res DeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.DeviceName = ""
	res.ContainerNo = ""
	res.ContainerType = CONTAINER_TYPE_NONE
	res.DeviceType = DEVICE_TYPE_NONE
	res.SerialNumber = ""
	res.DeviceStatus = DEVICE_STATU_NONE
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
	res.ReaderAppStatus = STATU_PASSIVE
	res.ReaderAppLastOkTime = ""
	res.ReaderConnStatus = STATU_PASSIVE
	res.ReaderConnLastOkTime = ""
	res.ReaderStatus = STATU_PASSIVE
	res.ReaderLastOkTime = ""
	res.CamAppStatus = STATU_PASSIVE
	res.CamAppLastOkTime = ""
	res.CamConnStatus = STATU_PASSIVE
	res.CamConnLastOkTime = ""
	res.CamStatus = STATU_PASSIVE
	res.CamLastOkTime = ""
	res.GpsAppStatus = STATU_PASSIVE
	res.GpsAppLastOkTime = ""
	res.GpsConnStatus = STATU_PASSIVE
	res.GpsConnLastOkTime = ""
	res.GpsStatus = STATU_PASSIVE
	res.GpsLastOkTime = ""
	res.ThermAppStatus = STATU_PASSIVE
	res.ThermAppLastOkTime = ""
	res.TransferAppStatus = STATU_PASSIVE
	res.TransferAppLastOkTime = ""
	res.SystemAppStatus = STATU_PASSIVE
	res.SystemAppLastOkTime = ""
	res.UpdaterAppStatus = STATU_PASSIVE
	res.UpdaterAppLastOkTime = ""
	res.ContactStatus = STATU_PASSIVE
	res.ContactLastOkTime = ""
	res.Speed = -1
	res.Battery = "0"
	res.BatteryStatus = BATTERY_STATU_NONE
	res.BatteryTime = ""
	res.UltTime = ""
	res.UltRange = 0
	res.UltStatus = ULT_STATU_NONE
	res.TotalGlassCount = 0
	res.TotalPlasticCount = 0
	res.TotalMetalCount = 0
	res.DailyGlassCount = 0
	res.DailyPlasticCount = 0
	res.DailyMetalCount = 0
	res.RecyTime = ""
	res.MotorAppStatus = STATU_PASSIVE
	res.MotorAppLastOkTime = ""
	res.MotorConnStatus = STATU_PASSIVE
	res.MotorConnLastOkTime = ""
	res.MotorStatus = STATU_PASSIVE
}

//ToId String
func (res DeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res DeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res DeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res DeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To DeviceType
func ByteToDeviceType(retByte []byte) DeviceType {
	var retVal DeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To DeviceType
func StringToDeviceType(retStr string) DeviceType {
	return ByteToDeviceType([]byte(retStr))
}

//SelectSQL
func (res DeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerId            ,
	DeviceName            ,
	ContainerNo           ,
	ContainerType         ,
	DeviceType            ,
	SerialNumber          ,
	DeviceStatus          ,
	StatusTime            ,
	AliveStatus           ,
	AliveLastOkTime       ,
	Latitude              ,
	Longitude             ,
	GpsTime               ,
	AlarmStatus           ,
	AlarmTime             ,
	AlarmType             ,
	Alarm                 ,
	Therm                 ,
	ThermTime             ,
	ThermStatus           ,
	Active                ,
	CreateTime            ,
	ReaderAppStatus       ,
	ReaderAppLastOkTime   ,
	ReaderConnStatus      ,
	ReaderConnLastOkTime  ,
	ReaderStatus          ,
	ReaderLastOkTime      ,
	CamAppStatus          ,
	CamAppLastOkTime      ,
	CamConnStatus         ,
	CamConnLastOkTime     ,
	CamStatus             ,
	CamLastOkTime         ,
	GpsAppStatus          ,
	GpsAppLastOkTime      ,
	GpsConnStatus         ,
	GpsConnLastOkTime     ,
	GpsStatus             ,
	GpsLastOkTime         ,
	ThermAppStatus        ,
	ThermAppLastOkTime    ,
	TransferAppStatus     ,
	TransferAppLastOkTime ,
	ContactStatus         ,
	ContactLastOkTime     ,
	Speed                 ,
	Battery               ,
	BatteryStatus         ,
	BatteryTime           ,
	UltTime               ,
	UltRange              ,
	UltStatus             ,
	TotalGlassCount       ,
	TotalPlasticCount     ,
	TotalMetalCount       ,
	DailyGlassCount       ,
	DailyPlasticCount     ,
	DailyMetalCount       ,
	RecyTime              ,
	MotorAppStatus        ,
	MotorAppLastOkTime    ,
	MotorConnStatus       ,
	MotorConnLastOkTime   ,
	MotorStatus           ,
	UpdaterAppStatus     ,
	UpdaterAppLastOkTime ,
	SystemAppStatus     ,
	SystemAppLastOkTime 
	 FROM public.devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res DeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.devices 
	(DeviceType,SerialNumber,DeviceName,CustomerId,ContainerNo,ContainerType) 
	  VALUES ('%s','%s','%s',%f,'%s','%s')   
	  RETURNING DeviceId;`,
		res.DeviceType, res.SerialNumber, res.DeviceName,
		res.CustomerId, res.ContainerNo, res.ContainerType)
}

//InsertDeviceDataSQL
func (res DeviceType) InsertDeviceDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.devicedata 
	(DeviceId              ,
		CustomerId            ,
		DeviceName            ,
		ContainerNo           ,
		ContainerType         ,
		DeviceType            ,
		SerialNumber          ,
		DeviceStatus          ,
		StatusTime            ,
		AliveStatus           ,
		AliveLastOkTime       ,
		Latitude              ,
		Longitude             ,
		GpsTime               ,
		AlarmStatus           ,
		AlarmTime             ,
		AlarmType             ,
		Alarm                 ,
		Therm                 ,
		ThermTime             ,
		ThermStatus           ,
		ReaderAppStatus       ,
		ReaderAppLastOkTime   ,
		ReaderConnStatus      ,
		ReaderConnLastOkTime  ,
		ReaderStatus          ,
		ReaderLastOkTime      ,
		CamAppStatus          ,
		CamAppLastOkTime      ,
		CamConnStatus         ,
		CamConnLastOkTime     ,
		CamStatus             ,
		CamLastOkTime         ,
		GpsAppStatus          ,
		GpsAppLastOkTime      ,
		GpsConnStatus         ,
		GpsConnLastOkTime     ,
		GpsStatus             ,
		GpsLastOkTime         ,
		ThermAppStatus        ,
		ThermAppLastOkTime    ,
		TransferAppStatus     ,
		TransferAppLastOkTime ,
		ContactStatus         ,
		ContactLastOkTime     ,
		Speed                 ,
		Battery               ,
		BatteryStatus         ,
		BatteryTime           ,
		UltTime               ,
		UltRange              ,
		UltStatus             ,
		TotalGlassCount       ,
		TotalPlasticCount     ,
		TotalMetalCount       ,
		DailyGlassCount       ,
		DailyPlasticCount     ,
		DailyMetalCount       ,
		RecyTime              ,
		MotorAppStatus        ,
		MotorAppLastOkTime    ,
		MotorConnStatus       ,
		MotorConnLastOkTime   ,
		MotorStatus           ,
		UpdaterAppStatus     ,
		UpdaterAppLastOkTime ,
		SystemAppStatus     ,
		SystemAppLastOkTime 
	) 
	  VALUES (%f,%f,'%s','%s','%s','%s','%s','%s','%s','%s',
	  '%s',%f,%f,'%s','%s','%s','%s','%s',
	  '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',
	  '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',
	  '%s','%s','%s','%s','%s','%s','%s',%f,'%s','%s',
	  '%s','%s',%f,'%s',%f,%f,%f,%f,%f,%f,
	  '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s') 
	  RETURNING DataId;`,
		res.DeviceId,
		res.CustomerId,
		res.DeviceName,
		res.ContainerNo,
		res.ContainerType,
		res.DeviceType,
		res.SerialNumber,
		res.DeviceStatus,
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
		res.ReaderAppStatus,
		res.ReaderAppLastOkTime,
		res.ReaderConnStatus,
		res.ReaderConnLastOkTime,
		res.ReaderStatus,
		res.ReaderLastOkTime,
		res.CamAppStatus,
		res.CamAppLastOkTime,
		res.CamConnStatus,
		res.CamConnLastOkTime,
		res.CamStatus,
		res.CamLastOkTime,
		res.GpsAppStatus,
		res.GpsAppLastOkTime,
		res.GpsConnStatus,
		res.GpsConnLastOkTime,
		res.GpsStatus,
		res.GpsLastOkTime,
		res.ThermAppStatus,
		res.ThermAppLastOkTime,
		res.TransferAppStatus,
		res.TransferAppLastOkTime,
		res.ContactStatus,
		res.ContactLastOkTime,
		res.Speed,
		res.Battery,
		res.BatteryStatus,
		res.BatteryTime,
		res.UltTime,
		res.UltRange,
		res.UltStatus,
		res.TotalGlassCount,
		res.TotalPlasticCount,
		res.TotalMetalCount,
		res.DailyGlassCount,
		res.DailyPlasticCount,
		res.DailyMetalCount,
		res.RecyTime,
		res.MotorAppStatus,
		res.MotorAppLastOkTime,
		res.MotorConnStatus,
		res.MotorConnLastOkTime,
		res.MotorStatus,
		res.UpdaterAppStatus,
		res.UpdaterAppLastOkTime,
		res.SystemAppStatus,
		res.SystemAppLastOkTime)
}

//UpdateSQL
func (res DeviceType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.devices 
				SET DeviceType='%s',SerialNumber='%s',DeviceName='%s',CustomerId=%f 
				,ContainerNo='%s',ContainerType='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.DeviceType, res.SerialNumber, res.DeviceName,
		res.CustomerId, res.ContainerNo, res.ContainerType, res.DeviceId)
}

//UpdateGpsSQL
func (res DeviceType) UpdateGpsSQL() string {
	return fmt.Sprintf(`UPDATE public.devices
			   SET GpsTime='%s',Latitude=%f,Longitude=%f,Speed=%f
			   WHERE DeviceId=%f AND CustomerId=%f 
			   RETURNING DeviceId;`, res.GpsTime,
		res.Latitude, res.Longitude, res.Speed,
		res.DeviceId, res.CustomerId)
}

//UpdateStatuSQL
func (res DeviceType) UpdateStatuSQL() string {
	var execSqlExt = ""
	if res.ReaderAppStatus == STATU_ACTIVE {
		execSqlExt += ",ReaderAppLastOkTime='" + res.ReaderAppLastOkTime + "'"
	}
	if res.ReaderConnStatus == STATU_ACTIVE {
		execSqlExt += ",ReaderConnLastOkTime='" + res.ReaderConnLastOkTime + "'"
	}
	if res.ReaderStatus == STATU_ACTIVE {
		execSqlExt += ",ReaderLastOkTime='" + res.ReaderLastOkTime + "'"
	}
	if res.CamAppStatus == STATU_ACTIVE {
		execSqlExt += ",CamAppLastOkTime='" + res.CamAppLastOkTime + "'"
	}
	if res.CamConnStatus == STATU_ACTIVE {
		execSqlExt += ",CamConnLastOkTime='" + res.CamConnLastOkTime + "'"
	}
	if res.CamStatus == STATU_ACTIVE {
		execSqlExt += ",CamLastOkTime='" + res.CamLastOkTime + "'"
	}

	if res.GpsAppStatus == STATU_ACTIVE {
		execSqlExt += ",GpsAppLastOkTime='" + res.GpsAppLastOkTime + "'"
	}
	if res.GpsConnStatus == STATU_ACTIVE {
		execSqlExt += ",GpsConnLastOkTime='" + res.GpsConnLastOkTime + "'"
	}
	if res.GpsStatus == STATU_ACTIVE {
		execSqlExt += ",GpsLastOkTime='" + res.GpsLastOkTime + "'"
	}

	if res.ThermAppStatus == STATU_ACTIVE {
		execSqlExt += ",ThermAppLastOkTime='" + res.ThermAppLastOkTime + "'"
	}
	if res.TransferAppStatus == STATU_ACTIVE {
		execSqlExt += ",TransferAppLastOkTime='" + res.TransferAppLastOkTime + "'"
	}
	if res.SystemAppStatus == STATU_ACTIVE {
		execSqlExt += ",SystemAppLastOkTime='" + res.SystemAppLastOkTime + "'"
	}
	if res.UpdaterAppStatus == STATU_ACTIVE {
		execSqlExt += ",UpdaterAppLastOkTime='" + res.UpdaterAppLastOkTime + "'"
	}
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}
	if res.ContactStatus == STATU_ACTIVE {
		execSqlExt += ",ContactLastOkTime='" + res.ContactLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.devices
	SET StatusTime='%s',
	ReaderAppStatus='%s',ReaderConnStatus='%s',ReaderStatus='%s',CamAppStatus='%s',CamConnStatus='%s',
	CamStatus='%s',GpsAppStatus='%s',GpsConnStatus='%s',GpsStatus='%s',ThermAppStatus='%s',
	TransferAppStatus='%s',AliveStatus='%s',ContactStatus='%s',
	UpdaterAppStatus='%s',SystemAppStatus='%s'`+execSqlExt+`
   WHERE DeviceId=%f AND CustomerId=%f 
   RETURNING DeviceId;`, res.StatusTime,
		res.ReaderAppStatus, res.ReaderConnStatus, res.ReaderStatus,
		res.CamAppStatus, res.CamConnStatus, res.CamStatus,
		res.GpsAppStatus, res.GpsConnStatus, res.GpsStatus,
		res.ThermAppStatus, res.TransferAppStatus, res.AliveStatus,
		res.ContactStatus, res.UpdaterAppStatus, res.SystemAppStatus, res.DeviceId, res.CustomerId)
}

//UpdateThermSQL
func (res DeviceType) UpdateThermSQL() string {
	return fmt.Sprintf(`UPDATE public.devices
		 SET Therm='%s',ThermTime='%s'
		 WHERE DeviceId=%f AND CustomerId=%f 
		 RETURNING DeviceId;`,
		res.Therm, res.ThermTime, res.DeviceId, res.CustomerId)
}
