package main

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.dekutyavin.net/internal/assert"
)

func TestPing(t *testing.T) {
	app := &application{
		logger: slog.New(slog.DiscardHandler),
	}
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/ping", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
