package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RecyDeviceType
type RecyDeviceType struct {
	DeviceId              float64
	CustomerId            float64
	ContainerNo           string
	DeviceType            string
	SerialNumber          string
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
	ThermAppStatus        string
	ThermAppLastOkTime    string
	TransferAppStatus     string
	TransferAppLastOkTime string
	SystemAppStatus       string
	SystemAppLastOkTime   string
	UpdaterAppStatus      string
	UpdaterAppLastOkTime  string
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
	MotorLastOkTime       string
	WebAppStatus          string
	WebAppLastOkTime      string
	WebAppVersion         string
	MotorAppVersion       string
	ThermAppVersion       string
	TransferAppVersion    string
	CheckerAppVersion     string
	CamAppVersion         string
	ReaderAppVersion      string
	SystemAppVersion      string
}

//New
func (res RecyDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.ContainerNo = ""
	res.DeviceType = DEVICE_TYPE_NONE
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
	res.ThermAppStatus = STATU_PASSIVE
	res.ThermAppLastOkTime = ""
	res.TransferAppStatus = STATU_PASSIVE
	res.TransferAppLastOkTime = ""
	res.SystemAppStatus = STATU_PASSIVE
	res.SystemAppLastOkTime = ""
	res.UpdaterAppStatus = STATU_PASSIVE
	res.UpdaterAppLastOkTime = ""
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
	res.MotorLastOkTime = ""
	res.WebAppStatus = STATU_PASSIVE
	res.WebAppLastOkTime = ""
	res.WebAppVersion = "1"
	res.MotorAppVersion = "1"
	res.ThermAppVersion = "1"
	res.TransferAppVersion = "1"
	res.CheckerAppVersion = "1"
	res.CamAppVersion = "1"
	res.ReaderAppVersion = "1"
	res.SystemAppVersion = "1"
}

//ToId String
func (res RecyDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res RecyDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res RecyDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RecyDeviceType) ToString() string {
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

//SelectSQL
func (res RecyDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerId,
	ContainerNo,
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
	ReaderAppStatus,
	ReaderAppLastOkTime,
	ReaderConnStatus,
	ReaderConnLastOkTime,
	ReaderStatus,
	ReaderLastOkTime,
	CamAppStatus,
	CamAppLastOkTime,
	CamConnStatus,
	CamConnLastOkTime,
	CamStatus,
	CamLastOkTime,
	ThermAppStatus,
	ThermAppLastOkTime,
	TransferAppStatus,
	TransferAppLastOkTime,
	SystemAppStatus,
	SystemAppLastOkTime,
	UpdaterAppStatus,
	UpdaterAppLastOkTime,
	TotalGlassCount,
	TotalPlasticCount,
	TotalMetalCount,
	DailyGlassCount,
	DailyPlasticCount,
	DailyMetalCount,
	RecyTime,
	MotorAppStatus,
	MotorAppLastOkTime,
	MotorConnStatus,
	MotorConnLastOkTime,
	MotorStatus,
	MotorLastOkTime,
	WebAppStatus,
	WebAppLastOkTime,
	WebAppVersion,
	MotorAppVersion,
	ThermAppVersion,
	TransferAppVersion,
	CheckerAppVersion,    
	CamAppVersion,
	ReaderAppVersion,
	SystemAppVersion 
	 FROM public.recy_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RecyDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_devices 
	(SerialNumber,CustomerId) 
	  VALUES ('%s',%f)   
	  RETURNING DeviceId;`,
		res.SerialNumber, res.CustomerId)
}

//InsertDeviceDataSQL
func (res RecyDeviceType) InsertDeviceDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.devicedata 
	(DeviceId,
		CustomerId,
		ContainerNo,
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
		ReaderAppStatus,
		ReaderAppLastOkTime,
		ReaderConnStatus,
		ReaderConnLastOkTime,
		ReaderStatus,
		ReaderLastOkTime,
		CamAppStatus,
		CamAppLastOkTime,
		CamConnStatus,
		CamConnLastOkTime,
		CamStatus,
		CamLastOkTime,
		ThermAppStatus,
		ThermAppLastOkTime,
		TransferAppStatus,
		TransferAppLastOkTime,
		SystemAppStatus,
		SystemAppLastOkTime,
		UpdaterAppStatus,
		UpdaterAppLastOkTime,
		TotalGlassCount,
		TotalPlasticCount,
		TotalMetalCount,
		DailyGlassCount,
		DailyPlasticCount,
		DailyMetalCount,
		RecyTime,
		MotorAppStatus,
		MotorAppLastOkTime,
		MotorConnStatus,
		MotorConnLastOkTime,
		MotorStatus,
		MotorLastOkTime,
		WebAppStatus,
		WebAppLastOkTime,
		WebAppVersion,
		MotorAppVersion,
		ThermAppVersion,
		TransferAppVersion,
		CheckerAppVersion,    
		CamAppVersion,
		ReaderAppVersion,
		SystemAppVersion) 
	  VALUES (%f,%f
		,'%s','%s','%s','%s','%s'
		,'%s'
	    ,%f,%f
		,'%s','%s','%s','%s','%s'
		,%f
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s'
		,%f,%f,%f,%f,%f
		,%f
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
		,'%s','%s') 
	  RETURNING DataId;`,
		res.DeviceId,
		res.CustomerId,
		res.ContainerNo,
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
		res.Active,
		res.CreateTime,
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
		res.ThermAppStatus,
		res.ThermAppLastOkTime,
		res.TransferAppStatus,
		res.TransferAppLastOkTime,
		res.SystemAppStatus,
		res.SystemAppLastOkTime,
		res.UpdaterAppStatus,
		res.UpdaterAppLastOkTime,
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
		res.MotorLastOkTime,
		res.WebAppStatus,
		res.WebAppLastOkTime,
		res.WebAppVersion,
		res.MotorAppVersion,
		res.ThermAppVersion,
		res.TransferAppVersion,
		res.CheckerAppVersion,
		res.CamAppVersion,
		res.ReaderAppVersion,
		res.SystemAppVersion)
}

//UpdateBasicSQL
func (res RecyDeviceType) UpdateBasicSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices 
				SET DeviceType='%s',ContainerNo='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.DeviceType, res.ContainerNo, res.DeviceId)
}

//UpdateCustomerSQL
func (res RecyDeviceType) UpdateCustomerSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices 
				SET CustomerId=%f 
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.CustomerId, res.DeviceId)
}

//UpdateVersionSQL
func (res RecyDeviceType) UpdateVersionSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices 
				SET WebAppVersion='%s',MotorAppVersion='%s',ThermAppVersion='%s',TransferAppVersion='%s' 
				,CheckerAppVersion='%s',CamAppVersion='%s',ReaderAppVersion='%s'
				,SystemAppVersion='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.WebAppVersion, res.MotorAppVersion, res.ThermAppVersion, res.TransferAppVersion,
		res.CheckerAppVersion, res.CamAppVersion, res.ReaderAppVersion, res.SystemAppVersion, res.DeviceId)
}

//UpdateGpsSQL
func (res RecyDeviceType) UpdateGpsSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices
			   SET GpsTime='%s',Latitude=%f,Longitude=%f
			   WHERE DeviceId=%f AND CustomerId=%f 
			   RETURNING DeviceId;`, res.GpsTime,
		res.Latitude, res.Longitude,
		res.DeviceId, res.CustomerId)
}

//UpdateStatuSQL
func (res RecyDeviceType) UpdateStatuSQL() string {
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

	if res.MotorAppStatus == STATU_ACTIVE {
		execSqlExt += ",MotorAppLastOkTime='" + res.MotorAppLastOkTime + "'"
	}
	if res.MotorConnStatus == STATU_ACTIVE {
		execSqlExt += ",MotorConnLastOkTime='" + res.MotorConnLastOkTime + "'"
	}
	if res.MotorStatus == STATU_ACTIVE {
		execSqlExt += ",MotorLastOkTime='" + res.MotorLastOkTime + "'"
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
	if res.WebAppStatus == STATU_ACTIVE {
		execSqlExt += ",WebAppLastOkTime='" + res.WebAppLastOkTime + "'"
	}
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.recy_devices
	SET StatusTime='%s',
	ReaderAppStatus='%s',ReaderConnStatus='%s',ReaderStatus='%s',CamAppStatus='%s',CamConnStatus='%s',
	CamStatus='%s',MotorAppStatus='%s',MotorConnStatus='%s',MotorStatus='%s',ThermAppStatus='%s',
	TransferAppStatus='%s',AliveStatus='%s',
	UpdaterAppStatus='%s',SystemAppStatus='%s',WebAppStatus='%s'`+execSqlExt+`
   WHERE DeviceId=%f AND CustomerId=%f 
   RETURNING DeviceId;`, res.StatusTime,
		res.ReaderAppStatus, res.ReaderConnStatus, res.ReaderStatus,
		res.CamAppStatus, res.CamConnStatus, res.CamStatus,
		res.MotorAppStatus, res.MotorConnStatus, res.MotorStatus,
		res.ThermAppStatus, res.TransferAppStatus, res.AliveStatus,
		res.UpdaterAppStatus, res.SystemAppStatus, res.WebAppStatus, res.DeviceId, res.CustomerId)
}

//UpdateThermSQL
func (res RecyDeviceType) UpdateThermSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_devices
		 SET Therm='%s',ThermTime='%s'
		 WHERE DeviceId=%f AND CustomerId=%f 
		 RETURNING DeviceId;`,
		res.Therm, res.ThermTime, res.DeviceId, res.CustomerId)
}
