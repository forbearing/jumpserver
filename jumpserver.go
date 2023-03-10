package jumperserver

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	apiUsers   = "/api/v1/users/users/"
	apiAssets  = "/api/v1/assets/assets/"
	apiNodes   = "/api/v1/assets/nodes/"
	apiMetrics = "/api/v1/prometheus/metrics"
)

var (
	client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	gmtFmt = "Mon, 02 Jan 2006 15:04:05 GMT"
)

type Jumpserver struct {
	jmsUrl    string
	accessKey string
	secretKey string

	auth *SigAuth
}

// New
func New(jmsUrl, accessKey, secretKey string, opts ...Option) (*Jumpserver, error) {
	jms := &Jumpserver{
		jmsUrl:    jmsUrl,
		accessKey: accessKey,
		secretKey: secretKey,
		auth: &SigAuth{
			KeyID:    accessKey,
			SecretID: secretKey,
		},
	}

	for _, fn := range opts {
		if fn == nil {
			continue
		}
		if err := fn.apply(jms); err != nil {
			return nil, err
		}
	}
	return jms, nil
}

func (j *Jumpserver) GetNodes() ([]byte, error)   { return j.get(apiNodes) }
func (j *Jumpserver) GetAssets() ([]byte, error)  { return j.get(apiAssets) }
func (j *Jumpserver) GetUsers() ([]byte, error)   { return j.get(apiUsers) }
func (j *Jumpserver) GetMetrics() ([]byte, error) { return j.get(apiMetrics) }

func (j *Jumpserver) MustGetNodes() []byte   { return j.mustGet(apiNodes) }
func (j *Jumpserver) MustGetAssets() []byte  { return j.mustGet(apiAssets) }
func (j *Jumpserver) MustGetUsers() []byte   { return j.mustGet(apiUsers) }
func (j *Jumpserver) MustGetMetrics() []byte { return j.mustGet(apiMetrics) }

// mustGet
func (j *Jumpserver) mustGet(api string) []byte {
	data, err := j.get(api)
	if err != nil {
		panic(err)
	}
	return data
}

// get
func (j *Jumpserver) get(api string) ([]byte, error) {
	l, err := url.JoinPath(j.jmsUrl, api)
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
	if err := j.auth.Sign(req); err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
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
