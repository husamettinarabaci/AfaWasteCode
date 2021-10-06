package WasteLibrary

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

//CheckAuth
func CheckAuth(data url.Values, customerId string, userRole string) ResultType {
	var resultVal ResultType
	resultVal.Result = RESULT_FAIL
	data.Add(HTTP_CUSTOMERID, customerId)
	data.Add(HTTP_USERROLE, userRole)
	resultVal = HttpPostReq("http://waste-authapi-cluster-ip/checkAuth", data)
	return resultVal
}

//GenerateToken
func GenerateToken(tokenVal string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(tokenVal), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	return base64.StdEncoding.EncodeToString(hash)
}
