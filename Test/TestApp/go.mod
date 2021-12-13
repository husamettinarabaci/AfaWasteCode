module afatek.com.tr/devafatek/test/TestApp

go 1.17

replace github.com/devafatek/WasteLibrary => ../../Host/WasteLibrary

require (
	github.com/devafatek/WasteLibrary v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.4
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
)
