package pkg

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

const (
	httpTimeout                  = 5 * time.Second
	httpConnectTimeout           = 3 * time.Second
	httpKeepAlive                = 30 * time.Second * 30
	httpIdleConnTimeout          = 90 * time.Second
	httpExpectContinueTimeout    = 1 * time.Second
	httpClientMaxIdleConnections = 100
)

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   httpConnectTimeout,
				KeepAlive: httpKeepAlive,
			}).DialContext,
			MaxIdleConns:          httpClientMaxIdleConnections,
			IdleConnTimeout:       httpIdleConnTimeout,
			TLSHandshakeTimeout:   httpConnectTimeout,
			ExpectContinueTimeout: httpExpectContinueTimeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}
}
