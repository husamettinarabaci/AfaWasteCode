package WasteLibrary

import (
	"encoding/base64"
	"net/url"
	"strings"

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
func GenerateToken(tokenVal string, userId string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(tokenVal), bcrypt.DefaultCost)
	if err != nil {
		LogErr(err)
	}
	var lastHash string = userId + "#" + string(hash)
	return base64.StdEncoding.EncodeToString([]byte(lastHash))
}

//DecodeToken
func GetUserIdByToken(tokenVal string) string {
	lastHash, err := base64.StdEncoding.DecodeString(tokenVal)
	if err != nil {
		LogErr(err)
	}
	spData := strings.Split(string(lastHash), "#")
	userId := spData[0]

	return userId
}
