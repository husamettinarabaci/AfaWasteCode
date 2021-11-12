module github.com/devafatek/WasteStoreApiForReaderDB

go 1.14

replace github.com/devafatek/WasteLibrary => ../WasteLibrary

require (
	github.com/AfatekDevelopers/result_lib_go v0.0.0-20210831140827-985022b03085
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/devafatek/WasteLibrary v1.0.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/lib/pq v1.10.2

)
