package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//LocalConfigType
type LocalConfigType struct {
	CustomerId float64
	Active     string
	CreateTime string
	Locs       map[string]string
}

//New
func (res *LocalConfigType) New() {
	res.CustomerId = 1
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
	res.Locs = make(map[string]string)
	res.Locs[RESULT_SUCCESS_OK] = "İşlem Başarılı"
	res.Locs[RESULT_ERROR_NONE] = "Bilinmeyen Hata"
	res.Locs[RESULT_ERROR_HTTP_PARSE] = "İstek Çözülmedi"
	res.Locs[RESULT_ERROR_HTTP_POST] = "İstek Hatası"
	res.Locs[RESULT_ERROR_USER_AUTH] = "Hatalı Kullanıcı"
	res.Locs[RESULT_ERROR_READERTYPE] = "Okuma Türü Hatalı"
	res.Locs[RESULT_ERROR_DATATYPE] = "Veri Türü Hatalı"
	res.Locs[RESULT_ERROR_DEVICETYPE] = "Cihaz Tipi Hatalı"
	res.Locs[RESULT_ERROR_DEVICE_NOTFOUND] = "Cihaz Bulunamadı"
	res.Locs[RESULT_ERROR_DEVICES_NOTFOUND] = "Cihazlar Bulunamadı"
	res.Locs[RESULT_ERROR_DB_SAVE] = "Veri Tabanı Kayıt Hatası"
	res.Locs[RESULT_ERROR_DB_GET] = "Veri Tabanı Sorgu Hatası"
	res.Locs[RESULT_ERROR_REDIS_SAVE] = "Hafıza Kayıt Hatası"
	res.Locs[RESULT_ERROR_REDIS_GET] = "Hafıza Sorgu Hatası"
	res.Locs[RESULT_ERROR_CUSTOMER_NOTFOUND] = "Müşteri Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMER_INVALID] = "Hatalı Müşteri"
	res.Locs[RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND] = "Müşteri Ayarı Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND] = "Admin Ayarı Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMER_LOCALCONFIG_NOTFOUND] = "Dil Ayarı Bulunamdı"
	res.Locs[RESULT_ERROR_CUSTOMER_TAGS_NOTFOUND] = "Tagler Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMER_DEVICES_NOTFOUND] = "Müşteri Cihazları Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMER_NFCS_NOTFOUND] = "Müşteri Kartları Bulunamadı"
	res.Locs[RESULT_ERROR_CUSTOMERS_NOTFOUND] = "Müşteriler Bulunamadı"
	res.Locs[RESULT_ERROR_TAG_NOTFOUND] = "Tag Bulunamadı"
	res.Locs[RESULT_ERROR_TAG_CUSTOMER_NOTFOUND] = "Tag Müşteri Ataması Bulunamadı"
	res.Locs[RESULT_ERROR_NFC_NOTFOUND] = "Kart Bulunamadı"
	res.Locs[RESULT_ERROR_NFC_CUSTOMER_NOTFOUND] = "Kart Müşteri Ataması Bulunamadı"
	res.Locs[RESULT_ERROR_APP_STARTED] = "Uygulama Başlatılamadı"
	res.Locs[RESULT_ERROR_IGNORE_FIRST_DATA] = "Başlangıç Verisi Hatalı"
}

//GetByRedis
func (res *LocalConfigType) GetByRedis() ResultType {
	resultVal := GetRedisForStoreApi("0", REDIS_CUSTOMER_LOCALCONFIG, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *LocalConfigType) SaveToRedis() ResultType {
	resultVal := SaveRedisForStoreApi(REDIS_CUSTOMER_LOCALCONFIG, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *LocalConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *LocalConfigType) ToString() string {
	return string(res.ToByte())

}

//ByteToType
func (res *LocalConfigType) ByteToType(retByte []byte) {
	res.New()
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *LocalConfigType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
