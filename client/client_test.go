package client

import (
	"fmt"
	"reflect"
	"testing"
)

var clientMock = Client{
	&Soccerama{
		APIToken:   "api_token",
		APIVersion: "api_version",
		APIHost:    "api_host",
	},
}

func TestNew(t *testing.T) {
	observed, err := New("../resources/test/config_mock.yml")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(observed, clientMock) {
		t.Errorf("observed: %+v\nexpected: %+v", clientMock.String(), clientMock.String())
	}
}

func TestString(t *testing.T) {
	expected := fmt.Sprintf("api token: %s\napi_version: %s\napi_host: %s\n",
		clientMock.APIToken, clientMock.APIVersion, clientMock.APIHost)

	observed := clientMock.String()
	if observed != expected {
		t.Errorf("observed: %q\nexpected: %q", observed, expected)
	}
}
