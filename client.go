package minappapi

import (
	"bytes"
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

//HTTPPostJSON  通过传入url和内容，提交内容后，提取出页面中的值
func HTTPPostJSON(url string, body []byte, response interface{}) error {
	httpResp, err := httpClient.Post(url, "application/json; charset=utf-8", bytes.NewReader(body))
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
