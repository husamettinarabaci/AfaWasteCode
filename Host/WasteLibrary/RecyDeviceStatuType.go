package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceStatuType
type RecyDeviceStatuType struct {
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
	ThermAppStatus        string
	ThermAppLastOkTime    string
	TransferAppStatus     string
	TransferAppLastOkTime string
	SystemAppStatus       string
	SystemAppLastOkTime   string
	UpdaterAppStatus      string
	UpdaterAppLastOkTime  string
	MotorAppStatus        string
	MotorAppLastOkTime    string
	MotorConnStatus       string
	MotorConnLastOkTime   string
	MotorStatus           string
	MotorLastOkTime       string
	WebAppStatus          string
	WebAppLastOkTime      string
	NewData               bool
}

//New
func (res *RecyDeviceStatuType) New() {
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
	res.ThermAppStatus = STATU_PASSIVE
	res.ThermAppLastOkTime = GetTime()
	res.TransferAppStatus = STATU_PASSIVE
	res.TransferAppLastOkTime = GetTime()
	res.SystemAppStatus = STATU_PASSIVE
	res.SystemAppLastOkTime = GetTime()
	res.UpdaterAppStatus = STATU_PASSIVE
	res.UpdaterAppLastOkTime = GetTime()
	res.MotorAppStatus = STATU_PASSIVE
	res.MotorAppLastOkTime = GetTime()
	res.MotorConnStatus = STATU_PASSIVE
	res.MotorConnLastOkTime = GetTime()
	res.MotorStatus = STATU_PASSIVE
	res.MotorLastOkTime = GetTime()
	res.WebAppStatus = STATU_PASSIVE
	res.WebAppLastOkTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceStatuType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_STATU_DEVICES, res.ToIdString())
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
func (res *RecyDeviceStatuType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_STATU_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceStatuType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_STATU_DEVICE

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
func (res *RecyDeviceStatuType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_STATU_DEVICE

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
func (res *RecyDeviceStatuType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceStatuType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceStatuType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceStatuType
func ByteToRecyDeviceStatuType(retByte []byte) RecyDeviceStatuType {
	var retVal RecyDeviceStatuType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceStatuType
func StringToRecyDeviceStatuType(retStr string) RecyDeviceStatuType {
	return ByteToRecyDeviceStatuType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceStatuType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceStatuType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceStatuType) SelectSQL() string {
	return fmt.Sprintf(`SELECT StatusTime,AliveStatus,AliveLastOkTime,ReaderAppStatus,ReaderAppLastOkTime,ReaderConnStatus,
	ReaderConnLastOkTime,ReaderStatus,ReaderLastOkTime,CamAppStatus,CamAppLastOkTime,CamConnStatus,CamConnLastOkTime,
	CamStatus,CamLastOkTime,ThermAppStatus,ThermAppLastOkTime,TransferAppStatus,TransferAppLastOkTime,SystemAppStatus,
	SystemAppLastOkTime,UpdaterAppStatus,UpdaterAppLastOkTime,MotorAppStatus,MotorAppLastOkTime,MotorConnStatus,
	MotorConnLastOkTime,MotorStatus,MotorLastOkTime,WebAppStatus,WebAppLastOkTime
	 FROM public.recy_statu_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceStatuType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_statu_devices (DeviceId,
	StatusTime,AliveStatus,AliveLastOkTime,ReaderAppStatus,ReaderAppLastOkTime,ReaderConnStatus,
	ReaderConnLastOkTime,ReaderStatus,ReaderLastOkTime,CamAppStatus,CamAppLastOkTime,CamConnStatus,CamConnLastOkTime,
	CamStatus,CamLastOkTime,ThermAppStatus,ThermAppLastOkTime,TransferAppStatus,TransferAppLastOkTime,SystemAppStatus,
	SystemAppLastOkTime,UpdaterAppStatus,UpdaterAppLastOkTime,MotorAppStatus,MotorAppLastOkTime,MotorConnStatus,
	MotorConnLastOkTime,MotorStatus,MotorLastOkTime,WebAppStatus,WebAppLastOkTime) 
	  VALUES (%f
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'
	,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s'
	,'%s') 
	  RETURNING DeviceId;`, res.DeviceId,
		res.StatusTime, res.AliveStatus, res.AliveLastOkTime, res.ReaderAppStatus, res.ReaderAppLastOkTime, res.ReaderConnStatus,
		res.ReaderConnLastOkTime, res.ReaderStatus, res.ReaderLastOkTime, res.CamAppStatus, res.CamAppLastOkTime, res.CamConnStatus, res.CamConnLastOkTime,
		res.CamStatus, res.CamLastOkTime, res.ThermAppStatus, res.ThermAppLastOkTime, res.TransferAppStatus, res.TransferAppLastOkTime, res.SystemAppStatus,
		res.SystemAppLastOkTime, res.UpdaterAppStatus, res.UpdaterAppLastOkTime, res.MotorAppStatus, res.MotorAppLastOkTime, res.MotorConnStatus,
		res.MotorConnLastOkTime, res.MotorStatus, res.MotorLastOkTime, res.WebAppStatus, res.WebAppLastOkTime)
}

//UpdateSQL
func (res *RecyDeviceStatuType) UpdateSQL() string {
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
	if res.MotorAppStatus == STATU_ACTIVE {
		execSqlExt += ",MotorAppLastOkTime='" + res.MotorAppLastOkTime + "'"
	}
	if res.MotorConnStatus == STATU_ACTIVE {
		execSqlExt += ",MotorConnLastOkTime='" + res.MotorConnLastOkTime + "'"
	}
	if res.MotorStatus == STATU_ACTIVE {
		execSqlExt += ",MotorLastOkTime='" + res.MotorLastOkTime + "'"
	}
	if res.WebAppStatus == STATU_ACTIVE {
		execSqlExt += ",WebAppLastOkTime='" + res.WebAppLastOkTime + "'"
	}

	return fmt.Sprintf(`UPDATE public.recy_statu_devices 
	  SET StatusTime='%s',AliveStatus='%s',ReaderAppStatus='%s',ReaderConnStatus='%s',
	  ReaderStatus='%s',CamAppStatus='%s',CamConnStatus='%s',
	  CamStatus='%s',ThermAppStatus='%s',TransferAppStatus='%s',SystemAppStatus='%s',
	  UpdaterAppStatus='%s',MotorAppStatus='%s',MotorConnStatus='%s',
	  MotorStatus='%s',WebAppStatus='%s'`+execSqlExt+`
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.StatusTime, res.AliveStatus, res.ReaderAppStatus, res.ReaderConnStatus,
		res.ReaderStatus, res.CamAppStatus, res.CamConnStatus,
		res.CamStatus, res.ThermAppStatus, res.TransferAppStatus, res.SystemAppStatus,
		res.UpdaterAppStatus, res.MotorAppStatus, res.MotorConnStatus,
		res.MotorStatus, res.WebAppStatus, res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceStatuType) SelectWithDb(db *sql.DB) error {
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
		&res.MotorAppStatus,
		&res.MotorAppLastOkTime,
		&res.MotorConnStatus,
		&res.MotorConnLastOkTime,
		&res.MotorStatus,
		&res.MotorLastOkTime,
		&res.WebAppStatus,
		&res.WebAppLastOkTime)
	return errDb
}
