package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

const (
	FONDY_CHEKCOUT_URL      = "https://pay.fondy.eu/api/checkout/url"
	FONDY_MERCHANT_ID       = "1396424"
	FONDY_MERCHANT_PASSWORD = "test"
)

func main() {
	id := uuid.New()
	// Example usage
	checkoutRequest := CheckoutRequest{
		OrderID:           id.String(),
		MerchantID:        FONDY_MERCHANT_ID,
		OrderDesc:         "Test Order",
		Amount:            "1000",
		Currency:          "KZT",
		ServerCallbackURL: "http://188.227.35.6:8081",
	}

	checkoutRequest.SetSignature(FONDY_MERCHANT_PASSWORD)

	request := APIRequest{
		Request: checkoutRequest}

	requestBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return
	}
	fmt.Println(checkoutRequest)
	resp, err := http.Post(FONDY_CHEKCOUT_URL, "application/json", strings.NewReader(string(requestBody)))

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(body)
	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Response from server:", apiResponse.Response)
}

type APIResponse struct {
	Response interface{} `json:"response"`
}

type APIRequest struct {
	Request interface{} `json:"request"`
}

type CheckoutRequest struct {
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	OrderDesc         string `json:"order_desc"`
	Signature         string `json:"signature"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	ResponseURL       string `json:"response_url"`
	ServerCallbackURL string `json:"server_callback_url"`
	SenderEmail       string `json:"sender_email"`
	Language          string `json:"language"`
	ProductID         string `json:"product_id"`
}

func (r *CheckoutRequest) SetSignature(password string) {

	params := structs.Map(r)

	var key []string
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	values := []string{}

	for _, v := range key {
		value := params[v].(string)
		if value != "" {
			values = append(values, value)
		}
	}

	r.Signature = generateSignature(password, values)
}

func generateSignature(password string, values []string) string {
	newValues := []string{password}
	newValues = append(newValues, values...)
	signatureString := strings.Join(newValues, "|")
	fmt.Println("Signature String:", signatureString)
	hash := sha1.New()
	hash.Write([]byte(signatureString))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
