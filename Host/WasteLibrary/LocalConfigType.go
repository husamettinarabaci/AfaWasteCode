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

	Durum       string
	KonteynerID string
	Konum       string
	GrafikDetay string
	Uyari       string

	Ozet    string
	GunOnce string

	Arama   string
	AraAdet string

	Kamyonlar        string
	Kamyon           string
	Vinc             string
	Tip              string
	Plaka            string
	Sofor            string
	ParkHalinde      string
	HareketHalinde   string
	GunlukMesafe     string
	GunlukTopKonAdet string
	VerimlilikOrani  string
	MesaiBitisi      string
	MesaiBitti       string

	RfidTakipSistemi    string
	Toplanmayan         string
	Toplanan            string
	GunlukToplanmaOrani string
	ToplanmaOranlari    string
	Toplandi            string
	Toplanmadi          string
	SonToplanmaTarihi   string
	SonOkunmaTarihi     string
	ToplanmaSikligi     string
	GundurToplanmadi    string

	DolulukSensorleri     string
	BelirliDoluluk        string
	GunlukDolulukOrani    string
	Bos                   string
	AzDolu                string
	OrtaDolu              string
	Dolu                  string
	DolulukOrani          string
	DolulukOranlari       string
	SicaklikDegeri        string
	KonteynerTipi         string
	DolmaHiziPuani        string
	GecmisDolulukOranlari string

	GeriDonusumCihazlari string
	Plastik              string
	Cam                  string
	Metal                string
	KurtAgacSay          string
	GunlukGeriDonOran    string
	ToplamGeriDonSayisi  string
	PlastikSayisi        string
	CamSayisi            string
	MetalSayisi          string
}

//New
func (res LocalConfigType) New() {
	res.CustomerId = 0
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()

	res.Durum = "Durum"
	res.KonteynerID = "Konteyner ID"
	res.Konum = "Konum"
	res.GrafikDetay = "Grafik için tıkayın"
	res.Uyari = "Uyarı"

	res.Ozet = "Özet"
	res.GunOnce = "gün önce"

	res.Arama = "Arama"
	res.AraAdet = "Kamyon ve konteyner bilgisi için en az 3 karakter girmelisiniz"

	res.Kamyonlar = "Kamyonlar"
	res.Kamyon = "Kamyon"
	res.Vinc = "Vinç"
	res.Tip = "Tip"
	res.Plaka = "Plaka No"
	res.Sofor = "Şöför"
	res.ParkHalinde = "Park Halinde"
	res.HareketHalinde = "Hareket Ediyor"
	res.GunlukMesafe = "Günlük Mesafe"
	res.GunlukTopKonAdet = "Günlük Toplanılan Konteyner Sayısı"
	res.VerimlilikOrani = "Verimlilik Oranı"
	res.MesaiBitisi = "Mesai Bitişi"
	res.MesaiBitti = "Mesai Bitti"

	res.RfidTakipSistemi = "RFID Takip Sistemi"
	res.Toplanmayan = "Toplanmayan"
	res.Toplanan = "Toplanan"
	res.GunlukToplanmaOrani = "Günlük Toplanma Oranı"
	res.ToplanmaOranlari = "Toplanma Oranları"
	res.Toplandi = "Toplandı"
	res.Toplanmadi = "Toplanmadı"
	res.SonToplanmaTarihi = "Son Toplanma Tarihi"
	res.SonOkunmaTarihi = "Son Okunma Tarihi"
	res.ToplanmaSikligi = "Toplanma Sıklığı"
	res.GundurToplanmadi = "gündür toplanmadı"

	res.DolulukSensorleri = "Doluluk Sensörleri"
	res.BelirliDoluluk = "Belirli doluluk oranları aralığındaki toplam konteyner sayısını gösterir."
	res.GunlukDolulukOrani = "Günlük Doluluk Oranı"
	res.Bos = "Boş"
	res.AzDolu = "Az Dolu"
	res.OrtaDolu = "Orta Dolu"
	res.Dolu = "Dolu"
	res.DolulukOrani = "Doluluk Oranı"
	res.DolulukOranlari = "Doluluk Oranları"
	res.SicaklikDegeri = "Sıcaklık Değeri"
	res.KonteynerTipi = "Konteyner Tipi"
	res.DolmaHiziPuani = "Dolma Hızı Puanı"
	res.GecmisDolulukOranlari = "Geçmiş Doluluk Oranları"

	res.GeriDonusumCihazlari = "Geri Dönüşüm Cihazları"
	res.Plastik = "Plastik"
	res.Cam = "Cam"
	res.Metal = "Metal"
	res.KurtAgacSay = "Kurtarılan Ağaç Sayısı"
	res.GunlukGeriDonOran = "Günlük Geri Dönüşüm Oranları"
	res.ToplamGeriDonSayisi = "Toplam Geri Dönüşüm Sayısı"
	res.PlastikSayisi = "Plastik Sayısı"
	res.CamSayisi = "Cam Sayısı"
	res.MetalSayisi = "Metal Sayısı"

}

//ToId String
func (res LocalConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res LocalConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res LocalConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To LocalConfigType
func ByteToLocalConfigType(retByte []byte) LocalConfigType {
	var retVal LocalConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To LocalConfigType
func StringToLocalConfigType(retStr string) LocalConfigType {
	return ByteToLocalConfigType([]byte(retStr))
}
