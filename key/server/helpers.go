package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//ebsClien extracts links of provided URL
func request(buf []byte, url string) (EbsResponse, error) {
	verifyTLS := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ebsClient := http.Client{
		Timeout:   5 * time.Second,
		Transport: verifyTLS,
	}

	log.Printf("The sent request is: %v\n\n", string(buf))
	// url := ip + "/" + endpoint
	reqBuilder, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(buf))

	reqBuilder.Header.Add("content-type", "application/json")
	res, err := ebsClient.Do(reqBuilder)
	if err != nil {
		log.Printf("The error is: %v", err)
		return EbsResponse{}, errors.New("it doesn't work")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("The error is: %v", err)
		return EbsResponse{}, errors.New("it doesn't work")
	}
	defer res.Body.Close()

	var result EbsResponse
	if res.StatusCode == 200 {
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Printf("The error is: %v", err)
			return EbsResponse{}, errors.New("it doesn't work")
		}
		log.Printf("the inner response is: %v", result)
		return result, nil
	}

	return EbsResponse{}, fmt.Errorf("%s", body)

}

// EBSResponse struct template
// something should solve it
type EbsResponse struct {
	WorkingKey      string `json:"workingKey"`
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"ResponseMessage"`
}

type EbsRequest struct {
	STAN         int    `json:"systemTraceAuditNumber"`
	TranDateTime string `json:"tranDateTime"`
	ClientID     string `json:"clientId"`
	TerminalID   string `json:"terminalId"`
}

type PurchaseRequest struct {
	STAN         int    `json:"systemTraceAuditNumber"`
	TranDateTime string `json:"tranDateTime"`
	ClientID     string `json:"clientId"`
	TerminalID   string `json:"terminalId"`
	CardFields
	AmountFields
}

type CardFields struct {
	Pan     string `json:"PAN"`
	Expdate string `json:"expDate"`
	Pin     string `json:"PIN"`
}

type AmountFields struct {
	Amount   float32 `json:"tranAmount"`
	Currency string  `json:"tranCurrencyCode"`
}
