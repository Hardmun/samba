package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type jsonStruct struct {
	Uuid   string `json:"uuid"`
	Number string `json:"number"`
	Date   string `json:"date"`
	Rk_id  string `json:"rk_id"`
	Amount int32  `json:"amount"`
}

func main() {
	newRequest := jsonStruct{
		Uuid:   "805da498-83f7-11ef-816c-0050568a4702",
		Number: "TEST PO #11",
		Date:   "20241008",
		Rk_id:  "У191994",
		Amount: 15000,
	}

	jsonRequest, err := json.Marshal(newRequest)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	sendRequest, err := http.NewRequest(http.MethodPut, "https://1cplatform.mitcoms.ru/Dev/U_MediaFinance_for_testing/hs/samba/po",
		bytes.NewBuffer(jsonRequest))
	if err != nil {
		fmt.Errorf(err.Error())
	}
	sendRequest.SetBasicAuth("user1c", "123")
	sendRequest.Header.Set("Content-Type", "application/json")

	client1c := &http.Client{}

	resp, respErr := client1c.Do(sendRequest)
	if respErr != nil {
		fmt.Errorf(respErr.Error())
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)

	jFinal, errR := io.ReadAll(resp.Body)
	if errR != nil {
		fmt.Errorf(errR.Error())
	}

	var finalStruct interface{}
	errMarh := json.Unmarshal(jFinal, &finalStruct)
	if errMarh != nil{
		fmt.Errorf(errMarh.Error())
	}

	base64Data := finalStruct.(map[string]interface{})["result"].(map[string]interface{})["po_file"].(string)
	pdfData, pdfErr := base64.StdEncoding.DecodeString(base64Data)
	if pdfErr != nil{
		fmt.Errorf(pdfErr.Error())
	}
	os.WriteFile("./test.pdf", pdfData, os.ModePerm)
}
