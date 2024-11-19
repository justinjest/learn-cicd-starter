package auth

import (
	"net/http/httptest"
	"testing"
)

func TestHeaderProcessing(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "ApiKey abc")

	// Then you can test header processing
	gotValue, err := GetAPIKey(req.Header)
	if gotValue != "abc" || err != nil {
		t.Errorf("expected %v, got %v", "abc", gotValue)
	}
}
