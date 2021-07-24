// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestApiEndpoint(t *testing.T) {
	fmt.Println("Corriendo E2E test para el endpoint de prueba")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/test")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
