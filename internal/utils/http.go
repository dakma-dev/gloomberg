package utils

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

type HTTPClient struct{}

var HTTP = &HTTPClient{}

func (h *HTTPClient) Head(ctx context.Context, url string) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodHead, http.Header{}, nil, tls.VersionTLS12)
}

func (h *HTTPClient) Get(ctx context.Context, url string) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodGet, http.Header{}, nil, tls.VersionTLS13)
}

func (h *HTTPClient) GetWithTLS12(ctx context.Context, url string) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodGet, http.Header{}, nil, tls.VersionTLS12)
}

func (h *HTTPClient) GetWithTLS12AndHeader(ctx context.Context, url string, customHeader http.Header) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodGet, customHeader, nil, tls.VersionTLS12)
}

func (h *HTTPClient) GetWithHeader(ctx context.Context, url string, customHeader http.Header) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodGet, customHeader, nil, tls.VersionTLS13)
}

func (h *HTTPClient) Post(ctx context.Context, url string, payload *strings.Reader) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodPost, http.Header{}, payload, tls.VersionTLS13)
}

func (h *HTTPClient) PostWithHeader(ctx context.Context, url string, customHeader http.Header, payload *strings.Reader) (*http.Response, error) {
	return h.httpCall(ctx, url, http.MethodPost, customHeader, payload, tls.VersionTLS13)
}

// func (h *HTTPClient) httpGet(url string, customHeader http.Header, tlsVersion uint16) (*http.Response, error) {.
func (h *HTTPClient) httpCall(ctx context.Context, url string, method string, customHeader http.Header, payload *strings.Reader, tlsVersion uint16) (*http.Response, error) {
	client, err := CreateHTTPClient(viper.GetDuration("http.timeout"), tlsVersion)
	if err != nil {
		return nil, err
	}

	request, err := createRequest(ctx, strings.TrimSpace(url), method, customHeader, payload)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

// func createGetRequest(url string, customHeader http.Header) (*http.Request, error) {.
func createRequest(ctx context.Context, url string, method string, customHeader http.Header, payload *strings.Reader) (*http.Request, error) {
	var request *http.Request

	var err error

	if payload != nil {
		request, err = http.NewRequestWithContext(ctx, method, url, payload)
	} else {
		request, err = http.NewRequestWithContext(ctx, method, url, nil)
	}

	if err != nil {
		gbl.Log.Errorf("❌ error creating %+v request: %+v", method, err)

		return nil, err
	}

	header := &http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Cache-Control", "no-cache")

	for h := range customHeader {
		header.Add(h, customHeader.Get(h))
	}

	request.Header = *header

	return request, nil
}

func CreateHTTPClient(requestTimeout time.Duration, tlsVersion uint16) (*http.Client, error) {
	tlsConfig := &tls.Config{ //nolint:gosec
		MinVersion: tlsVersion,
	}

	transport := &http.Transport{
		MaxIdleConnsPerHost:   25,
		TLSClientConfig:       tlsConfig,
		ResponseHeaderTimeout: viper.GetDuration("http.timeout"),
		// IdleConnTimeout:       17 * time.Second,
		// ResponseHeaderTimeout: 7 * time.Second,
		// TLSHandshakeTimeout:   5 * time.Second,
	}

	// explicitly use http2
	if err := http2.ConfigureTransport(transport); err != nil {
		gbl.Log.Errorf("❌ error configuring http2 transport: %+v", err.Error())

		return nil, err
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   requestTimeout,
	}

	return client, nil
}
