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
	ResultSendOTPRegister struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		ErrorDescription string `json:"errorDescription" `
		DeveloperMessage string `json:"developerMessage" `
		TimeStamp        int    `json:"timeStamp" `
		Result           struct {
			RefCode string `json:"refCode"`
		} `json:"result" `
		Message string `json:"message" `
	}

	ResultValidateOTPByPhone struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		ErrorDescription string `json:"errorDescription" `
		DeveloperMessage string `json:"developerMessage" `
		TimeStamp        int    `json:"timeStamp" `
		Result           string `json:"result" `
		Message          string `json:"message" `
	}

	ResultValidateBeforeRegister struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		ErrorDescription string `json:"errorDescription" `
		DeveloperMessage string `json:"developerMessage" `
		TimeStamp        int    `json:"timeStamp" `
		Result           struct {
			PersonalID string `json:"personalId"`
			Email      string `json:"email"`
		} `json:"result" `
		Message string `json:"message" `
	}

	ResultRegister struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		ErrorDescription string `json:"errorDescription" `
		DeveloperMessage string `json:"developerMessage" `
		TimeStamp        int    `json:"timeStamp" `
		Result           struct {
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			MobilePhoneNo   string `json:"mobilePhoneNo"`
			Email           string `json:"email"`
			Address         string `json:"address"`
			PersonalID      string `json:"personalId"`
			AvatarImageLink string `json:"avatarImageLink"`
			MemberLevel     int    `json:"memberLevel"`
			BirthDate       string `json:"birthDate"`
			UsageLimit      string `json:"usageLimit"`
			MemberCode      string `json:"memberCode"`
			Wallet          struct {
				AccountNumber    string  `json:"accountNumber"`
				BalanceAmount    float32 `json:"balanceAmount"`
				UsedAmount       float32 `json:"usedAmount"`
				BalanceAmountStr string  `json:"balanceAmountStr"`
				UsedAmountStr    string  `json:"usedAmountStr"`
				FileUrl          string  `json:"fileUrl"`
			} `json:"wallet" `
			FullName string `json:"fullName"`
			FileUrl  string `json:"fileUrl"`
			Gender   string `json:"gender"`
			Country  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"country" `
			Job               string `json:"job"`
			SubJob            string `json:"subJob"`
			JobDescription    string `json:"jobDescription"`
			SubJobDescription string `json:"subJobDescription"`
			Province          string `json:"province"`
			ProvinceID        string `json:"provinceId"`
			Postcode          string `json:"postcode"`
			MsgUnread         string `json:"msgUnread"`
		} `json:"result" `
		Message string `json:"message" `
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

	ReqSendReqToCore struct {
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		MobilePhoneNo string `json:"mobilePhoneNo"`
		Email         string `json:"email"`
		Password      string `json:"password"`
		PersonalID    string `json:"personalID"`
		RefCode       string `json:"refCode"`
		CountryID     string `json:"countryId"`
	}

	ResSendReqToCore struct {
		Success          bool   `json:"success"`
		ResultCode       string `json:"resultCode" `
		ErrorDescription string `json:"errorDescription" `
		DeveloperMessage string `json:"developerMessage" `
		TimeStamp        int    `json:"timeStamp" `
		Result           struct {
			FirstName     string `json:"firstName"`
			LastName      string `json:"lastName"`
			MobilePhoneNo string `json:"mobilePhoneNo"`
			Email         string `json:"email"`
			Password      string `json:"password"`
			PersonalID    string `json:"personalID"`
			RefCode       string `json:"refCode"`
			CountryID     string `json:"countryId"`
		} `json:"result" `
		Message string `json:"message" `
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
	var data ResultSendOTPRegister

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
	var data ResultValidateOTPByPhone
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
	var data ResultValidateBeforeRegister
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
	var data ResultRegister
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, data)
}

func SendReqToCore(c echo.Context) (err error) {

	var reqBody ReqSendReqToCore

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
	var data ResultValidateOTPByPhone
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, data)

}
