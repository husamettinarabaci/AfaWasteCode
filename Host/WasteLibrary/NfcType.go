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

//GetByRedis
func (res *NfcType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType

	res.NfcMain.NfcId = res.NfcId
	resultVal = res.NfcMain.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.NfcBase.NfcId = res.NfcId
	resultVal = res.NfcBase.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.NfcStatu.NfcId = res.NfcId
	resultVal = res.NfcStatu.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.NfcReader.NfcId = res.NfcId
	resultVal = res.NfcReader.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
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
