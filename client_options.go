package jumpserver

import (
	"net/http"

	"golang.org/x/time/rate"
	"gopkg.in/twindagger/httpsig.v1"
)

type Option interface {
	apply(*Client) error
}

type optionFunc func(*Client) error

func (fn optionFunc) apply(c *Client) error {
	return fn(c)
}

// AuthWithAccessKey, Access Key 对 Http Header 进行签名
func AuthWithAccessKey(accessKey, secretKey string) Option {
	return optionFunc(func(c *Client) error {
		c.sigAuth = &SigAuth{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}
		return nil
	})
}

// AuthWithPrivateToken, Private Token 为永久 Token.
func AuthWithPrivateToken(token string) Option {
	return optionFunc(func(c *Client) error {
		return nil
	})
}

// AuthWithToken, 该 Token 为一次性 Token, 该 Token 有有效期, 过期作废.
func AuthWithToken(token string) Option {
	return optionFunc(func(c *Client) error {
		return nil
	})
}

// AuthWithSession, 登录后可以直接使用 session_id 作为认证方式
func AuthWithSession(sessionID string) Option {
	return optionFunc(func(c *Client) error {
		return nil
	})
}

// WithHTTPClient
func WithHTTPClient(httpClient *http.Client) Option {
	return optionFunc(func(c *Client) error {
		return nil
	})
}

// WithRateLimiter
func WithRateLimiter(limiter *rate.Limiter) Option {
	return optionFunc(func(c *Client) error {
		c.limiter = limiter
		return nil
	})
}

type SigAuth struct {
	AccessKey string
	SecretKey string
}

func (auth *SigAuth) Sign(r *http.Request) error {
	headers := []string{"(request-target)", "date"}
	signer, err := httpsig.NewRequestSigner(auth.AccessKey, auth.SecretKey, "hmac-sha256")
	if err != nil {
		return err
	}
	return signer.SignRequest(r, headers, nil)
}
