package services

import (
	"encoding/json"
	"io"
	"mzapi-sdk/sdk"
	"mzapi-sdk/sdk/request"
	"mzapi-sdk/services/common/models"
	"net/http"
)

type Client struct {
	sdk.Client
}

// NewClient 自定义配置
func NewClient(config *sdk.Config, credential *sdk.Credential) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		config = sdk.NewConfig()
	}
	client.WithCredential(credential).WithConfig(config)
	return
}

// NewClientWithSecret 用两个key初始化
func NewClientWithSecret(accessKey, secretKey string) (client *Client, err error) {
	client = &Client{}
	config := sdk.NewConfig()
	client.WithSecret(accessKey, secretKey).WithConfig(config)
	return
}

func (c *Client) GetNameByGETMethod(name string) (string, error) {
	var err error
	url := "http://localhost:9123/api/name?name=" + name
	req := &request.BaseRequest{
		URL:     url,
		Method:  "GET",
		Header:  nil,
		Version: "v1",
		Body:    []byte{},
	}
	resp := &http.Response{}
	resp, err = c.Send(req)
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	buf, _ := io.ReadAll(body)
	return string(buf), err
}

func (c *Client) GetNameByPOSTMethod(name string) (string, error) {
	var err error
	url := "http://localhost:9123/api/name?name=" + name
	req := &request.BaseRequest{
		URL:     url,
		Method:  "POST",
		Header:  nil,
		Version: "v1",
		Body:    []byte{},
	}
	resp := &http.Response{}
	resp, err = c.Send(req)
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	buf, _ := io.ReadAll(body)
	return string(buf), err
}

func (c *Client) GetNameByJSON(user models.User) (string, error) {
	var err error
	url := `http://localhost:9123/api/json`
	requestBody, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	req := &request.BaseRequest{
		URL:     url,
		Method:  "POST",
		Header:  nil,
		Version: "v1",
		Body:    requestBody,
	}
	resp := &http.Response{}
	resp, err = c.Send(req)
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	buf, _ := io.ReadAll(body)
	return string(buf), err
}
