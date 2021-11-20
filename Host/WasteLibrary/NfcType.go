package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//NfcType
type NfcType struct {
	NfcId     float64
	NfcMain   NfcMainType
	NfcBase   NfcBaseType
	NfcStatu  NfcStatuType
	NfcReader NfcReaderType
}

//New
func (res *NfcType) New() {
	res.NfcId = 0
	res.NfcMain.New()
	res.NfcBase.New()
	res.NfcStatu.New()
	res.NfcReader.New()
}

//GetAll
func (res *NfcType) GetAll() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_NFC_MAINS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.NfcMain = StringToNfcMainType(resultVal.Retval.(string))
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_NFC_BASES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.NfcBase = StringToNfcBaseType(resultVal.Retval.(string))
		res.NfcBase.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_NFC_STATUS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.NfcStatu = StringToNfcStatuType(resultVal.Retval.(string))
		res.NfcStatu.NewData = false
	} else {
		return resultVal
	}
	resultVal = GetRedisForStoreApi(REDIS_NFC_READERS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.NfcReader = StringToNfcReaderType(resultVal.Retval.(string))
		res.NfcReader.NewData = false
	} else {
		return resultVal
	}
	return resultVal
}

//ToId String
func (res *NfcType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToNfcId String
func (res *NfcType) ToNfcIdString() string {
	return fmt.Sprintf("%.0f", res.NfcId)
}

//ToByte
func (res *NfcType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *NfcType) ToString() string {
	return string(res.ToByte())

}

//Byte To NfcType
func ByteToNfcType(retByte []byte) NfcType {
	var retVal NfcType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To NfcType
func StringToNfcType(retStr string) NfcType {
	return ByteToNfcType([]byte(retStr))
}
