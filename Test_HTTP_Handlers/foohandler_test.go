package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFooRR(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handleGetFoo(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", rr.Result().StatusCode)
	}

	defer rr.Result().Body.Close()
	expected := "foo"
	b, err := io.ReadAll(rr.Result().Body)

	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but got %s", expected, string(b))
	}

}

func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	expected := "foo"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but got %s", expected, string(b))
	}

}
