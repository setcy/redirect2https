package redirect2https_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunalwaysknows/redirect2https"
)

func TestDemo(t *testing.T) {
	cfg := redirect2https.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := redirect2https.New(ctx, next, cfg, "redirect2https-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
}
