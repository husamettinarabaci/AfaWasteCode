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
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMERS, res.ToIdString())
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
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMER_LINK, link)
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
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMERS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToRedisLink
func (res *CustomerType) SaveToRedisLink() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_CUSTOMER_LINK, res.CustomerLink, res.ToIdString())
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
	resultVal = SaveStaticDbMainForStoreApi(data)
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

//Byte To CustomerType
func ByteToCustomerType(retByte []byte) CustomerType {
	var retVal CustomerType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomerType
func StringToCustomerType(retStr string) CustomerType {
	return ByteToCustomerType([]byte(retStr))
}

//ByteToType
func (res *CustomerType) ByteToType(retByte []byte) {
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
	FROM public.customers 
	WHERE CustomerId=%f  ;`, res.CustomerId)
}

//InsertSQL
func (res *CustomerType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.customers(
		CustomerName,CustomerLink,RfIdApp,UltApp,RecyApp)
  VALUES ('%s','%s','%s','%s','%s')  RETURNING CustomerId;`,
		res.CustomerName, res.CustomerLink, res.RfIdApp,
		res.UltApp, res.RecyApp)
}

//UpdateSQL
func (res *CustomerType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.customers
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
