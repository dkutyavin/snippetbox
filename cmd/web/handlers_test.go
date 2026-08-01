package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.dekutyavin.net/internal/assert"
)

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	ping(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
