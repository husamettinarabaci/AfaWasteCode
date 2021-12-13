package WasteLibrary

import "encoding/json"

//StringArrayToByte
func StringArrayToByte(res []string) []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//StringArrayToString Get JSON
func StringArrayToString(res []string) string {
	return string(StringArrayToByte(res))

}

//ByteToStringArray
func ByteToStringArray(retByte []byte) []string {
	var res []string
	json.Unmarshal(retByte, &res)
	return res
}

//StringToStringArray
func StringToStringArray(retStr string) []string {
	res := ByteToStringArray([]byte(retStr))
	return res
}

//MapStringStringToByte
func MapStringStringToByte(res map[string]string) []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//MapStringStringToString Get JSON
func MapStringStringToString(res map[string]string) string {
	return string(MapStringStringToByte(res))

}

//ByteToMapStringString
func ByteToMapStringString(retByte []byte) map[string]string {
	var res map[string]string
	json.Unmarshal(retByte, &res)
	return res
}

//StringToMapStringString
func StringToMapStringString(retStr string) map[string]string {
	res := ByteToMapStringString([]byte(retStr))
	return res
}
