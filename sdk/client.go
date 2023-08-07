package sdk

import (
	"errors"
	"fmt"
	"github.com/al-online/mzapi-client-sdk/pkg/jwt"
	"github.com/al-online/mzapi-client-sdk/sdk/request"
	"log"
	"net/http"
	"runtime"
	"strings"
)

const Version = "1.0.0"

var defaultUserAgent = fmt.Sprintf("MEDUSDKGo/%s (%s; %s) Golang/%s", Version, runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"))

// Client is the base struct of service clients
type Client struct {
	Config     *Config
	Credential *Credential
}

type SignFunc func(*http.Request) error

// Send 每个请求都需要通过Send
func (c *Client) Send(req request.Request) (*http.Response, error) {
	//获取请求方法
	method := req.GetMethod()
	body := req.GetBody()
	reqUrl := req.GetURL()
	//签名，将token添加进请求头，这一步可以封装成一个signer使用不同的加密方法，这里为了方便就只使用了jwt
	sign := func(r *http.Request) error {
		if c.Credential.AccessKey == "" {
			return errors.New("accessKey不能为空")
		}
		tokenString, err := jwt.GenToken(c.Credential.AccessKey, c.Credential.SecretKey, string(body))
		r.Header.Set("Accesskey", c.Credential.AccessKey)
		r.Header.Set("Authorization", tokenString)
		return err
	}
	return c.doSend(method, reqUrl, string(body), req.GetHeaders(), sign)
}

func (c *Client) WithSecret(secretID, secretKey string) *Client {
	c.Credential = NewCredentials(secretID, secretKey)
	return c
}

func (c *Client) WithCredential(cred *Credential) *Client {
	c.Credential = cred
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	return c
}

func (c *Client) doSend(method, url, data string, header map[string]string, sign SignFunc) (*http.Response, error) {
	client := &http.Client{Timeout: c.Config.Timeout}

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	c.setHeader(req, header)
	//签名,并添加到请求头
	err = sign(req)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(req)
	return response, err
}

func (c *Client) setHeader(req *http.Request, header map[string]string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", defaultUserAgent)

	for k, v := range header {
		req.Header.Set(k, v)
	}
}
