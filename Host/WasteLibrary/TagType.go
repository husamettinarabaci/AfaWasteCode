package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagType
type TagType struct {
	TagId     float64
	TagMain   TagMainType
	TagBase   TagBaseType
	TagStatu  TagStatuType
	TagGps    TagGpsType
	TagReader TagReaderType
}

//New
func (res *TagType) New() {
	res.TagId = 0
	res.TagMain.New()
	res.TagBase.New()
	res.TagStatu.New()
	res.TagGps.New()
	res.TagReader.New()
}

//GetAll
func (res *TagType) GetAll() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_MAINS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.TagMain = StringToTagMainType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_TAG_BASES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.TagBase = StringToTagBaseType(resultVal.Retval.(string))
		res.TagBase.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_TAG_GPSES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.TagGps = StringToTagGpsType(resultVal.Retval.(string))
		res.TagGps.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_TAG_STATUS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.TagStatu = StringToTagStatuType(resultVal.Retval.(string))
		res.TagStatu.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_TAG_READERS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.TagReader = StringToTagReaderType(resultVal.Retval.(string))
		res.TagReader.NewData = false
	} else {
		return resultVal
	}
	return resultVal
}

//ToId String
func (res *TagType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToTagId String
func (res *TagType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagType
func ByteToTagType(retByte []byte) TagType {
	var retVal TagType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagType
func StringToTagType(retStr string) TagType {
	return ByteToTagType([]byte(retStr))
}
