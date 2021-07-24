// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetTickets(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/ticket")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostTicket(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"user": "Sherlock", "status": "abierto"}`).
		Post("http://localhost:8080/api/ticket")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
