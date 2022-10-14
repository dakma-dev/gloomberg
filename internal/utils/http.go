package utils

import (
	"crypto/tls"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

type HTTPClient struct{}

var HTTP = &HTTPClient{}

func (h *HTTPClient) Get(url string) (*http.Response, error) {
	return h.httpGet(url, http.Header{}, tls.VersionTLS13)
}

func (h *HTTPClient) GetWithTLS12(url string) (*http.Response, error) {
	return h.httpGet(url, http.Header{}, tls.VersionTLS12)
}

func (h *HTTPClient) GetWithHeader(url string, customHeader http.Header) (*http.Response, error) {
	return h.httpGet(url, customHeader, tls.VersionTLS13)
}

func (h *HTTPClient) httpGet(url string, customHeader http.Header, tlsVersion uint16) (*http.Response, error) {
	url = ReplaceSchemeWithGateway(url)

	client, err := createHTTPClient(viper.GetDuration("http.timeout"), tlsVersion)
	if err != nil {
		return nil, err
	}

	request, err := createGetRequest(strings.TrimSpace(url), customHeader)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func createGetRequest(url string, customHeader http.Header) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		gbl.Log.Errorf("❌ error creating request for %s: %s", url, err)
		return nil, err
	}

	header := &http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Cache-Control", "no-cache")

	gbl.Log.Debugf("request: %+v | header: %+v", request, header)

	for h := range customHeader {
		header.Add(h, customHeader.Get(h))
	}

	request.Header = *header

	return request, nil
}

func createHTTPClient(requestTimeout time.Duration, tlsVersion uint16) (*http.Client, error) {
	tlsConfig := &tls.Config{MinVersion: tlsVersion}

	transport := &http.Transport{
		MaxIdleConnsPerHost: 25,
		TLSClientConfig:     tlsConfig,
		// IdleConnTimeout:       17 * time.Second,
		// ResponseHeaderTimeout: 7 * time.Second,
		// TLSHandshakeTimeout:   5 * time.Second,
	}

	// explicitly use http2
	if err := http2.ConfigureTransport(transport); err != nil {
		gbl.Log.Errorf("❌ error configuring http2 transport: %+v", err.Error())
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   requestTimeout,
	}

	return client, nil
}
