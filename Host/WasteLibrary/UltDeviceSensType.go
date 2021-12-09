package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UltDeviceSensType
type UltDeviceSensType struct {
	DeviceId   float64
	UltTime    string
	UltCount   float64
	UltRange1  float64
	UltRange2  float64
	UltRange3  float64
	UltRange4  float64
	UltRange5  float64
	UltRange6  float64
	UltRange7  float64
	UltRange8  float64
	UltRange9  float64
	UltRange10 float64
	UltRange11 float64
	UltRange12 float64
	UltRange13 float64
	UltRange14 float64
	UltRange15 float64
	UltRange16 float64
	UltRange17 float64
	UltRange18 float64
	UltRange19 float64
	UltRange20 float64
	UltRange21 float64
	UltRange22 float64
	UltRange23 float64
	UltRange24 float64

	NewData bool
}

//New
func (res *UltDeviceSensType) New() {
	res.DeviceId = 0
	res.UltTime = GetTime()
	res.UltCount = 0
	res.UltRange1 = 0
	res.UltRange2 = 0
	res.UltRange3 = 0
	res.UltRange4 = 0
	res.UltRange5 = 0
	res.UltRange6 = 0
	res.UltRange7 = 0
	res.UltRange8 = 0
	res.UltRange9 = 0
	res.UltRange10 = 0
	res.UltRange11 = 0
	res.UltRange12 = 0
	res.UltRange13 = 0
	res.UltRange14 = 0
	res.UltRange15 = 0
	res.UltRange16 = 0
	res.UltRange17 = 0
	res.UltRange18 = 0
	res.UltRange19 = 0
	res.UltRange20 = 0
	res.UltRange21 = 0
	res.UltRange22 = 0
	res.UltRange23 = 0
	res.UltRange24 = 0

	res.NewData = false
}

//GetByRedis
func (res *UltDeviceSensType) GetByRedis(dbIndex string) ResultType {
	resultVal := GetRedisForStoreApi(dbIndex, REDIS_ULT_SENS, res.ToIdString())
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
func (res *UltDeviceSensType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_ULT_SENS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UltDeviceSensType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_SENS

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
func (res *UltDeviceSensType) SaveToReaderDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_ULT_SENS

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
func (res *UltDeviceSensType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceSensType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UltDeviceSensType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *UltDeviceSensType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UltDeviceSensType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UltDeviceSensType) SelectSQL() string {
	return fmt.Sprintf(`SELECT UltTime,
	 UltCount,
	 UltRange1,
     UltRange2,
     UltRange3,
     UltRange4,
     UltRange5,
     UltRange6,
     UltRange7,
     UltRange8,
     UltRange9,
     UltRange10,
     UltRange11,
     UltRange12,
     UltRange13,
     UltRange14,
     UltRange15,
     UltRange16,
     UltRange17,
     UltRange18,
     UltRange19,
     UltRange20,
     UltRange21,
     UltRange22,
     UltRange23,
     UltRange24
	 FROM public.`+DATATYPE_ULT_SENS+` 
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *UltDeviceSensType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_ULT_SENS+`  (DeviceId,UltTime,
		UltCount,
		UltRange1,
        UltRange2,
        UltRange3,
        UltRange4,
        UltRange5,
        UltRange6,
        UltRange7,
        UltRange8,
        UltRange9,
        UltRange10,
        UltRange11,
        UltRange12,
        UltRange13,
        UltRange14,
        UltRange15,
        UltRange16,
        UltRange17,
        UltRange18,
        UltRange19,
        UltRange20,
        UltRange21,
        UltRange22,
        UltRange23,
        UltRange24) 
	  VALUES (%f,'%s',%f,%f,%f
	  ,%f,%f,%f,%f,%f
	  ,%f,%f,%f,%f,%f
	  ,%f,%f,%f,%f,%f
	  ,%f,%f,%f,%f,%f
	  ,%f,%f) 
	  RETURNING DeviceId;`, res.DeviceId,
		res.UltTime,
		res.UltCount,
		res.UltRange1,
		res.UltRange2,
		res.UltRange3,
		res.UltRange4,
		res.UltRange5,
		res.UltRange6,
		res.UltRange7,
		res.UltRange8,
		res.UltRange9,
		res.UltRange10,
		res.UltRange11,
		res.UltRange12,
		res.UltRange13,
		res.UltRange14,
		res.UltRange15,
		res.UltRange16,
		res.UltRange17,
		res.UltRange18,
		res.UltRange19,
		res.UltRange20,
		res.UltRange21,
		res.UltRange22,
		res.UltRange23,
		res.UltRange24)
}

//UpdateSQL
func (res *UltDeviceSensType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_ULT_SENS+`  
	  SET UltTime='%s',
	  UltCount=%f,
	  UltRange1=%f,
      UltRange2=%f,
      UltRange3=%f,
      UltRange4=%f,
      UltRange5=%f,
      UltRange6=%f,
      UltRange7=%f,
      UltRange8=%f,
      UltRange9=%f,
      UltRange10=%f,
      UltRange11=%f,
      UltRange12=%f,
      UltRange13=%f,
      UltRange14=%f,
      UltRange15=%f,
      UltRange16=%f,
      UltRange17=%f,
      UltRange18=%f,
      UltRange19=%f,
      UltRange20=%f,
      UltRange21=%f,
      UltRange22=%f,
      UltRange23=%f,
      UltRange24=%f
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.UltTime,
		res.UltCount,
		res.UltRange1,
		res.UltRange2,
		res.UltRange3,
		res.UltRange4,
		res.UltRange5,
		res.UltRange6,
		res.UltRange7,
		res.UltRange8,
		res.UltRange9,
		res.UltRange10,
		res.UltRange11,
		res.UltRange12,
		res.UltRange13,
		res.UltRange14,
		res.UltRange15,
		res.UltRange16,
		res.UltRange17,
		res.UltRange18,
		res.UltRange19,
		res.UltRange20,
		res.UltRange21,
		res.UltRange22,
		res.UltRange23,
		res.UltRange24,
		res.DeviceId)
}

//SelectWithDb
func (res *UltDeviceSensType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.UltTime,
		&res.UltCount,
		&res.UltRange1,
		&res.UltRange2,
		&res.UltRange3,
		&res.UltRange4,
		&res.UltRange5,
		&res.UltRange6,
		&res.UltRange7,
		&res.UltRange8,
		&res.UltRange9,
		&res.UltRange10,
		&res.UltRange11,
		&res.UltRange12,
		&res.UltRange13,
		&res.UltRange14,
		&res.UltRange15,
		&res.UltRange16,
		&res.UltRange17,
		&res.UltRange18,
		&res.UltRange19,
		&res.UltRange20,
		&res.UltRange21,
		&res.UltRange22,
		&res.UltRange23,
		&res.UltRange24)
	return errDb
}

//CreateDb
func (res *UltDeviceSensType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_ULT_SENS + `  (
	DataId serial PRIMARY KEY,
	DeviceId INT NOT NULL DEFAULT -1,
	UltTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	UltCount INT NOT NULL DEFAULT 0,
	UltRange1  INT NOT NULL DEFAULT 0,
	UltRange2  INT NOT NULL DEFAULT 0,
	UltRange3  INT NOT NULL DEFAULT 0,
	UltRange4  INT NOT NULL DEFAULT 0,
	UltRange5  INT NOT NULL DEFAULT 0,
	UltRange6  INT NOT NULL DEFAULT 0,
	UltRange7  INT NOT NULL DEFAULT 0,
	UltRange8  INT NOT NULL DEFAULT 0,
	UltRange9  INT NOT NULL DEFAULT 0,
	UltRange10 INT NOT NULL DEFAULT 0,
	UltRange11 INT NOT NULL DEFAULT 0,
	UltRange12 INT NOT NULL DEFAULT 0,
	UltRange13 INT NOT NULL DEFAULT 0,
	UltRange14 INT NOT NULL DEFAULT 0,
	UltRange15 INT NOT NULL DEFAULT 0,
	UltRange16 INT NOT NULL DEFAULT 0,
	UltRange17 INT NOT NULL DEFAULT 0,
	UltRange18 INT NOT NULL DEFAULT 0,
	UltRange19 INT NOT NULL DEFAULT 0,
	UltRange20 INT NOT NULL DEFAULT 0,
	UltRange21 INT NOT NULL DEFAULT 0,
	UltRange22 INT NOT NULL DEFAULT 0,
	UltRange23 INT NOT NULL DEFAULT 0,
	UltRange24 INT NOT NULL DEFAULT 0,
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
