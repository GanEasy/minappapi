package minappapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//httpClient 默认http.Client
var httpClient *http.Client

func init() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	httpClient = &client
}

// HTTPGetJSON 通过传入url和结构，提取出页面中的值
func HTTPGetJSON(url string, response interface{}) error {
	httpResp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return DecodeJSONHttpResponse(httpResp.Body, response)
}

//DecodeJSONHttpResponse 解决json
func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
