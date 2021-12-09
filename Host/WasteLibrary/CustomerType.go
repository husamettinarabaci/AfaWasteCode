package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//CustomerType
type CustomerType struct {
	CustomerId   float64
	CustomerName string
	CustomerLink string
	RfIdApp      string
	UltApp       string
	RecyApp      string
	Active       string
	CreateTime   string
}

//New
func (res *CustomerType) New() {
	res.CustomerId = 1
	res.CustomerName = ""
	res.CustomerLink = ""
	res.RfIdApp = STATU_PASSIVE
	res.UltApp = STATU_PASSIVE
	res.RecyApp = STATU_PASSIVE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//
//GetByRedis
func (res *CustomerType) GetByRedis() ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_CUSTOMERS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByLink
func (res *CustomerType) GetByRedisByLink(link string) ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_CUSTOMER_LINK, link)
	if resultVal.Result == RESULT_OK {
		var customerId string = resultVal.Retval.(string)
		res.CustomerId = StringIdToFloat64(customerId)
		resultVal = res.GetByRedis()
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *CustomerType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMERS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisLink
func (res *CustomerType) SaveToRedisLink() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_LINK, res.CustomerLink, res.ToIdString())
	return resultVal
}

//SaveToDb
func (res *CustomerType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_CUSTOMER

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveConfigDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.CustomerId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *CustomerType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *CustomerType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *CustomerType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *CustomerType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomerType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *CustomerType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
	CustomerName,
	CustomerLink,
	RfIdApp,
	UltApp,
	RecyApp,
	Active,
	CreateTime 
	FROM public.`+DATATYPE_CUSTOMER+`  
	WHERE CustomerId=%f  ;`, res.CustomerId)
}

//InsertSQL
func (res *CustomerType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_CUSTOMER+` (
		CustomerName,CustomerLink,RfIdApp,UltApp,RecyApp)
  VALUES ('%s','%s','%s','%s','%s')  RETURNING CustomerId;`,
		res.CustomerName, res.CustomerLink, res.RfIdApp,
		res.UltApp, res.RecyApp)
}

//UpdateSQL
func (res *CustomerType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_CUSTOMER+` 
		 SET CustomerName='%s',CustomerLink='%s',RfIdApp='%s',UltApp='%s',RecyApp='%s'
		 WHERE CustomerId=%f  RETURNING CustomerId;`,
		res.CustomerName, res.CustomerLink, res.RfIdApp,
		res.UltApp, res.RecyApp, res.CustomerId)
}

//SelectWithDb
func (res *CustomerType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerName,
		&res.CustomerLink,
		&res.RfIdApp,
		&res.UltApp,
		&res.RecyApp,
		&res.Active,
		&res.CreateTime)
	return errDb
}

//CreateDb
func (res *CustomerType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_CUSTOMER + ` ( 
	CustomerId serial PRIMARY KEY,
	CustomerName varchar(50) NOT NULL DEFAULT '',
	CustomerLink varchar(50) NOT NULL DEFAULT '',
	RfIdApp varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	UltApp varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	RecyApp varchar(50) NOT NULL DEFAULT '` + STATU_PASSIVE + `',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
