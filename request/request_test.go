package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "ok")
	}

	mockServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer mockServer.Close()

	rURL := mockServer.URL
	params := map[string]string{"param1": "bar"}

	body, err := Get(rURL, params)
	if err != nil {
		t.Fatal(err)
	}

	defer body.Close()
	resp, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "ok"
	observed := string(resp)

	if observed != expected {
		t.Errorf("observed %q, expected: %q", observed, expected)
	}
}
