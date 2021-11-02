package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RfidDeviceStatuType
type RfidDeviceStatuType struct {
	DeviceId              float64
	StatusTime            string
	AliveStatus           string
	AliveLastOkTime       string
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
	NewData               bool
}

//New
func (res *RfidDeviceStatuType) New() {
	res.DeviceId = 0
	res.StatusTime = ""
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = ""
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
	res.NewData = false
}

//ToId String
func (res RfidDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RfidDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RfidDeviceStatuType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceStatuType
func ByteToRfidDeviceStatuType(retByte []byte) RfidDeviceStatuType {
	var retVal RfidDeviceStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceStatuType
func StringToRfidDeviceStatuType(retStr string) RfidDeviceStatuType {
	return ByteToRfidDeviceStatuType([]byte(retStr))
}

//SelectSQL
func (res RfidDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime,ReaderAppStatus,ReaderAppLastOkTime,ReaderConnStatus,
	ReaderConnLastOkTime,ReaderStatus,ReaderLastOkTime,CamAppStatus,CamAppLastOkTime,CamConnStatus,CamConnLastOkTime,
	CamStatus,CamLastOkTime,ThermAppStatus,ThermAppLastOkTime,TransferAppStatus,TransferAppLastOkTime,SystemAppStatus,
	SystemAppLastOkTime,UpdaterAppStatus,UpdaterAppLastOkTime,GpsAppStatus,GpsAppLastOkTime,GpsConnStatus,
	GpsConnLastOkTime,GpsStatus,GpsLastOkTime
	 FROM public.rfid_statu_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RfidDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.rfid_statu_devices (DeviceId,
	StatusTime,AliveStatus,AliveLastOkTime,ReaderAppStatus,ReaderAppLastOkTime,ReaderConnStatus,
	ReaderConnLastOkTime,ReaderStatus,ReaderLastOkTime,CamAppStatus,CamAppLastOkTime,CamConnStatus,CamConnLastOkTime,
	CamStatus,CamLastOkTime,ThermAppStatus,ThermAppLastOkTime,TransferAppStatus,TransferAppLastOkTime,SystemAppStatus,
	SystemAppLastOkTime,UpdaterAppStatus,UpdaterAppLastOkTime,GpsAppStatus,GpsAppLastOkTime,GpsConnStatus,
	GpsConnLastOkTime,GpsStatus,GpsLastOkTime) 
	  VALUES (%f
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId,
		res.StatusTime, res.AliveStatus, res.AliveLastOkTime, res.ReaderAppStatus, res.ReaderAppLastOkTime, res.ReaderConnStatus,
		res.ReaderConnLastOkTime, res.ReaderStatus, res.ReaderLastOkTime, res.CamAppStatus, res.CamAppLastOkTime, res.CamConnStatus, res.CamConnLastOkTime,
		res.CamStatus, res.CamLastOkTime, res.ThermAppStatus, res.ThermAppLastOkTime, res.TransferAppStatus, res.TransferAppLastOkTime, res.SystemAppStatus,
		res.SystemAppLastOkTime, res.UpdaterAppStatus, res.UpdaterAppLastOkTime, res.GpsAppStatus, res.GpsAppLastOkTime, res.GpsConnStatus,
		res.GpsConnLastOkTime, res.GpsStatus, res.GpsLastOkTime)
}

//UpdateSQL
func (res RfidDeviceStatuType) UpdateSQL() string {
	var execSqlExt = ""
	if res.AliveStatus == STATU_ACTIVE {
		execSqlExt += ",AliveLastOkTime='" + res.AliveLastOkTime + "'"
	}
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
	if res.GpsAppStatus == STATU_ACTIVE {
		execSqlExt += ",GpsAppLastOkTime='" + res.GpsAppLastOkTime + "'"
	}
	if res.GpsConnStatus == STATU_ACTIVE {
		execSqlExt += ",GpsConnLastOkTime='" + res.GpsConnLastOkTime + "'"
	}
	if res.GpsStatus == STATU_ACTIVE {
		execSqlExt += ",GpsLastOkTime='" + res.GpsLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.rfid_statu_devices 
	  SET StatusTime='%s',AliveStatus='%s',ReaderAppStatus='%s',ReaderConnStatus='%s',
	  ReaderStatus='%s',CamAppStatus='%s',CamConnStatus='%s',
	  CamStatus='%s',ThermAppStatus='%s',TransferAppStatus='%s',SystemAppStatus='%s',
	  UpdaterAppStatus='%s',GpsAppStatus='%s',GpsConnStatus='%s',
	  GpsStatus='%s'`+execSqlExt+`
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.StatusTime, res.AliveStatus, res.ReaderAppStatus, res.ReaderConnStatus,
		res.ReaderStatus, res.CamAppStatus, res.CamConnStatus,
		res.CamStatus, res.ThermAppStatus, res.TransferAppStatus, res.SystemAppStatus,
		res.UpdaterAppStatus, res.GpsAppStatus, res.GpsConnStatus,
		res.GpsStatus, res.DeviceId)
}

//SelectWithDb
func (res RfidDeviceStatuType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.StatusTime,
		&res.AliveStatus,
		&res.AliveLastOkTime,
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
		&res.ThermAppStatus,
		&res.ThermAppLastOkTime,
		&res.TransferAppStatus,
		&res.TransferAppLastOkTime,
		&res.SystemAppStatus,
		&res.SystemAppLastOkTime,
		&res.UpdaterAppStatus,
		&res.UpdaterAppLastOkTime,
		&res.GpsAppStatus,
		&res.GpsAppLastOkTime,
		&res.GpsConnStatus,
		&res.GpsConnLastOkTime,
		&res.GpsStatus,
		&res.GpsLastOkTime)
	return errDb
}
