package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceType
type RfidDeviceType struct {
	DeviceId              float64
	CustomerId            float64
	DeviceType            string
	TruckType             string
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
	PlateNo               string
	DriverName            string
	DriverSurName         string
	GpsAppVersion         string
	ThermAppVersion       string
	TransferAppVersion    string
	CheckerAppVersion     string
	CamAppVersion         string
	ReaderAppVersion      string
	SystemAppVersion      string
}

//New
func (res *RfidDeviceType) New() {
	res.DeviceId = 0
	res.CustomerId = 0
	res.DeviceType = RFID_DEVICE_TYPE_NONE
	res.TruckType = TRUCK_TYPE_NONE
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
	res.PlateNo = ""
	res.DriverName = ""
	res.DriverSurName = ""
	res.GpsAppVersion = "1"
	res.ThermAppVersion = "1"
	res.TransferAppVersion = "1"
	res.CheckerAppVersion = "1"
	res.CamAppVersion = "1"
	res.ReaderAppVersion = "1"
	res.SystemAppVersion = "1"
}

//ToId String
func (res RfidDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToCustomerId String
func (res RfidDeviceType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res RfidDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RfidDeviceType) ToString() string {
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

//SelectSQL
func (res RfidDeviceType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerId,
	DeviceType,
	TruckType,
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
	GpsAppStatus,
	GpsAppLastOkTime,
	GpsConnStatus,
	GpsConnLastOkTime,
	GpsStatus,
	GpsLastOkTime,
	ThermAppStatus,
	ThermAppLastOkTime,
	TransferAppStatus,
	TransferAppLastOkTime,
	ContactStatus,
	ContactLastOkTime,
	Speed,
	UpdaterAppStatus,
	UpdaterAppLastOkTime,
	SystemAppStatus,
	SystemAppLastOkTime,
	PlateNo,
	DriverName,
	DriverSurName,
	GpsAppVersion,
	ThermAppVersion,
	TransferAppVersion,
	CheckerAppVersion,    
	CamAppVersion,
	ReaderAppVersion,
	SystemAppVersion  
	 FROM public.rfid_devices
	 WHERE DeviceId=%f;`, res.DeviceId)
}

//InsertSQL
func (res RfidDeviceType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_devices 
	(SerialNumber,CustomerId) 
	  VALUES ('%s',%f)   
	  RETURNING DeviceId;`,
		res.SerialNumber, res.CustomerId)
}

//InsertDeviceDataSQL
func (res RfidDeviceType) InsertDeviceDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_device_data 
	   (DeviceId,
		CustomerId,
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
		GpsAppStatus,
		GpsAppLastOkTime,
		GpsConnStatus,
		GpsConnLastOkTime,
		GpsStatus,
		GpsLastOkTime,
		ThermAppStatus,
		ThermAppLastOkTime,
		TransferAppStatus,
		TransferAppLastOkTime,
		ContactStatus,
		ContactLastOkTime,
		Speed,
		UpdaterAppStatus,
		UpdaterAppLastOkTime,
		SystemAppStatus,
		SystemAppLastOkTime,
		PlateNo,
		DriverName,
		DriverSurName,
		GpsAppVersion,
		ThermAppVersion,
		TransferAppVersion,
		CheckerAppVersion,    
		CamAppVersion,
		ReaderAppVersion,
		SystemAppVersion,
		TruckType) 
	  VALUES (%f,%f
		,'%s','%s','%s','%s','%s'
	    ,%f,%f
		,'%s','%s','%s','%s','%s'
	    ,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
	    ,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s'
	    ,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s'
		,%f
		,'%s','%s',,'%s','%s'
		,%f
		,'%s','%s','%s','%s','%s'
	    ,'%s','%s','%s','%s','%s'
		,'%s','%s','%s','%s','%s') 
	  RETURNING DataId;`,
		res.DeviceId,
		res.CustomerId,
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
		res.UpdaterAppStatus,
		res.UpdaterAppLastOkTime,
		res.SystemAppStatus,
		res.SystemAppLastOkTime,
		res.PlateNo,
		res.DriverName,
		res.DriverSurName,
		res.GpsAppVersion,
		res.ThermAppVersion,
		res.TransferAppVersion,
		res.CheckerAppVersion,
		res.CamAppVersion,
		res.ReaderAppVersion,
		res.SystemAppVersion,
		res.TruckType)
}

//UpdateBasicSQL
func (res RfidDeviceType) UpdateBasicSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices 
				SET DeviceType='%s',PlateNo='%s',DriverName='%s',DriverSurName='%s',TruckType='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.DeviceType, res.PlateNo, res.DriverName,
		res.DriverSurName, res.TruckType, res.DeviceId)
}

//UpdateCustomerSQL
func (res RfidDeviceType) UpdateCustomerSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices 
				SET CustomerId=%f 
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.CustomerId, res.DeviceId)
}

//UpdateVersionSQL
func (res RfidDeviceType) UpdateVersionSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices 
				SET GpsAppVersion='%s',ThermAppVersion='%s',TransferAppVersion='%s' 
				,CheckerAppVersion='%s',CamAppVersion='%s',ReaderAppVersion='%s'
				,SystemAppVersion='%s'
	  			WHERE DeviceId=%f  
				RETURNING DeviceId;`,
		res.GpsAppVersion, res.ThermAppVersion, res.TransferAppVersion,
		res.CheckerAppVersion, res.CamAppVersion, res.ReaderAppVersion, res.SystemAppVersion, res.DeviceId)
}

//UpdateGpsSQL
func (res RfidDeviceType) UpdateGpsSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices
			   SET GpsTime='%s',Latitude=%f,Longitude=%f,Speed=%f
			   WHERE DeviceId=%f AND CustomerId=%f 
			   RETURNING DeviceId;`, res.GpsTime,
		res.Latitude, res.Longitude, res.Speed,
		res.DeviceId, res.CustomerId)
}

//UpdateStatuSQL
func (res RfidDeviceType) UpdateStatuSQL() string {
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

	return fmt.Sprintf(`UPDATE public.rfid_devices
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
func (res RfidDeviceType) UpdateThermSQL() string {
	return fmt.Sprintf(`UPDATE public.rfid_devices
		 SET Therm='%s',ThermTime='%s'
		 WHERE DeviceId=%f AND CustomerId=%f 
		 RETURNING DeviceId;`,
		res.Therm, res.ThermTime, res.DeviceId, res.CustomerId)
}

//SelectWithDb
func (res RfidDeviceType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(&res.CustomerId,
		&res.DeviceType,
		&res.TruckType,
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
		&res.ReaderAppStatus,
		&res.ReaderAppLastOkTime,
		&res.ReaderConnStatus,
		&res.ReaderConnLastOkTime,
		&res.ReaderStatus,
		&res.ReaderLastOkTime,
		&res.CamAppStatus,
		&res.CamAppLastOkTime,
		&res.CamConnStatus,
		&res.CamConnLastOkTime,
		&res.CamStatus,
		&res.CamLastOkTime,
		&res.GpsAppStatus,
		&res.GpsAppLastOkTime,
		&res.GpsConnStatus,
		&res.GpsConnLastOkTime,
		&res.GpsStatus,
		&res.GpsLastOkTime,
		&res.ThermAppStatus,
		&res.ThermAppLastOkTime,
		&res.TransferAppStatus,
		&res.TransferAppLastOkTime,
		&res.ContactStatus,
		&res.ContactLastOkTime,
		&res.Speed,
		&res.UpdaterAppStatus,
		&res.UpdaterAppLastOkTime,
		&res.SystemAppStatus,
		&res.SystemAppLastOkTime,
		&res.PlateNo,
		&res.DriverName,
		&res.DriverSurName,
		&res.GpsAppVersion,
		&res.ThermAppVersion,
		&res.TransferAppVersion,
		&res.CheckerAppVersion,
		&res.CamAppVersion,
		&res.ReaderAppVersion,
		&res.SystemAppVersion)
	return errDb
}
