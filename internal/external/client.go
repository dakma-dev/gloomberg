package external

import (
	"crypto/tls"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

// Creates a new Client, with reasonable defaults
func newClient() (*http.Client, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	transport := &http.Transport{
		TLSClientConfig:     tlsConfig,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     13 * time.Second,
	}

	// explicitly use http2
	_ = http2.ConfigureTransport(transport)

	client := &http.Client{
		Timeout:   13 * time.Second,
		Transport: transport,
	}

	return client, nil
}
