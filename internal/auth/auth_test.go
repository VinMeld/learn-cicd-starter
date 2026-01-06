package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("Didn't get an error!")
	}

}
