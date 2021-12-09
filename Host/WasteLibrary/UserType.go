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
	FirstName  string
	LastName   string
	UserRole   string
	Password   string
	Email      string
	Token      string
	Active     string
	CreateTime string
}

//New
func (res *UserType) New() {
	res.UserId = 0
	res.CustomerId = 1
	res.FirstName = ""
	res.LastName = ""
	res.UserRole = USER_ROLE_GUEST
	res.Password = ""
	res.Email = ""
	res.Token = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *UserType) GetByRedis() ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_USERS, res.ToIdString())
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
	resultVal := SaveRedisForStoreApi(REDIS_USERS, res.ToIdString(), res.ToString())
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

//ByteToType
func (res *UserType) ByteToType(retByte []byte) {
	res.New()
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
			FirstName,
			LastName,
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
	(UserRole,Email,FirstName,LastName,CustomerId,Password) 
	  VALUES ('%s','%s','%s','%s',%f,'%s')   
	  RETURNING UserId;`,
		res.UserRole, res.Email, res.FirstName, res.LastName,
		res.CustomerId, res.Password)
}

//UpdateSQL
func (res *UserType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.`+DATATYPE_USER+`  
	SET UserRole='%s',Email='%s',FirstName='%s',LastName='%s'
	  WHERE UserId=%f  
	RETURNING UserId;`,
		res.UserRole, res.Email, res.FirstName, res.LastName, res.UserId)
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
		&res.FirstName,
		&res.LastName,
		&res.UserRole,
		&res.Password,
		&res.Email,
		&res.Active,
		&res.CreateTime)
	return errDb
}

//CreateDb
func (res *UserType) CreateDb(currentDb *sql.DB) {
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + DATATYPE_USER + `  ( 
	UserId serial PRIMARY KEY,
	CustomerId INT NOT NULL DEFAULT -1,
	FirstName varchar(50) NOT NULL DEFAULT '',
	LastName varchar(50) NOT NULL DEFAULT '',
	UserRole varchar(50) NOT NULL DEFAULT '` + USER_ROLE_GUEST + `',
	Password varchar(50) NOT NULL DEFAULT '',
	Email varchar(50) NOT NULL DEFAULT '',
	Active varchar(50) NOT NULL DEFAULT '` + STATU_ACTIVE + `',
	CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`)
	_, err := currentDb.Exec(createSQL)
	LogErr(err)
}
