package WasteLibrary

const (
	RESULT_FAIL = "FAIL"
	RESULT_OK   = "OK"
)

const (
	STATU_ACTIVE  = "ACTIVE"
	STATU_PASSIVE = "PASSIVE"
)

const (
	HTTP_HEADER     = "HEADER"
	HTTP_DATA       = "DATA"
	HTTP_CHECKTYPE  = "CHECKTYPE"
	HTTP_USERROLE   = "USERROLE"
	HTTP_CUSTOMERID = "CUSTOMERID"
	HTTP_READERTYPE = "READERTYPE"
)

const (
	LOGGER_CONTAINER = "CONTAINER"
	LOGGER_LOGTYPE   = "LOGTYPE"
	LOGGER_FUNC      = "FUNC"
	LOGGER_LOG       = "LOG"
	LOGGER_ERROR     = "ERROR"
	LOGGER_INFO      = "INFO"
)

const (
	READERTYPE_NONE        = "NONE"
	READERTYPE_STATUS      = "STATUS"
	READERTYPE_CAM         = "CAM"
	READERTYPE_GPS         = "GPS"
	READERTYPE_THERM       = "THERM"
	READERTYPE_ULT         = "ULT"
	READERTYPE_RF          = "RF"
	READERTYPE_CAMTRIGGER  = "CAMTRIGGER"
	READERTYPE_MOTORRIGGER = "MOTORTRIGGER"
	READERTYPE_WEBTRIGGER  = "WEBTRIGGER"
	READERTYPE_ALARM       = "ALARM"
	READERTYPE_VERSION     = "VERSION"
	READERTYPE_UPDATE      = "UPDATE"
	READERTYPE_MOTOR       = "MOTOR"
	READERTYPE_WEB         = "WEB"
)

const (
	CHECKTYPE_NONE   = "NONE"
	CHECKTYPE_APP    = "APP"
	CHECKTYPE_CONN   = "CONN"
	CHECKTYPE_DEVICE = "DEVICE"
)

const (
	RFID_APPNAME_NONE     = "None"
	RFID_APPNAME_GPS      = "GpsApp"
	RFID_APPNAME_CHECKER  = "CheckerApp"
	RFID_APPNAME_READER   = "ReaderApp"
	RFID_APPNAME_TRANSFER = "TransferApp"
	RFID_APPNAME_CAM      = "CamApp"
	RFID_APPNAME_THERM    = "ThermApp"
	RFID_APPNAME_SYSTEM   = "SystemApp"
	RFID_APPNAME_UPDATER  = "UpdaterApp"
)

const (
	RFID_APPTYPE_NONE     = "NONE"
	RFID_APPTYPE_GPS      = "GPS"
	RFID_APPTYPE_CHECKER  = "STATUS"
	RFID_APPTYPE_READER   = "RF"
	RFID_APPTYPE_TRANSFER = "TRANSFER"
	RFID_APPTYPE_CAM      = "CAM"
	RFID_APPTYPE_THERM    = "THERM"
	RFID_APPTYPE_SYSTEM   = "SYSTEM"
)

const (
	RECY_APPNAME_NONE     = "None"
	RECY_APPNAME_WEB      = "WebApp"
	RECY_APPNAME_MOTOR    = "MotorApp"
	RECY_APPNAME_CHECKER  = "CheckerApp"
	RECY_APPNAME_READER   = "ReaderApp"
	RECY_APPNAME_TRANSFER = "TransferApp"
	RECY_APPNAME_CAM      = "CamApp"
	RECY_APPNAME_THERM    = "ThermApp"
	RECY_APPNAME_SYSTEM   = "SystemApp"
	RECY_APPNAME_UPDATER  = "UpdaterApp"
)

const (
	RECY_APPTYPE_NONE     = "NONE"
	RECY_APPTYPE_WEB      = "WEB"
	RECY_APPTYPE_MOTOR    = "MOTOR"
	RECY_APPTYPE_CHECKER  = "STATUS"
	RECY_APPTYPE_READER   = "RF"
	RECY_APPTYPE_TRANSFER = "TRANSFER"
	RECY_APPTYPE_CAM      = "CAM"
	RECY_APPTYPE_THERM    = "THERM"
	RECY_APPTYPE_SYSTEM   = "SYSTEM"
)

const (
	CONTAINER_FULLNESS_STATU_NONE   = "NONE"
	CONTAINER_FULLNESS_STATU_EMPTY  = "EMPTY"
	CONTAINER_FULLNESS_STATU_LITTLE = "LITTLE"
	CONTAINER_FULLNESS_STATU_MEDIUM = "MEDIUM"
	CONTAINER_FULLNESS_STATU_FULL   = "FULL"
	CONTAINER_FULLNESS_STATU_ERROR  = "ERROR"
)

const (
	TAG_STATU_NONE  = "NONE"
	TAG_STATU_READ  = "READ"
	TAG_STATU_STOP  = "STOP"
	TAG_STATU_ERROR = "ERROR"
)

const (
	RECY_ITEM_STATU_NONE    = "NONE"
	RECY_ITEM_STATU_PLASTIC = "PLASTIC"
	RECY_ITEM_STATU_GLASS   = "GLASS"
	RECY_ITEM_STATU_METAL   = "METAL"
)

const (
	NFC_STATU_NONE = "NONE"
)

const (
	ULT_STATU_NONE = "NONE"
)

const (
	RFID_READER_OKBIT    = "5379"
	RFID_READER_STARTBIT = "4354"
	RFID_READER_CHECKBIT = "45"
	RFID_TAG_PATTERN     = "AFA09012018AFA"
)

const (
	DATATYPE_NONE                 = "NONE"
	DATATYPE_CUSTOMER             = "CUSTOMER"
	DATATYPE_USER                 = "USER"
	DATATYPE_CUSTOMERCONFIG       = "CUSTOMERCONFIG"
	DATATYPE_ADMINCONFIG          = "ADMINCONFIG"
	DATATYPE_LOCALCONFIG          = "LOCALCONFIG"
	DATATYPE_RECY_MAIN_DEVICE     = "RECY_MAIN_DEVICE"
	DATATYPE_RECY_BASE_DEVICE     = "RECY_BASE_DEVICE"
	DATATYPE_RECY_GPS_DEVICE      = "RECY_GPS_DEVICE"
	DATATYPE_RECY_THERM_DEVICE    = "RECY_THERM_DEVICE"
	DATATYPE_RECY_VERSION_DEVICE  = "RECY_VERSION_DEVICE"
	DATATYPE_RECY_ALARM_DEVICE    = "RECY_ALARM_DEVICE"
	DATATYPE_RECY_STATU_DEVICE    = "RECY_STATU_DEVICE"
	DATATYPE_RECY_DETAIL_DEVICE   = "RECY_DETAIL_DEVICE"
	DATATYPE_ULT_MAIN_DEVICE      = "ULT_MAIN_DEVICE"
	DATATYPE_ULT_BASE_DEVICE      = "ULT_BASE_DEVICE"
	DATATYPE_ULT_BATTERY_DEVICE   = "ULT_BATTERY_DEVICE"
	DATATYPE_ULT_GPS_DEVICE       = "ULT_GPS_DEVICE"
	DATATYPE_ULT_ALARM_DEVICE     = "ULT_ALARM_DEVICE"
	DATATYPE_ULT_THERM_DEVICE     = "ULT_THERM_DEVICE"
	DATATYPE_ULT_VERSION_DEVICE   = "ULT_VERSION_DEVICE"
	DATATYPE_ULT_STATU_DEVICE     = "ULT_STATU_DEVICE"
	DATATYPE_ULT_SENS_DEVICE      = "ULT_SENS_DEVICE"
	DATATYPE_RFID_MAIN_DEVICE     = "RFID_MAIN_DEVICE"
	DATATYPE_RFID_GPS_DEVICE      = "RFID_GPS_DEVICE"
	DATATYPE_RFID_ALARM_DEVICE    = "RFID_ALARM_DEVICE"
	DATATYPE_RFID_THERM_DEVICE    = "RFID_THERM_DEVICE"
	DATATYPE_RFID_VERSION_DEVICE  = "RFID_VERSION_DEVICE"
	DATATYPE_RFID_STATU_DEVICE    = "RFID_STATU_DEVICE"
	DATATYPE_RFID_BASE_DEVICE     = "RFID_BASE_DEVICE"
	DATATYPE_RFID_DETAIL_DEVICE   = "RFID_DETAIL_DEVICE"
	DATATYPE_RFID_WORKHOUR_DEVICE = "RFID_WORKHOUR_DEVICE"
	DATATYPE_TAG_MAIN             = "TAG_MAIN"
	DATATYPE_TAG_BASE             = "TAG_BASE"
	DATATYPE_TAG_STATU            = "TAG_STATU"
	DATATYPE_TAG_GPS              = "TAG_GPS"
	DATATYPE_TAG_READER           = "TAG_READER"
	DATATYPE_NFC_MAIN             = "NFC_MAIN"
	DATATYPE_NFC_BASE             = "NFC_BASE"
	DATATYPE_NFC_STATU            = "NFC_STATU"
	DATATYPE_NFC_READER           = "NFC_READER"
)

const (
	ALARMTYPE_NONE = "NONE"
	ALARMTYPE_ATMP = "THERM"
	ALARMTYPE_AGPS = "GPS"
)

const (
	ALARMSTATU_NONE  = "NONE"
	ALARMSTATU_ALARM = "ALARM"
)
const (
	THERMSTATU_NONE   = "NONE"
	THERMSTATU_NORMAL = "NORMAL"
	THERMSTATU_HIGH   = "HIGH"
)

const (
	BATTERYSTATU_NONE   = "NONE"
	BATTERYSTATU_NORMAL = "NORMAL"
	BATTERYSTATU_LOW    = "LOW"
)

const (
	REDIS_CUSTOMER_CHANNEL        = "ch-customer-"
	REDIS_APP_LOG_CHANNEL         = "ch-app-log"
	REDIS_SERIAL_RFID_DEVICE      = "serial-rfid-device"
	REDIS_SERIAL_ULT_DEVICE       = "serial-ult-device"
	REDIS_SERIAL_RECY_DEVICE      = "serial-recy-device"
	REDIS_RECY_MAIN_DEVICES       = "recy-main-devices"
	REDIS_RECY_BASE_DEVICES       = "recy-base-devices"
	REDIS_RECY_GPS_DEVICES        = "recy-gps-devices"
	REDIS_RECY_THERM_DEVICES      = "recy-therm-devices"
	REDIS_RECY_VERSION_DEVICES    = "recy-version-devices"
	REDIS_RECY_ALARM_DEVICES      = "recy-alarm-devices"
	REDIS_RECY_STATU_DEVICES      = "recy-statu-devices"
	REDIS_RECY_DETAIL_DEVICES     = "recy-detail-devices"
	REDIS_ULT_MAIN_DEVICES        = "ult-main-devices"
	REDIS_ULT_BASE_DEVICES        = "ult-base-devices"
	REDIS_ULT_BATTERY_DEVICES     = "ult-battery-devices"
	REDIS_ULT_GPS_DEVICES         = "ult-gps-devices"
	REDIS_ULT_ALARM_DEVICES       = "ult-alarm-devices"
	REDIS_ULT_THERM_DEVICES       = "ult-therm-devices"
	REDIS_ULT_VERSION_DEVICES     = "ult-version-devices"
	REDIS_ULT_STATU_DEVICES       = "ult-statu-devices"
	REDIS_ULT_SENS_DEVICES        = "ult-sens-devices"
	REDIS_RFID_MAIN_DEVICES       = "rfid-main-devices"
	REDIS_RFID_GPS_DEVICES        = "rfid-gps-devices"
	REDIS_RFID_ALARM_DEVICES      = "rfid-alarm-devices"
	REDIS_RFID_THERM_DEVICES      = "rfid-therm-devices"
	REDIS_RFID_VERSION_DEVICES    = "rfid-version-devices"
	REDIS_RFID_STATU_DEVICES      = "rfid-statu-devices"
	REDIS_RFID_BASE_DEVICES       = "rfid-base-devices"
	REDIS_RFID_DETAIL_DEVICES     = "rfid-detail-devices"
	REDIS_RFID_WORKHOUR_DEVICES   = "rfid-workhour-devices"
	REDIS_TAG_EPC                 = "tag-epc"
	REDIS_TAG_MAINS               = "tag-mains"
	REDIS_TAG_BASES               = "tag-bases"
	REDIS_TAG_STATUS              = "tag-status"
	REDIS_TAG_GPSES               = "tag-gpses"
	REDIS_TAG_READERS             = "tag-readers"
	REDIS_NFC_EPC                 = "nfc-epc"
	REDIS_NFC_MAINS               = "nfc-mains"
	REDIS_NFC_BASES               = "nfc-bases"
	REDIS_NFC_STATUS              = "nfc-status"
	REDIS_NFC_READERS             = "nfc-readers"
	REDIS_CUSTOMER_TAGVIEWS       = "customer-tagviews"
	REDIS_CUSTOMER_TAGVIEWS_REEL  = "customer-tagviews-reel"
	REDIS_CUSTOMER_TAGS           = "customer-tags"
	REDIS_CUSTOMER_NFCS           = "customer-nfcs"
	REDIS_CUSTOMER_RFID_DEVICES   = "customer-rfid-devices"
	REDIS_CUSTOMER_RECY_DEVICES   = "customer-recy-devices"
	REDIS_CUSTOMER_ULT_DEVICES    = "customer-ult-devices"
	REDIS_CUSTOMERS               = "customers"
	REDIS_CUSTOMER_CUSTOMERCONFIG = "customer-customerconfig"
	REDIS_CUSTOMER_ADMINCONFIG    = "customer-adminconfig"
	REDIS_CUSTOMER_LOCALCONFIG    = "customer-localconfig"
	REDIS_CUSTOMER_LINK           = "customer-link"
	REDIS_USERS                   = "users"
	REDIS_CUSTOMER_USERS          = "customer-users"
	REDIS_USER_TOKENENDDATE       = "user-tokenenddate"
	REDIS_USER_TOKEN              = "user-token"
	REDIS_SERIAL_ALARM            = "serial-alarm"
	REDIS_HASHKEY                 = "HASHKEY"
	REDIS_HASHBASEKEY             = "HASHBASEKEY"
	REDIS_CHANNELKEY              = "CHANNELKEY"
	REDIS_SUBKEY                  = "SUBKEY"
	REDIS_KEYVALUE                = "KEYVALUE"
)

const (
	USER_ROLE_GUEST  = "GUEST"
	USER_ROLE_ADMIN  = "ADMIN"
	USER_ROLE_REPORT = "REPORT"
)

const (
	CONTAINERTYPE_NONE = "NONE"
)

const (
	TRUCKTYPE_NONE = "NONE"
)

const (
	DEVICETYPE_NONE = "NONE"
	DEVICETYPE_RFID = "RFID"
	DEVICETYPE_ULT  = "ULT"
	DEVICETYPE_RECY = "RECY"
)

const (
	APPTYPE_NONE = "NONE"
)

const (
	RFID_DEVICE_TYPE_NONE = "NONE"
)

const (
	ULT_DEVICE_TYPE_NONE = "NONE"
)

const (
	RECY_DEVICE_TYPE_NONE = "NONE"
)

const (
	RESULT_ERROR_HTTP_PARSE                       = "HTTP_PARSE"
	RESULT_ERROR_HTTP_POST                        = "HTTP_POST"
	RESULT_ERROR_USER_INVALIDTOKEN                = "USER_INVALIDTOKEN"
	RESULT_ERROR_USER_ENDTOKEN                    = "USER_ENDTOKEN"
	RESULT_ERROR_USER_INVALIDUSER                 = "USER_INVALIDUSER"
	RESULT_ERROR_USER_INVALIDUSERROLE             = "USER_INVALIDUSERROLE"
	RESULT_ERROR_USER_INVALIDPASSWORD             = "USER_INVALIDPASSWORD"
	RESULT_ERROR_USER_USERNAMEEXIST               = "USER_USERNAMEEXIST"
	RESULT_ERROR_USER_USEREMAILEXIST              = "USER_USEREMAILEXIST"
	RESULT_ERROR_USER_PASSWORDEMPTY               = "USER_PASSWORDEMPTY"
	RESULT_ERROR_READERTYPE                       = "READERTYPE"
	RESULT_ERROR_DEVICETYPE                       = "DEVICETYPE"
	RESULT_ERROR_DATATYPE                         = "DATATYPE"
	RESULT_ERROR_CUSTOMER_NOTFOUND                = "CUSTOMER_NOTFOUND"
	RESULT_ERROR_USER_NOTFOUND                    = "USER_NOTFOUND"
	RESULT_ERROR_DEVICE_NOTFOUND                  = "DEVICE_NOTFOUND"
	RESULT_ERROR_DB_SAVE                          = "DB_SAVE"
	RESULT_ERROR_REDIS_SAVE                       = "REDIS_SAVE"
	RESULT_ERROR_DB_GET                           = "DB_GET"
	RESULT_ERROR_REDIS_GET                        = "REDIS_GET"
	RESULT_ERROR_CUSTOMER_INVALID                 = "CUSTOMER_INVALID"
	RESULT_ERROR_CUSTOMER_CUSTOMERCONFIG_NOTFOUND = "CUSTOMER_CUSTOMERCONFIG_NOTFOUND"
	RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND    = "CUSTOMER_ADMINCONFIG_NOTFOUND"
	RESULT_ERROR_CUSTOMER_LOCALCONFIG_NOTFOUND    = "CUSTOMER_LOCALCONFIG_NOTFOUND"
	RESULT_ERROR_CUSTOMER_TAGS_NOTFOUND           = "CUSTOMER_TAGS_NOTFOUND"
	RESULT_ERROR_CUSTOMER_NFCS_NOTFOUND           = "CUSTOMER_NFCS_NOTFOUND"
	RESULT_ERROR_CUSTOMERS_NOTFOUND               = "CUSTOMERS_NOTFOUND"
	RESULT_ERROR_DEVICES_NOTFOUND                 = "DEVICES_NOTFOUND"
	RESULT_ERROR_CUSTOMERS_DEVICES_NOTFOUND       = "CUSTOMERS_DEVICES_NOTFOUND"
	RESULT_ERROR_TAG_NOTFOUND                     = "TAG_NOTFOUND"
	RESULT_ERROR_TAG_CUSTOMER_NOTFOUND            = "TAG_CUSTOMER_NOTFOUND"
	RESULT_ERROR_NFC_NOTFOUND                     = "NFC_NOTFOUND"
	RESULT_ERROR_NFC_CUSTOMER_NOTFOUND            = "NFC_CUSTOMER_NOTFOUND"
	RESULT_ERROR_APP_STARTED                      = "APP_STARTED"
	RESULT_ERROR_IGNORE_FIRST_DATA                = "IGNORE_FIRST_DATA"
)

const (
	WEB_APP_TYPE_RFID = "RfIdApp"
	WEB_APP_TYPE_ULT  = "UltApp"
	WEB_APP_TYPE_RECY = "RecyApp"
)

const (
	RECY_SOCKET_INDEX   = "INDEX"
	RECY_SOCKET_ANALYZE = "ANALYZE"
	RECY_SOCKET_FINISH  = "FINISH"
)
