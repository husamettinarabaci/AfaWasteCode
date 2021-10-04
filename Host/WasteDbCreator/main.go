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
			data_id serial PRIMARY KEY,
			app_type varchar(50) NOT NULL DEFAULT '',
			device_id INT NOT NULL DEFAULT -1,
			optype varchar(50) NOT NULL DEFAULT '',
			data TEXT NOT NULL DEFAULT '',
			customer_id INT NOT NULL DEFAULT -1,
			  data_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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
			data_id serial PRIMARY KEY,
			hashkey TEXT NOT NULL DEFAULT '',
			subkey TEXT NOT NULL DEFAULT '', 
			keyvalue TEXT NOT NULL DEFAULT '',
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
		data_id serial PRIMARY KEY,
		tag_id serial PRIMARY KEY,
		customer_id INT NOT NULL DEFAULT -1,
		device_id INT NOT NULL DEFAULT -1,
		epc varchar(50) NOT NULL DEFAULT '',
		uid varchar(50) NOT NULL DEFAULT '',
		container_no varchar(50) NOT NULL DEFAULT '',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.EMPTY + `',
		image_statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `',
		read_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		check_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = readerDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS devicedata ( 
		data_id serial PRIMARY KEY,
		device_id serial PRIMARY KEY,
		device_type varchar(50) NOT NULL DEFAULT '',
		serial_number varchar(50) NOT NULL DEFAULT '',
		device_name varchar(50) NOT NULL DEFAULT '',
		customer_id INT NOT NULL DEFAULT -1,
		reader_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		reader_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		reader_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		transfer_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		alive_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		contact_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm varchar(50) NOT NULL DEFAULT '0',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_range NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		device_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		total_glass_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		total_metal_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		total_plastic_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		alarm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		alarm_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		battery_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		battery varchar(50) NOT NULL DEFAULT '0',
		alarm_type varchar(50) NOT NULL DEFAULT '` + WasteLibrary.NONE + `',
		alarm varchar(50) NOT NULL DEFAULT '',
		recy_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		contact_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `', 
		therm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		gps_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		status_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		battery_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		therm_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		transfer_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		alive_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		contact_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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
		tag_id serial PRIMARY KEY,
		customer_id INT NOT NULL DEFAULT -1,
		device_id INT NOT NULL DEFAULT -1,
		epc varchar(50) NOT NULL DEFAULT '',
		uid varchar(50) NOT NULL DEFAULT '',
		container_no varchar(50) NOT NULL DEFAULT '',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.EMPTY + `',
		image_statu varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `',
		read_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		check_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS devices ( 
		device_id serial PRIMARY KEY,
		device_type varchar(50) NOT NULL DEFAULT '',
		serial_number varchar(50) NOT NULL DEFAULT '',
		device_name varchar(50) NOT NULL DEFAULT '',
		customer_id INT NOT NULL DEFAULT -1,
		reader_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		reader_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		reader_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		cam_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_conn_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		gps_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		transfer_app_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		alive_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		contact_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm varchar(50) NOT NULL DEFAULT '0',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		speed NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_range NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		device_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		total_glass_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		total_metal_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		total_plastic_count NUMERIC(14, 11)  NOT NULL DEFAULT 0,
		ult_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		alarm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		alarm_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		therm_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		battery_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		battery varchar(50) NOT NULL DEFAULT '0',
		alarm_type varchar(50) NOT NULL DEFAULT '` + WasteLibrary.NONE + `',
		alarm varchar(50) NOT NULL DEFAULT '',
		recy_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		contact_status varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `', 
		therm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		gps_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		status_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reader_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		gps_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		battery_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_conn_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cam_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		therm_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		transfer_app_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		alive_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		contact_last_ok_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS customers ( 
		customer_id serial PRIMARY KEY,
		customer_name varchar(50) NOT NULL DEFAULT '',
		customer_link varchar(50) NOT NULL DEFAULT '',
		rfid_app varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		ult_app varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		recy_app varchar(50) NOT NULL DEFAULT '` + WasteLibrary.PASSIVE + `',
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `',
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users (
		user_id serial PRIMARY KEY,
		user_type varchar(50) NOT NULL DEFAULT '` + WasteLibrary.USER_GUEST + `',
		email varchar(50) NOT NULL DEFAULT '',
		user_name varchar(50) NOT NULL DEFAULT '',
		customer_id INT NOT NULL DEFAULT -1,
		active varchar(50) NOT NULL DEFAULT '` + WasteLibrary.ACTIVE + `',
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = staticDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	staticDb.Close()
}
