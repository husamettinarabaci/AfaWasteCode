module afatek.com.tr/devafatek/TestApp

go 1.17

replace github.com/devafatek/WasteLibrary => ../../Host/WasteLibrary

require github.com/devafatek/WasteLibrary v0.0.0-00010101000000-000000000000

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)
