package api

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

func TestAPI_ListMemberAuth_Success(t *testing.T) {
	a := New("mockCorpID", "mockCorpSecret", "mockToken", "mockAESKey")

	mockTransport := &mockRoundTripper{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/cgi-bin/gettoken") {
				respBody := `{"access_token":"valid-token","expires_in":7200}`
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					Header:     make(http.Header),
				}, nil
			}

			if strings.Contains(req.URL.Path, "/cgi-bin/user/list_member_auth") {
				respBody := `{"errcode":0,"errmsg":"ok","next_cursor":"cursor2","member_auth_list":[{"open_userid":"user1"}]}`
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

	a.Client.SetHTTPClient(&http.Client{Transport: mockTransport})

	res, err := a.ListMemberAuth("cursor1", 10)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if res == nil {
		t.Fatal("Expected non-nil result")
	}

	if res.NextCursor != "cursor2" {
		t.Errorf("Expected next_cursor to be 'cursor2', got %s", res.NextCursor)
	}

	if len(res.MemberAuthList) != 1 || res.MemberAuthList[0].OpenUserid != "user1" {
		t.Errorf("Unexpected member auth list content: %v", res.MemberAuthList)
	}
}
