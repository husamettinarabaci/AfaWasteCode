package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceDetailType
type RecyDeviceDetailType struct {
	DeviceId          float64
	TotalGlassCount   float64
	TotalPlasticCount float64
	TotalMetalCount   float64
	DailyGlassCount   float64
	DailyPlasticCount float64
	DailyMetalCount   float64
	RecyTime          string
	NewData           bool
}

//New
func (res *RecyDeviceDetailType) New() {
	res.DeviceId = 0
	res.TotalGlassCount = 0
	res.TotalPlasticCount = 0
	res.TotalMetalCount = 0
	res.DailyGlassCount = 0
	res.DailyPlasticCount = 0
	res.DailyMetalCount = 0
	res.RecyTime = GetTime()
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceDetailType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RECY_DETAIL_DEVICES, res.ToIdString())
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
func (res *RecyDeviceDetailType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_DETAIL_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceDetailType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_DETAIL_DEVICE

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
func (res *RecyDeviceDetailType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceDetailType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceDetailType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceDetailType
func ByteToRecyDeviceDetailType(retByte []byte) RecyDeviceDetailType {
	var retVal RecyDeviceDetailType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceDetailType
func StringToRecyDeviceDetailType(retStr string) RecyDeviceDetailType {
	return ByteToRecyDeviceDetailType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceDetailType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceDetailType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceDetailType) SelectSQL() string {

	return fmt.Sprintf(`SELECT TotalGlassCount,TotalPlasticCount,TotalMetalCount,DailyGlassCount,DailyPlasticCount,DailyMetalCount,RecyTime
	 FROM public.`+DATATYPE_RECY_DETAIL_DEVICE+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceDetailType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_RECY_DETAIL_DEVICE+`  (DeviceId,TotalGlassCount,TotalPlasticCount,TotalMetalCount,DailyGlassCount,DailyPlasticCount,DailyMetalCount,RecyTime) 
	  VALUES (%f,%f,%f,%f,%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.TotalGlassCount, res.TotalPlasticCount,
		res.TotalMetalCount, res.DailyGlassCount, res.DailyPlasticCount, res.DailyMetalCount, res.RecyTime)
}

//UpdateSQL
func (res *RecyDeviceDetailType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_RECY_DETAIL_DEVICE+`  
	  SET TotalGlassCount=%f,TotalPlasticCount=%f,TotalMetalCount=%f,DailyGlassCount=%f,DailyPlasticCount=%f,DailyMetalCount=%f,RecyTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.TotalGlassCount,
		res.TotalPlasticCount,
		res.TotalMetalCount,
		res.DailyGlassCount,
		res.DailyPlasticCount,
		res.DailyMetalCount,
		res.RecyTime,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceDetailType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.TotalGlassCount,
		&res.TotalPlasticCount,
		&res.TotalMetalCount,
		&res.DailyGlassCount,
		&res.DailyPlasticCount,
		&res.DailyMetalCount,
		&res.RecyTime)
	return errDb
}
