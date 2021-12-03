package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
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
	res.StatusTime = GetTime()
	res.AliveStatus = STATU_PASSIVE
	res.AliveLastOkTime = GetTime()
	res.ReaderAppStatus = STATU_PASSIVE
	res.ReaderAppLastOkTime = GetTime()
	res.ReaderConnStatus = STATU_PASSIVE
	res.ReaderConnLastOkTime = GetTime()
	res.ReaderStatus = STATU_PASSIVE
	res.ReaderLastOkTime = GetTime()
	res.CamAppStatus = STATU_PASSIVE
	res.CamAppLastOkTime = GetTime()
	res.CamConnStatus = STATU_PASSIVE
	res.CamConnLastOkTime = GetTime()
	res.CamStatus = STATU_PASSIVE
	res.CamLastOkTime = GetTime()
	res.GpsAppStatus = STATU_PASSIVE
	res.GpsAppLastOkTime = GetTime()
	res.GpsConnStatus = STATU_PASSIVE
	res.GpsConnLastOkTime = GetTime()
	res.GpsStatus = STATU_PASSIVE
	res.GpsLastOkTime = GetTime()
	res.ThermAppStatus = STATU_PASSIVE
	res.ThermAppLastOkTime = GetTime()
	res.TransferAppStatus = STATU_PASSIVE
	res.TransferAppLastOkTime = GetTime()
	res.SystemAppStatus = STATU_PASSIVE
	res.SystemAppLastOkTime = GetTime()
	res.UpdaterAppStatus = STATU_PASSIVE
	res.UpdaterAppLastOkTime = GetTime()
	res.ContactStatus = STATU_PASSIVE
	res.ContactLastOkTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RfidDeviceStatuType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_STATU_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceStatuType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_STATU_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RfidDeviceStatuType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_STATU_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//SaveToReaderDb
func (res *RfidDeviceStatuType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RFID_STATU_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveReaderDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *RfidDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceStatuType) ToString() string {
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

//ByteToType
func (res *RfidDeviceStatuType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RfidDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime,ReaderAppStatus,ReaderAppLastOkTime,ReaderConnStatus,
	ReaderConnLastOkTime,ReaderStatus,ReaderLastOkTime,CamAppStatus,CamAppLastOkTime,CamConnStatus,CamConnLastOkTime,
	CamStatus,CamLastOkTime,ThermAppStatus,ThermAppLastOkTime,TransferAppStatus,TransferAppLastOkTime,SystemAppStatus,
	SystemAppLastOkTime,UpdaterAppStatus,UpdaterAppLastOkTime,GpsAppStatus,GpsAppLastOkTime,GpsConnStatus,
	GpsConnLastOkTime,GpsStatus,GpsLastOkTime
	 FROM public.`+DATATYPE_RFID_STATU_DEVICE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RfidDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RFID_STATU_DEVICE+`  (DeviceId,
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
func (res *RfidDeviceStatuType) UpdateSQL() string {
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

	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RFID_STATU_DEVICE+`  
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
func (res *RfidDeviceStatuType) SelectWithDb(db *sql.DB) error {
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
