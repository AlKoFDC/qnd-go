package main

import (
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	httptest.NewServer()
}
