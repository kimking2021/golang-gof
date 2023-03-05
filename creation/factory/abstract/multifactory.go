package abstract

import (
	"net/http"
	"net/http/httptest"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHTTPClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct {
}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	//a new mock request generator
	res := httptest.NewRecorder()
	//calling the `Result` method gives us
	//the default empty *http.Response object
	return res.Result(), nil
}

func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}
