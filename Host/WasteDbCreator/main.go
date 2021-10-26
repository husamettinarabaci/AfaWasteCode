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
			DevıceId INT NOT NULL DEFAULT -1,
			OpType varchar(50) NOT NULL DEFAULT '',
			Data TEXT NOT NULL DEFAULT '',
			CustomerId INT NOT NULL DEFAULT -1,
			DataTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
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

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tagdata ( 
		DataId serial PRIMARY KEY,
		TagID INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		Epc varchar(50) NOT NULL DEFAULT '',
		ContainerNo varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TAG_STATU_EMPTY + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_devicedata ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RFID_DEVICE_TYPE_NONE + `',
		TruckType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TRUCK_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		Speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		PlateNo varchar(50) NOT NULL DEFAULT '',
        DriverName varchar(50) NOT NULL DEFAULT '',
        DriverSurName varchar(50) NOT NULL DEFAULT '',
        GpsAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1'
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_devicedata ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_DEVICE_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Battery varchar(50) NOT NULL DEFAULT '0',
		BatteryStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.BATTERY_STATU_NONE + `',
		BatteryTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltRange INT NOT NULL DEFAULT 0,
		UltStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_STATU_NONE + `',
		Imei  varchar(50) NOT NULL DEFAULT '',
		Imsi  varchar(50) NOT NULL DEFAULT '',
		FirmwareVersion  varchar(50) NOT NULL DEFAULT '1',
		OldLatitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		OldLongitude NUMERIC(14, 11)  NOT NULL DEFAULT 0
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_devicedata ( 
		DataId serial PRIMARY KEY,
		DeviceId INT NOT NULL DEFAULT -1,
		CustomerId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_DEVICE_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		TotalGlassCount INT NOT NULL DEFAULT 0,
		TotalPlasticCount INT NOT NULL DEFAULT 0,
		TotalMetalCount INT NOT NULL DEFAULT 0,
		DailyGlassCount INT NOT NULL DEFAULT 0,
		DailyPlasticCount INT NOT NULL DEFAULT 0,
		DailyMetalCount INT NOT NULL DEFAULT 0,
		RecyTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		WebAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppVersion varchar(50) NOT NULL DEFAULT '1',
		MotorAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1'
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

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tags ( 
		TagID serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceId INT NOT NULL DEFAULT -1,
		UID varchar(50) NOT NULL DEFAULT '',
		Epc varchar(50) NOT NULL DEFAULT '',
		ContainerNo varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TAG_STATU_EMPTY + `',
		ImageStatu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		ReadTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		CheckTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS rfid_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RFID_DEVICE_TYPE_NONE + `',
		TruckType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.TRUCK_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		Speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		PlateNo varchar(50) NOT NULL DEFAULT '',
        DriverName varchar(50) NOT NULL DEFAULT '',
        DriverSurName varchar(50) NOT NULL DEFAULT '',
        GpsAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1'
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS ult_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_DEVICE_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Battery varchar(50) NOT NULL DEFAULT '0',
		BatteryStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.BATTERY_STATU_NONE + `',
		BatteryTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UltRange INT NOT NULL DEFAULT 0,
		UltStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ULT_STATU_NONE + `',
		Imei  varchar(50) NOT NULL DEFAULT '',
		Imsi  varchar(50) NOT NULL DEFAULT '',
		FirmwareVersion  varchar(50) NOT NULL DEFAULT '1',
		OldLatitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		OldLongitude NUMERIC(14, 11)  NOT NULL DEFAULT 0
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS recy_devices (
		DeviceId  serial PRIMARY KEY,
		CustomerId INT NOT NULL DEFAULT -1,
		ContainerNo  varchar(50) NOT NULL DEFAULT '',
		ContainerType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.CONTAINER_TYPE_NONE + `',
		DeviceType  varchar(50) NOT NULL DEFAULT '` + WasteLibrary.RECY_DEVICE_TYPE_NONE + `',
		SerialNumber  varchar(50) NOT NULL DEFAULT '',
		StatusTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AliveStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		AliveLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		Latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		Longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		GpsTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_STATU_NONE + `',
		AlarmTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		AlarmType varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ALARM_NONE + `',
		Alarm varchar(50) NOT NULL DEFAULT '',
		Therm varchar(50) NOT NULL DEFAULT '0',
		ThermTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		ThermStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.THERM_STATU_NONE + `',
		Active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_ACTIVE + `',
		CreateTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		TotalGlassCount INT NOT NULL DEFAULT 0,
		TotalPlasticCount INT NOT NULL DEFAULT 0,
		TotalMetalCount INT NOT NULL DEFAULT 0,
		DailyGlassCount INT NOT NULL DEFAULT 0,
		DailyPlasticCount INT NOT NULL DEFAULT 0,
		DailyMetalCount INT NOT NULL DEFAULT 0,
		RecyTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorConnStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorConnLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		MotorStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		MotorLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppStatus varchar(50) NOT NULL DEFAULT '` + WasteLibrary.STATU_PASSIVE + `',
		WebAppLastOkTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		WebAppVersion varchar(50) NOT NULL DEFAULT '1',
		MotorAppVersion varchar(50) NOT NULL DEFAULT '1',
        ThermAppVersion varchar(50) NOT NULL DEFAULT '1',
        TransferAppVersion varchar(50) NOT NULL DEFAULT '1',
        CheckerAppVersion varchar(50) NOT NULL DEFAULT '1',
        CamAppVersion varchar(50) NOT NULL DEFAULT '1',
        ReaderAppVersion varchar(50) NOT NULL DEFAULT '1',
        SystemAppVersion varchar(50) NOT NULL DEFAULT '1'
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
