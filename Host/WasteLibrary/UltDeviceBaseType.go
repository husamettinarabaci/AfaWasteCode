package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceBaseType
type UltDeviceBaseType struct {
	DeviceId      float64
	ContainerNo   string
	ContainerType string
	DeviceType    string
	Imei          string
	Imsi          string
	NewData       bool
}

//New
func (res *UltDeviceBaseType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.ContainerType = CONTAINERTYPE_NONE
	res.DeviceType = ULT_DEVICE_TYPE_NONE
	res.Imei = ""
	res.Imsi = ""
	res.NewData = false
}

//GetByRedis
func (res *UltDeviceBaseType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_ULT_BASE_DEVICES, res.ToIdString())
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
func (res *UltDeviceBaseType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_ULT_BASE_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceBaseType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_BASE_DEVICE

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
func (res *UltDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceBaseType
func ByteToUltDeviceBaseType(retByte []byte) UltDeviceBaseType {
	var retVal UltDeviceBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceBaseType
func StringToUltDeviceBaseType(retStr string) UltDeviceBaseType {
	return ByteToUltDeviceBaseType([]byte(retStr))
}

//ByteToType
func (res *UltDeviceBaseType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceBaseType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,ContainerType,DeviceType,Imei,Imsi
	 FROM public.`+DATATYPE_ULT_BASE_DEVICE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_BASE_DEVICE+`  (DeviceId,ContainerNo,ContainerType,DeviceType,Imei,Imsi) 
	  VALUES (%f,'%s','%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.ContainerNo, res.ContainerType, res.DeviceType, res.Imei, res.Imsi)
}

//UpdateSQL
func (res *UltDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_BASE_DEVICE+`  
	  SET ContainerNo='%s',ContainerType='%s',DeviceType='%s',Imei='%s',Imsi='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.ContainerNo,
		res.ContainerType,
		res.DeviceType,
		res.Imei,
		res.Imsi,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.ContainerType,
		&res.DeviceType,
		&res.Imei,
		&res.Imsi)
	return errDb
}
