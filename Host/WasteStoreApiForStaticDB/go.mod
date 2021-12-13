module github.com/devafatek/WasteStoreApiForStaticDB

go 1.14

replace github.com/devafatek/WasteLibrary => ../WasteLibrary

require (
	github.com/devafatek/WasteLibrary v1.0.0
	github.com/lib/pq v1.10.2
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect

)
