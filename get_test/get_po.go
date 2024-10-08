package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var (
		err      error
		request  *http.Request
		resp     *http.Response
		readBody []byte
	)

	baseURL := "https://1cplatform.mitcoms.ru/Dev/U_MediaFinance_for_testing/hs/samba/files/po"

	urlParams := url.Values{}
	urlParams.Add("uuid", "860c7845-8566-11ef-816c-0050568a4702")
	fullURL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())

	request, err = http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.SetBasicAuth("user1c", "123")
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err = client.Do(request)
	if err != nil {
		fmt.Println(resp)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp)
		return
	}

	readBody, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var respJson any
	err = json.Unmarshal(readBody, &respJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bs64 := respJson.(map[string]interface{})["result"].(string)
	readBody, err = base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	os.WriteFile("./test_get.pdf", readBody, os.ModePerm)
}
