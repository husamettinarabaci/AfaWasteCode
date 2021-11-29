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
	TagNote   TagNoteType
	TagAlarm  TagAlarmType
}

//New
func (res *TagType) New() {
	res.TagId = 0
	res.TagMain.New()
	res.TagBase.New()
	res.TagStatu.New()
	res.TagGps.New()
	res.TagReader.New()
	res.TagNote.New()
	res.TagAlarm.New()
}

//GetByRedis
func (res *TagType) GetByRedis(dbIndex int) ResultType {

	var resultVal ResultType

	res.TagMain.TagId = res.TagId
	resultVal = res.TagMain.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagBase.TagId = res.TagId
	resultVal = res.TagBase.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagGps.TagId = res.TagId
	resultVal = res.TagGps.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagStatu.TagId = res.TagId
	resultVal = res.TagStatu.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagReader.TagId = res.TagId
	resultVal = res.TagReader.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagNote.TagId = res.TagId
	resultVal = res.TagNote.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagAlarm.TagId = res.TagId
	resultVal = res.TagAlarm.GetByRedis()
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByEpc
func (res *TagType) GetByRedisByEpc(epc string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_TAG_EPC, epc)
	if resultVal.Result == RESULT_OK {
		var tagId string = resultVal.Retval.(string)
		res.TagId = StringIdToFloat64(tagId)
		resultVal = res.GetByRedis()
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
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
