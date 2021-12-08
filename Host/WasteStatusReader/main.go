package main

import (
	"net/http"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))

	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceStatu.DeviceId = currentData.DeviceId
			currentData.DeviceStatu.StatusTime = currentHttpHeader.Time

			if currentData.DeviceStatu.ReaderAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ReaderConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ReaderStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ReaderLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.CamAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.CamConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.CamStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.CamLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.GpsAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.GpsConnStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsConnLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.GpsStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.GpsLastOkTime = currentHttpHeader.Time
			}

			if currentData.DeviceStatu.ThermAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ThermAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.TransferAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.TransferAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.SystemAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.SystemAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.UpdaterAppStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.UpdaterAppLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.AliveStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.AliveLastOkTime = currentHttpHeader.Time
			}
			if currentData.DeviceStatu.ContactStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.ContactLastOkTime = currentHttpHeader.Time
			}

			resultVal = currentData.DeviceStatu.SaveToDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.DeviceStatu.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.DeviceStatu.SaveToReaderDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_STATU, currentData.DeviceStatu.ToString())
		} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
			var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceStatu.DeviceId = currentData.DeviceId
			currentData.DeviceStatu.StatusTime = currentHttpHeader.Time

			ultCm := currentData.DeviceSens.UltRange1 * 173 / 10000

			var step1 float64 = 50
			var step2 float64 = 100
			var step3 float64 = 150
			var containerHeigth float64 = 200

			//TO DO
			//calculate ult status by container type
			var oldData WasteLibrary.UltDeviceType
			oldData.DeviceId = currentData.DeviceId
			oldData.GetByRedis("0")

			switch oldData.DeviceBase.ContainerType {
			case WasteLibrary.CONTAINERTYPE_NONE:
				containerHeigth = 200
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_400LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_400LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_770LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_770LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_1100LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_1100LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_BOYALI_2500LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_BOYALI_2500LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_BOYALI_3000LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_BOYALI_3000LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_BOYALI_3750LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_BOYALI_3750LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_PASLANMAZ_2200LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_PASLANMAZ_2200LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_PLASTIK_3000LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_PLASTIK_3000LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_PLASTIK_5000LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_PLASTIK_5000LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_PASLANMAZ_3500LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_PASLANMAZ_3500LT
			case WasteLibrary.CONTAINERTYPE_TEKSIN_GALVANIZ_PASLANMAZ_5000LT:
				containerHeigth = WasteLibrary.CONTAINERHEIGTH_TEKSIN_GALVANIZ_PASLANMAZ_5000LT
			default:
				containerHeigth = 200
			}

			step1 = (containerHeigth * 25) / 100
			step2 = (containerHeigth * 50) / 100
			step3 = (containerHeigth * 75) / 100

			if ultCm > (step3 + 50) {
				ultCm = step3 + 50
			}

			currentData.DeviceStatu.SensPercent = 100 - ((ultCm * 100) / (step3 + 50))

			if ultCm < step1 {
				currentData.DeviceStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_FULL
			}
			if ultCm >= step1 && ultCm < step2 {
				currentData.DeviceStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_MEDIUM
			}

			if ultCm >= step2 && ultCm < step3 {
				currentData.DeviceStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_LITTLE
			}

			if ultCm >= step3 {
				currentData.DeviceStatu.ContainerStatu = WasteLibrary.CONTAINER_FULLNESS_STATU_EMPTY
			}

			if currentData.DeviceStatu.AliveStatus == WasteLibrary.STATU_ACTIVE {
				currentData.DeviceStatu.AliveLastOkTime = currentHttpHeader.Time
			}

			resultVal = currentData.DeviceStatu.SaveToDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.DeviceStatu.SaveToRedis()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
				w.Write(resultVal.ToByte())

				return
			}

			resultVal = currentData.DeviceStatu.SaveToReaderDb()
			if resultVal.Result != WasteLibrary.RESULT_OK {
				resultVal.Result = WasteLibrary.RESULT_FAIL
				resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
				w.Write(resultVal.ToByte())

				return
			}
			WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentHttpHeader.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_STATU, currentData.DeviceStatu.ToString())
		}

	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
