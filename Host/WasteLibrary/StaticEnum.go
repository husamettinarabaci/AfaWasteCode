package WasteLibrary

const (
	ULT_EMPTY  = "EMPTY"
	ULT_HALF   = "HALF"
	ULT_FULL   = "FULL"
	ULT_NORMAL = "NORMAL"
	ULT_NONE   = "NONE"
)

const (
	CONTAINER_EMPTY  = "EMPTY"
	CONTAINER_HALF   = "HALF"
	CONTAINER_FULL   = "FULL"
	CONTAINER_ERROR  = "ERROR"
	CONTAINER_NORMAL = "NORMAL"
	CONTAINER_NONE   = "NONE"
)

const (
	RESULT_FAIL = "FAIL"
	RESULT_OK   = "OK"
)

const (
	STATU_ACTIVE  = "ACTIVE"
	STATU_PASSIVE = "PASSIVE"
	STATU_APP     = "APP"
	STATU_CONN    = "CONN"
	STATU_READER  = "READER"
	STATU_CAM     = "CAM"
	STATU_GPS     = "GPS"
)

const (
	GENERAL_OPTYPE = "OPTYPE"
	GENERAL_PORT   = "PORT"
	GENERAL_TYPE   = "TYPE"
)

const (
	AFATEK_CUSTOMER = "CUSTOMER"
	AFATEK_DEVICE   = "DEVICE"
)

const (
	OPTYPE_USER = "USER"
	OPTYPE_TAG  = "TAG"
)

const (
	CONFIG_CUSTOMER = "CUSTOMER"
	CONFIG_ADMIN    = "ADMIN"
	CONFIG_LOCAL    = "LOCAL"
)

const (
	RFID_READER_OKBIT    = "5379"
	RFID_READER_STARTBIT = "4354"
	RFID_TAG_PATTERN     = "AFA09012018AFA"
)

const (
	APP_STATUS  = "STATUS"
	APP_THERM   = "THERM"
	APP_RF      = "RF"
	APP_CAM     = "CAM"
	APP_SENS    = "SENS"
	APP_ARVENTO = "ARVENTO"
	APP_GPS     = "GPS"
)

const (
	ALARM_NONE  = "NONE"
	ALARM_ALARM = "ALARM"
	ALARM_ATMP  = "ATMP"
	ALARM_AGPS  = "AGPS"
)

const (
	DEVICE_ULT  = "ULT"
	DEVICE_RFID = "RFID"
	DEVICE_RECY = "RECY"
)

const (
	HTTP_HEADER = "HEADER"
	HTTP_DATA   = "DATA"
)

const (
	LOGGER_CONTAINER = "CONTAINER"
	LOGGER_LOGTYPE   = "LOGTYPE"
	LOGGER_FUNC      = "FUNC"
	LOGGER_LOG       = "LOG"
	LOGGER_DEBUG     = "DEBUG"
	LOGGER_ERROR     = "ERROR"
	LOGGER_INFO      = "INFO"
)

const (
	REDIS_SERIAL_DEVICE           = "serial-device"
	REDIS_DEVICES                 = "devices"
	REDIS_CUSTOMER_TAGS           = "customer-tags"
	REDIS_CUSTOMER_DEVICES        = "customer-devices"
	REDIS_CUSTOMERS               = "customers"
	REDIS_CUSTOMER_CUSTOMERCONFIG = "customer-customerconfig"
	REDIS_CUSTOMER_ADMINCONFIG    = "customer-adminconfig"
	REDIS_CUSTOMER_LOCALCONFIG    = "customer-localconfig"
	REDIS_CUSTOMER_LINK           = "customer-link"
	REDIS_USERS                   = "users"
	REDIS_CUSTOMER_USERS          = "customer-users"
	REDIS_HASHKEY                 = "HASHKEY"
	REDIS_SUBKEY                  = "SUBKEY"
	REDIS_KEYVALUE                = "KEYVALUE"
	REDIS_NOT                     = "NOT"
)

const (
	USER_GUEST  = "GUEST"
	USER_ADMIN  = "ADMIN"
	USER_REPORT = "REPORT"
	USER_TOKEN  = "TOKEN"
)
