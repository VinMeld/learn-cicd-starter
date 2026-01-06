package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("Didn't get an error!")
	}

}
