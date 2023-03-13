package jumpserver

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	gourl "net/url"
	"time"

	"golang.org/x/time/rate"
)

var (
	defaultHttpClient = &http.Client{
		Timeout:   30 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	defaultRateLimiter = rate.NewLimiter(rate.Every(500*time.Millisecond), 10)
)

type Client struct {
	jmsUrl    string
	accessKey string
	secretKey string

	limiter    *rate.Limiter
	httpClient *http.Client

	sigAuth *SigAuth

	Logger
	StructuredLogger
}

// New
func New(jmsUrl string, opts ...Option) (*Client, error) {
	jms := &Client{
		jmsUrl: jmsUrl,
	}

	for _, fn := range opts {
		if fn == nil {
			continue
		}
		if err := fn.apply(jms); err != nil {
			return nil, err
		}
	}
	if jms.limiter == nil {
		jms.limiter = defaultRateLimiter
	}
	if jms.httpClient == nil {
		jms.httpClient = defaultHttpClient
	}

	return jms, nil
}

func (c *Client) User() *UserNode   { return &UserNode{client: c, api: api_users} }
func (c *Client) Asset() *AssetNode { return &AssetNode{client: c, api: api_assets} }

// request
func (c *Client) request(method, api string, playload []byte) (data []byte, code int, err error) {
	var (
		req  *http.Request
		body io.Reader
		url  string
	)

	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		if playload != nil {
			body = bytes.NewReader(playload)
		}
	case http.MethodGet, http.MethodDelete:
	default:
		return nil, 0, fmt.Errorf("not support http method: %s", method)
	}

	if url, err = gourl.JoinPath(c.jmsUrl, api); err != nil {
		return nil, 0, err
	}
	// NOTE: 这里一定要 PathUnescape, 否则 ? 会自动转换成 %3F,
	// 这样在 Jumpserver 中进行比如 Get Assets 是获取不到资产信息的.
	if url, err = gourl.PathUnescape(url); err != nil {
		return nil, 0, err
	}
	if req, err = http.NewRequest(method, url, body); err != nil {
		return nil, 0, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	req.Header.Set("User-Agent", userAgent)
	if err := c.sigAuth.Sign(req); err != nil {
		return nil, 0, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	if data, err = io.ReadAll(resp.Body); err != nil {
		return nil, 0, err
	}
	resp.Body.Close()
	return data, resp.StatusCode, nil
}

//// mustGet
//func (c *Client) mustGet(api string) []byte {
//    data, err := c.get(api)
//    if err != nil {
//        panic(err)
//    }
//    return data
//}

// get do http get Client
func (c *Client) get(api string) ([]byte, error) {
	l, err := url.JoinPath(c.jmsUrl, api)
	if err != nil {
		return nil, err
	}
	// NOTE: 这里一定要 PathUnescape, 否则 ? 会自动转换成 %3F,
	// 这样在 Jumpserver 中进行比如 Get Assets 是获取不到资产信息的.
	l, err = url.PathUnescape(l)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, l, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Date", time.Now().Format(gmtFmt))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-JMS-ORG", "00000000-0000-0000-0000-000000000002")
	if err := c.sigAuth.Sign(req); err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return body, nil
}
