package main

import (
	"context"
	"net/http"

	"github.com/devafatek/WasteLibrary"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var redisDb *redis.Client
var ctx = context.Background()

var currentCustomerList WasteLibrary.CustomersType

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "waste-redis-master-cluster-ip:6379",
		Password: "Amca151200!Furkan",
		DB:       0,
	})

	pong, err := redisDb.Ping(ctx).Result()
	WasteLibrary.LogErr(err)
	WasteLibrary.LogStr(pong)
	currentCustomerList.New()

}

func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/socket", socket)
	http.ListenAndServe(":80", nil)
}

var upgrader = websocket.Upgrader{}

func socket(w http.ResponseWriter, req *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,access-control-allow-origin, access-control-allow-headers")
	}

	var linkCustomer WasteLibrary.CustomerType
	resultVal := linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		WasteLibrary.LogErr(err)
		return
	}
	defer c.Close()

	subscriber := redisDb.Subscribe(ctx, WasteLibrary.REDIS_CUSTOMER_CHANNEL+linkCustomer.ToIdString())

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			//WasteLibrary.LogErr(err)
			break
		}
		err = c.WriteMessage(1, []byte(msg.Payload))
		if err != nil {
			//WasteLibrary.LogErr(err)
			break
		}
	}
}
