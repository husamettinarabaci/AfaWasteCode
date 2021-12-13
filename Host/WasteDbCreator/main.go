package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/devafatek/WasteLibrary"
	_ "github.com/lib/pq"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	time.Sleep(time.Second * 10)
}
func main() {

	initStart()
	bulkDbSet()
	configDbSet()
	sumDbSet()
	staticDbSet()
	readerDbSet()

}

func bulkDbSet() {
	var bulkDbHost string = "waste-bulkdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err := sql.Open("postgres", bulkDbInfo)
	WasteLibrary.LogErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	WasteLibrary.LogErr(err)
	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ` + WasteLibrary.DATATYPE_LISTENER_DATA + `  ( 
			DataId serial PRIMARY KEY,
			AppType varchar(50) NOT NULL DEFAULT '',
			DeviceNo varchar(50) NOT NULL DEFAULT '',
			DeviceId INT NOT NULL DEFAULT -1,
			CustomerId INT NOT NULL DEFAULT -1,
			Time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			Repeat varchar(50) NOT NULL DEFAULT '',
			DeviceType varchar(50) NOT NULL DEFAULT '',
			ReaderType varchar(50) NOT NULL DEFAULT '',
			DataType varchar(50) NOT NULL DEFAULT '',
			Token varchar(50) NOT NULL DEFAULT '',
			Data TEXT NOT NULL DEFAULT '',
			CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`)
	_, err = bulkDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	bulkDb.Close()
}

func configDbSet() {
	var configDbHost string = "waste-configdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	configDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configDbHost, port, user, password, dbname)

	configDb, err := sql.Open("postgres", configDbInfo)
	WasteLibrary.LogErr(err)
	defer configDb.Close()

	err = configDb.Ping()
	WasteLibrary.LogErr(err)

	var userT WasteLibrary.UserType
	userT.CreateDb(configDb)

	var customer WasteLibrary.CustomerType
	customer.CreateDb(configDb)

	configDb.Close()
}

func sumDbSet() {
	var sumDbHost string = "waste-sumdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	sumDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sumDbHost, port, user, password, dbname)

	sumDb, err := sql.Open("postgres", sumDbInfo)
	WasteLibrary.LogErr(err)
	defer sumDb.Close()

	err = sumDb.Ping()
	WasteLibrary.LogErr(err)

	for i := 0; i < 31; i++ {

		var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS `+WasteLibrary.DATATYPE_REDIS_DATA+`_%d ( 
			DataId serial PRIMARY KEY,
			HashKey TEXT NOT NULL DEFAULT '',
			SubKey TEXT NOT NULL DEFAULT '',
			KeyValue TEXT NOT NULL DEFAULT '',
			  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`, i)
		_, err = sumDb.Exec(createSQL)
		WasteLibrary.LogErr(err)
	}
	sumDb.Close()
}

func readerDbSet() {
	var readerDbHost string = "waste-readerdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	readerDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		readerDbHost, port, user, password, dbname)

	readerDb, err := sql.Open("postgres", readerDbInfo)
	WasteLibrary.LogErr(err)
	defer readerDb.Close()

	err = readerDb.Ping()
	WasteLibrary.LogErr(err)

	var tag WasteLibrary.TagType
	tag.TagMain.CreateReaderDb(readerDb)

	var nfc WasteLibrary.NfcType
	nfc.NfcMain.CreateReaderDb(readerDb)

	var rfid WasteLibrary.RfidDeviceType
	rfid.DeviceMain.CreateReaderDb(readerDb)

	var ult WasteLibrary.UltDeviceType
	ult.DeviceMain.CreateReaderDb(readerDb)

	var recy WasteLibrary.RecyDeviceType
	recy.DeviceMain.CreateReaderDb(readerDb)

	readerStaticDbCumulative(readerDb)
	readerDb.Close()
}

func staticDbSet() {
	var staticDbHost string = "waste-staticdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	staticDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		staticDbHost, port, user, password, dbname)

	staticDb, err := sql.Open("postgres", staticDbInfo)
	WasteLibrary.LogErr(err)
	defer staticDb.Close()

	err = staticDb.Ping()
	WasteLibrary.LogErr(err)

	var tag WasteLibrary.TagType
	tag.TagMain.CreateDb(staticDb)

	var nfc WasteLibrary.NfcType
	nfc.NfcMain.CreateDb(staticDb)

	var rfid WasteLibrary.RfidDeviceType
	rfid.DeviceMain.CreateDb(staticDb)

	var ult WasteLibrary.UltDeviceType
	ult.DeviceMain.CreateDb(staticDb)

	var recy WasteLibrary.RecyDeviceType
	recy.DeviceMain.CreateDb(staticDb)

	readerStaticDbCumulative(staticDb)
	staticDb.Close()
}

func readerStaticDbCumulative(currentDb *sql.DB) {

	var tag WasteLibrary.TagType
	tag.New()
	tag.TagBase.CreateDb(currentDb)
	tag.TagStatu.CreateDb(currentDb)
	tag.TagGps.CreateDb(currentDb)
	tag.TagReader.CreateDb(currentDb)
	tag.TagNote.CreateDb(currentDb)
	tag.TagAlarm.CreateDb(currentDb)

	var nfc WasteLibrary.NfcType
	nfc.NfcBase.CreateDb(currentDb)
	nfc.NfcStatu.CreateDb(currentDb)
	nfc.NfcReader.CreateDb(currentDb)

	var rfid WasteLibrary.RfidDeviceType
	rfid.DeviceBase.CreateDb(currentDb)
	rfid.DeviceStatu.CreateDb(currentDb)
	rfid.DeviceGps.CreateDb(currentDb)
	rfid.DeviceEmbededGps.CreateDb(currentDb)
	rfid.DeviceAlarm.CreateDb(currentDb)
	rfid.DeviceTherm.CreateDb(currentDb)
	rfid.DeviceVersion.CreateDb(currentDb)
	rfid.DeviceWorkHour.CreateDb(currentDb)
	rfid.DeviceNote.CreateDb(currentDb)
	rfid.DeviceReport.CreateDb(currentDb)

	var ult WasteLibrary.UltDeviceType
	ult.DeviceBase.CreateDb(currentDb)
	ult.DeviceStatu.CreateDb(currentDb)
	ult.DeviceBattery.CreateDb(currentDb)
	ult.DeviceGps.CreateDb(currentDb)
	ult.DeviceAlarm.CreateDb(currentDb)
	ult.DeviceTherm.CreateDb(currentDb)
	ult.DeviceVersion.CreateDb(currentDb)
	ult.DeviceSens.CreateDb(currentDb)
	ult.DeviceNote.CreateDb(currentDb)
	ult.DeviceSim.CreateDb(currentDb)

	var recy WasteLibrary.RecyDeviceType
	recy.DeviceBase.CreateDb(currentDb)
	recy.DeviceStatu.CreateDb(currentDb)
	recy.DeviceGps.CreateDb(currentDb)
	recy.DeviceAlarm.CreateDb(currentDb)
	recy.DeviceTherm.CreateDb(currentDb)
	recy.DeviceVersion.CreateDb(currentDb)
	recy.DeviceDetail.CreateDb(currentDb)
	recy.DeviceNote.CreateDb(currentDb)
}
