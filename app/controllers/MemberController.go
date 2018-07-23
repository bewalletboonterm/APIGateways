package MemberController

import (
	// "APIGateways/app/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type (
	ResultEntity struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		Path             string `json:"path,omitempty" `
		ErrorDescription string `json:"errorDescription,omitempty" `
		DeveloperMessage string `json:"developerMessage,omitempty" `
		TimeStamp        int    `json:"timeStamp" `
		Result           struct {
			RefCode string `json:"refCode,omitempty"`
			Rest    string `json:"rest,omitempty"`
		} `json:"result" `
		ServerError string `json:"serverError,omitempty" `
		Method      string `json:"method,omitempty" `
		Header      string `json:"header,omitempty" `
		RequestBody string `json:"requestBody,omitempty" `
		Message     string `json:"message,omitempty" `
	}

	ReqSendOTPRegister struct {
		MobilePhoneNo string `json:"mobilePhoneNo"`
	}

	ReqValidateOTPByPhone struct {
		MobilePhoneNo string `json:"mobilePhoneNo"`
		RefCode       string `json:"refCode"`
		OtpCode       string `json:"otpCode"`
	}

	ReqValidateBeforeRegister struct {
		PersonalID string `json:"personalId"`
		Email      string `json:"email"`
	}

	ReqRegister struct {
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		MobilePhoneNo string `json:"mobilePhoneNo"`
		Email         string `json:"email"`
		Password      string `json:"password"`
		PersonalID    string `json:"personalID"`
		RefCode       string `json:"refCode"`
		OtpCode       string `json:"otpCode"`
		CountryID     string `json:"countryId"`
	}
)

// Send otp to user
func SendOTPRegister(c echo.Context) (err error) {

	var reqBody ReqSendOTPRegister
	// reqBody := echo.Map{}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	// Build the request
	req, err := http.NewRequest("GET", os.Getenv("HOST_BEWALLET")+"/rest/APIGateway/sendOTPRegister?mobilePhoneNo="+reqBody.MobilePhoneNo, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the data with the data from the JSON
	var data ResultEntity

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	// fmt.Println(data.ResultCode);
	return c.JSON(http.StatusOK, data)

}

func ValidateOTPByPhone(c echo.Context) (err error) {

	var reqBody ReqValidateOTPByPhone

	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	jsonValue, _ := json.Marshal(reqBody)
	request, _ := http.NewRequest("POST", os.Getenv("HOST_BEWALLET")+"/rest/APIGateway/validateOTPByPhone", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer resp.Body.Close()
	var data ResultEntity
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, data)

}

func ValidateBeforeRegister(c echo.Context) (err error) {

	var reqBody ReqValidateBeforeRegister

	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	jsonValue, _ := json.Marshal(reqBody)
	request, _ := http.NewRequest("POST", os.Getenv("HOST_BEWALLET")+"/rest/APIGateway/validateBeforeRegister", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer resp.Body.Close()
	var data ResultEntity
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, data)

}

func Register(c echo.Context) error {

	var reqBody ReqRegister

	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	jsonValue, _ := json.Marshal(reqBody)
	request, _ := http.NewRequest("POST", os.Getenv("HOST_BEWALLET")+"/rest/APIGateway/register", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer resp.Body.Close()
	var data ResultEntity
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, data)
}
