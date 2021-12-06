package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceBaseType
type RecyDeviceBaseType struct {
	DeviceId    float64
	ContainerNo string
	DeviceType  string
	NewData     bool
}

//New
func (res *RecyDeviceBaseType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.DeviceType = RECY_DEVICE_TYPE_NONE
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceBaseType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_BASE, res.ToIdString())
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
func (res *RecyDeviceBaseType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_BASE, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_BASE

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

//ToId String
func (res *RecyDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceBaseType
func ByteToRecyDeviceBaseType(retByte []byte) RecyDeviceBaseType {
	var retVal RecyDeviceBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceBaseType
func StringToRecyDeviceBaseType(retStr string) RecyDeviceBaseType {
	return ByteToRecyDeviceBaseType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceBaseType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,DeviceType
	 FROM public.`+DATATYPE_RECY_BASE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_BASE+`  (DeviceId,ContainerNo,DeviceType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.ContainerNo, res.DeviceType)
}

//UpdateSQL
func (res *RecyDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_BASE+`  
	  SET ContainerNo='%s',DeviceType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.ContainerNo,
		res.DeviceType,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.DeviceType)
	return errDb
}

//CreateDb
func (res *RecyDeviceBaseType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_RECY_BASE + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	ContainerNo  varchar(50) NOT NULL DEFAULT '',
	DeviceType  varchar(50) NOT NULL DEFAULT '` + RECY_DEVICE_TYPE_NONE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}

// deleteSQL := fmt.Sprintf(`ALTER TABLE ult_sens_devices DROP COLUMN IF EXISTS `+WasteLibrary.DATATYPE_ULT_STATU+` ;`)
// _, err = currentDb.Exec(deleteSQL)
// WasteLibrary.LogErr(err)
