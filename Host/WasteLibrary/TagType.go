package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagType
type TagType struct {
	TagId             float64
	TagMain           TagMainType
	TagBase           TagBaseType
	TagStatu          TagStatuType
	TagGps            TagGpsType
	TagReader         TagReaderType
	TagNote           TagNoteType
	TagAlarm          TagAlarmType
	TagReadDevice     TagReadDeviceType
	TagPositionChange TagPositionChangeType
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
	res.TagReadDevice.New()
	res.TagPositionChange.New()
}

//GetByRedis
func (res *TagType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType

	res.TagMain.TagId = res.TagId
	resultVal = res.TagMain.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagBase.TagId = res.TagId
	resultVal = res.TagBase.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagGps.TagId = res.TagId
	resultVal = res.TagGps.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagStatu.TagId = res.TagId
	resultVal = res.TagStatu.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagReader.TagId = res.TagId
	resultVal = res.TagReader.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagNote.TagId = res.TagId
	resultVal = res.TagNote.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagAlarm.TagId = res.TagId
	resultVal = res.TagAlarm.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagReadDevice.TagId = res.TagId
	resultVal = res.TagReadDevice.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.TagPositionChange.TagId = res.TagId
	resultVal = res.TagPositionChange.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisByEpc
func (res *TagType) GetByRedisByEpc(epc string) ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_TAG_EPC, epc)
	if resultVal.Result == RESULT_OK {
		var tagId string = resultVal.Retval.(string)
		res.TagId = StringIdToFloat64(tagId)
		resultVal = res.GetByRedis("0")
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

//ByteToType
func (res *TagType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *TagType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
