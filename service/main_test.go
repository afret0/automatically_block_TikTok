package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"service/router"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sendVerificationCode?email=konglingyu2745", nil)
	router.RegisterRouter(route)
	route.ServeHTTP(w, req)
	mockResp := `{"code":1,"message":"succeed"}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, mockResp, w.Body.String())
}
