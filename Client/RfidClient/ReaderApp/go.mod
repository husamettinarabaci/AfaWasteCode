module afatek.com.tr/devafatek/rfid/ReaderApp

go 1.17

replace github.com/devafatek/WasteLibrary => ../../../Host/WasteLibrary

require (
	gitee.com/wiseai/go-rpio v4.2.0+incompatible
	github.com/AfatekDevelopers/serial_lib_go v0.0.0-20200525185650-8665739e675a
	github.com/devafatek/WasteLibrary v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.4 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55 // indirect
)
