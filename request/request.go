package request

import (
	"strings"

	"github.com/labstack/echo"
)

var (
	requestsMap = map[string]func(c echo.Context) (interface{}, map[string]string){
		"authRegister": validateRegisterRequest,
	}
)

func ValidateRequest(validator string, c echo.Context) (interface{}, map[string]string) {
	return requestsMap[validator](c)
}

func formatRequestErrors(err error) map[string]string {
	splt1 := strings.Split(strings.TrimSuffix(err.Error(), ";"), ";")
	reqErrors := make(map[string]string, len(splt1))

	for i := 0; i < len(splt1); i++ {
		if splt1[i] == "" {
			continue
		}

		splt2 := strings.Split(splt1[i], ":")
		for j := 0; j < len(splt2); j++ {
			reqErrors[splt2[0]] = strings.TrimSpace(splt2[1])
		}
	}

	return reqErrors
}
