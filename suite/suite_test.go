package suite

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockRoundTripper struct {
	roundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTripFunc(req)
}

func TestSuite_ProviderTokenFallback(t *testing.T) {
	s := New("mockSuiteID", "mockSuiteSecret", "mockSuiteToken", "mockAESKey")

	mockTransport := &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/cgi-bin/service/get_suite_token") {
				respBody := `{"suite_access_token":"mock-suite-token","expires_in":7200}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					Header:     make(http.Header),
				}, nil
			}

			if strings.Contains(req.URL.Path, "/cgi-bin/license/list_order") {
				// Verify that fallback token is suite_access_token when provider not set
				q := req.URL.Query()
				token := q.Get("provider_access_token")
				if token != "mock-suite-token" {
					t.Errorf("Expected fallback provider_access_token to be 'mock-suite-token', got '%s'", token)
				}

				respBody := `{"errcode":0,"errmsg":"ok","order_list":[]}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					Header:     make(http.Header),
				}, nil
			}

			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewBufferString("")),
				Header:     make(http.Header),
			}, nil
		},
	}

	s.client.SetHTTPClient(&http.Client{Transport: mockTransport})
	s.SetTicket("mockTicket")

	_, err := s.ListOrder("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestSuite_ProviderTokenConfigured(t *testing.T) {
	s := New("mockSuiteID", "mockSuiteSecret", "mockSuiteToken", "mockAESKey")

	mockTransport := &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/cgi-bin/service/get_provider_token") {
				respBody := `{"provider_access_token":"mock-provider-token","expires_in":7200}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					Header:     make(http.Header),
				}, nil
			}

			if strings.Contains(req.URL.Path, "/cgi-bin/license/list_order") {
				// Verify that configured token is mock-provider-token
				q := req.URL.Query()
				token := q.Get("provider_access_token")
				if token != "mock-provider-token" {
					t.Errorf("Expected provider_access_token to be 'mock-provider-token', got '%s'", token)
				}

				respBody := `{"errcode":0,"errmsg":"ok","order_list":[]}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					Header:     make(http.Header),
				}, nil
			}

			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewBufferString("")),
				Header:     make(http.Header),
			}, nil
		},
	}

	s.client.SetHTTPClient(&http.Client{Transport: mockTransport})
	s.SetProvider("mockCorpID", "mockProviderSecret")

	_, err := s.ListOrder("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
