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
	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS listenerdata ( 
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

	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users ( 
		UserId serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		UserName varchar(50) NOT NULL DEFAULT '',
		UserRole varchar(50) NOT NULL DEFAULT '` + WasteLibrary.USER_ROLE_GUEST + `',
		Password varchar(50) NOT NULL DEFAULT '',
		Email varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = configDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

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

	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS redisdata ( 
			DataId serial PRIMARY KEY,
			HashKey TEXT NOT NULL DEFAULT '',
			SubKey TEXT NOT NULL DEFAULT '',
			KeyValue TEXT NOT NULL DEFAULT '',
			  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`)
	_, err = sumDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
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

	var createSQL string

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_mains ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		Epc varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_bases ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		ContainerNo varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINERTYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_status ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		ContainerStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_FULLNESS_STATU_NONE + `',
		TagStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TAG_STATU_NONE + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_gpses ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_readers ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_mains ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		Epc varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_bases ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		Name varchar(50) NOT NULL DEFAULT '',
		SurName varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_status ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		ItemStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_ITEM_STATU_NONE + `',
		NfcStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.NFC_STATU_NONE + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_readers ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_main_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RFID_DEVICE_TYPE_NONE + `',
		TruckType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TRUCKTYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ThermAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		TransferAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		TransferAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		SystemAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		SystemAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UpdaterAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		UpdaterAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ContactStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ContactLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
        GpsAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_detail_devices ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		PlateNo varchar(50) NOT NULL DEFAULT '',
        DriverName varchar(50) NOT NULL DEFAULT '',
        DriverSurName varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_workhour_devices ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		WorkStartHour  INT NOT NULL DEFAULT 6,
		WorkStartMinute  INT NOT NULL DEFAULT 0,
		WorkEndHour  INT NOT NULL DEFAULT 18,
		WorkEndMinute  INT NOT NULL DEFAULT 30,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_main_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		OldLatitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		OldLongitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINERTYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_DEVICE_TYPE_NONE + `',
		Imei  varchar(50) NOT NULL DEFAULT '',
		Imsi  varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_STATU_NONE + `',
		SensPercent NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ContainerStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_FULLNESS_STATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_battery_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Battery varchar(50) NOT NULL DEFAULT '0',
		BatteryStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.BATTERYSTATU_NONE + `',
		BatteryTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		FirmwareVersion  varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_sens_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		UltTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltCount INT NOT NULL DEFAULT 0,
		UltRange1  INT NOT NULL DEFAULT 0,
        UltRange2  INT NOT NULL DEFAULT 0,
        UltRange3  INT NOT NULL DEFAULT 0,
        UltRange4  INT NOT NULL DEFAULT 0,
        UltRange5  INT NOT NULL DEFAULT 0,
        UltRange6  INT NOT NULL DEFAULT 0,
        UltRange7  INT NOT NULL DEFAULT 0,
        UltRange8  INT NOT NULL DEFAULT 0,
        UltRange9  INT NOT NULL DEFAULT 0,
        UltRange10 INT NOT NULL DEFAULT 0,
        UltRange11 INT NOT NULL DEFAULT 0,
        UltRange12 INT NOT NULL DEFAULT 0,
        UltRange13 INT NOT NULL DEFAULT 0,
        UltRange14 INT NOT NULL DEFAULT 0,
        UltRange15 INT NOT NULL DEFAULT 0,
        UltRange16 INT NOT NULL DEFAULT 0,
        UltRange17 INT NOT NULL DEFAULT 0,
        UltRange18 INT NOT NULL DEFAULT 0,
        UltRange19 INT NOT NULL DEFAULT 0,
        UltRange20 INT NOT NULL DEFAULT 0,
        UltRange21 INT NOT NULL DEFAULT 0,
        UltRange22 INT NOT NULL DEFAULT 0,
        UltRange23 INT NOT NULL DEFAULT 0,
        UltRange24 INT NOT NULL DEFAULT 0,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	// deleteSQL := fmt.Sprintf(`ALTER TABLE ult_sens_devices DROP COLUMN IF EXISTS UltStatus;`)
	// _, err = readerDb.Exec(deleteSQL)
	// WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_main_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_DEVICE_TYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ThermAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		TransferAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		TransferAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		SystemAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		SystemAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UpdaterAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		UpdaterAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		WebAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		WebAppVersion varchar(50) NOT NULL DEFAULT '1',
		MotorAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_detail_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		TotalGlassCount INT NOT NULL DEFAULT 0,
		TotalPlasticCount INT NOT NULL DEFAULT 0,
		TotalMetalCount INT NOT NULL DEFAULT 0,
		DailyGlassCount INT NOT NULL DEFAULT 0,
		DailyPlasticCount INT NOT NULL DEFAULT 0,
		DailyMetalCount INT NOT NULL DEFAULT 0,
		RecyTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
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
	var createSQL string

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_mains ( 
		TagID serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		Epc varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_bases ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		ContainerNo varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINERTYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_status ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		ContainerStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_FULLNESS_STATU_NONE + `',
		TagStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TAG_STATU_NONE + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_gpses ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tag_readers ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_mains ( 
		NfcID serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		Epc varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_bases ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		Name varchar(50) NOT NULL DEFAULT '',
		SurName varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_status ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		ItemStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_ITEM_STATU_NONE + `',
		NfcStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.NFC_STATU_NONE + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS nfc_readers ( 
		DataId serial PRIMARY KEY,
		NfcID INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_main_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RFID_DEVICE_TYPE_NONE + `',
		TruckType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TRUCKTYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		GpsStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		GpsLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ThermAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		TransferAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		TransferAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		SystemAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		SystemAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UpdaterAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		UpdaterAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ContactStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ContactLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
        GpsAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_detail_devices ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		PlateNo varchar(50) NOT NULL DEFAULT '',
        DriverName varchar(50) NOT NULL DEFAULT '',
        DriverSurName varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_workhour_devices ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		WorkStartHour  INT NOT NULL DEFAULT 6,
		WorkStartMinute  INT NOT NULL DEFAULT 0,
		WorkEndHour  INT NOT NULL DEFAULT 6,
		WorkEndMinute  INT NOT NULL DEFAULT 0,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_main_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		OldLatitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		OldLongitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINERTYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_DEVICE_TYPE_NONE + `',
		Imei  varchar(50) NOT NULL DEFAULT '',
		Imsi  varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_STATU_NONE + `',
		SensPercent NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ContainerStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_FULLNESS_STATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_battery_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Battery varchar(50) NOT NULL DEFAULT '0',
		BatteryStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.BATTERYSTATU_NONE + `',
		BatteryTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		FirmwareVersion  varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_sens_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		UltTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltCount INT NOT NULL DEFAULT 0,
		UltRange1  INT NOT NULL DEFAULT 0,
        UltRange2  INT NOT NULL DEFAULT 0,
        UltRange3  INT NOT NULL DEFAULT 0,
        UltRange4  INT NOT NULL DEFAULT 0,
        UltRange5  INT NOT NULL DEFAULT 0,
        UltRange6  INT NOT NULL DEFAULT 0,
        UltRange7  INT NOT NULL DEFAULT 0,
        UltRange8  INT NOT NULL DEFAULT 0,
        UltRange9  INT NOT NULL DEFAULT 0,
        UltRange10 INT NOT NULL DEFAULT 0,
        UltRange11 INT NOT NULL DEFAULT 0,
        UltRange12 INT NOT NULL DEFAULT 0,
        UltRange13 INT NOT NULL DEFAULT 0,
        UltRange14 INT NOT NULL DEFAULT 0,
        UltRange15 INT NOT NULL DEFAULT 0,
        UltRange16 INT NOT NULL DEFAULT 0,
        UltRange17 INT NOT NULL DEFAULT 0,
        UltRange18 INT NOT NULL DEFAULT 0,
        UltRange19 INT NOT NULL DEFAULT 0,
        UltRange20 INT NOT NULL DEFAULT 0,
        UltRange21 INT NOT NULL DEFAULT 0,
        UltRange22 INT NOT NULL DEFAULT 0,
        UltRange23 INT NOT NULL DEFAULT 0,
        UltRange24 INT NOT NULL DEFAULT 0,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_main_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_base_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_DEVICE_TYPE_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_statu_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ReaderStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ReaderLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CamStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		CamLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		ThermAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		TransferAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		TransferAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		SystemAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		SystemAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UpdaterAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		UpdaterAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		WebAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_gps_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_alarm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMSTATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARMTYPE_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_therm_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERMSTATU_NONE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_version_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		WebAppVersion varchar(50) NOT NULL DEFAULT '1',
		MotorAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_detail_devices (
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		TotalGlassCount INT NOT NULL DEFAULT 0,
		TotalPlasticCount INT NOT NULL DEFAULT 0,
		TotalMetalCount INT NOT NULL DEFAULT 0,
		DailyGlassCount INT NOT NULL DEFAULT 0,
		DailyPlasticCount INT NOT NULL DEFAULT 0,
		DailyMetalCount INT NOT NULL DEFAULT 0,
		RecyTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS customers ( 
		CustomerId serial PRIMARY KEY,
		CustomerName varchar(50) NOT NULL DEFAULT '',
		CustomerLink varchar(50) NOT NULL DEFAULT '',
		RfIdApp varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		UltApp varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		RecyApp varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	staticDb.Close()
}
