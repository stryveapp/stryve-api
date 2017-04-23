package request

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
)

type RegisterRequest struct {
	Username string `valid:"length(3|25)~username:Username must be between 3 and 25 characters long,username~username:Username can only contain alpanumeric and underscore characters,required~username:Username is required,unique_username~username:Username is already taken"`
	Email    string `valid:"email~email:Invalid email address,required~email:Email is required,unique_email~email:Email address is already registered"`
	Password string `valid:"length(8|99)~password:Password must be a minimum of 8 characters long,password~password:Password must contain both numbers and letters,required~password:Password is required"`
}

func validateRegisterRequest(c echo.Context) (interface{}, map[string]string) {
	var req RegisterRequest

	c.Bind(&req)
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return nil, formatRequestErrors(err)
	}

	return req, map[string]string{}
}

func validateLoginRequest(c echo.Context) map[string]string {
	// TODO
	return map[string]string{}
}
