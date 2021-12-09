module github.com/devafatek/WasteWebSocket

go 1.14

replace github.com/devafatek/WasteLibrary => ../WasteLibrary

require (
	github.com/devafatek/WasteLibrary v1.0.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gorilla/websocket v1.4.2

)
