package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//UserType
type UserType struct {
	UserId     float64
	CustomerId float64
	UserName   string
	UserRole   string
	Password   string
	Email      string
	Active     string
	CreateTime string
}

//New
func (res *UserType) New() {
	res.UserId = 0
	res.CustomerId = 1
	res.UserName = ""
	res.UserRole = USER_ROLE_GUEST
	res.Password = ""
	res.Email = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *UserType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_USERS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *UserType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_USERS, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *UserType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_USER

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveConfigDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.UserId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *UserType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.UserId)
}

//ToCustomerId String
func (res *UserType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *UserType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *UserType) ToString() string {
	return string(res.ToByte())

}

//Byte To UserType
func ByteToUserType(retByte []byte) UserType {
	var retVal UserType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UserType
func StringToUserType(retStr string) UserType {
	return ByteToUserType([]byte(retStr))
}

//ByteToType
func (res *UserType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *UserType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *UserType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
			CustomerId,
			UserName,
			UserRole,
			Password,
			Email,
			Active,
			CreateTime 
			 FROM public.`+DATATYPE_USER+`  
			 WHERE UserId=%f  ;`, res.UserId)
}

//InsertSQL
func (res *UserType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.`+DATATYPE_USER+`  
	(UserRole,Email,UserName,CustomerId,Password) 
	  VALUES ('%s','%s','%s',%f,'%s')   
	  RETURNING UserId;`,
		res.UserRole, res.Email, res.UserName,
		res.CustomerId, res.Password)
}

//UpdateSQL
func (res *UserType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_USER+`  
	SET UserRole='%s',Email='%s',UserName='%s'
	  WHERE UserId=%f  
	RETURNING UserId;`,
		res.UserRole, res.Email, res.UserName, res.UserId)
}

//UpdatePasswordSQL
func (res *UserType) UpdatePasswordSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_USER+`  
	SET Password='%s'
	  WHERE UserId=%f  
	RETURNING UserId;`,
		res.Password, res.UserId)
}

//SelectWithDb
func (res *UserType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(&res.CustomerId,
		&res.UserName,
		&res.UserRole,
		&res.Password,
		&res.Email,
		&res.Active,
		&res.CreateTime)
	return errDb
}
