package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"service/router"
	"strings"
	"testing"
)

func init() {
	router.RegisterRouter(route)
}

func TestSendVerificationCode(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sendVerificationCode?email=kongandmarx@163.com", nil)
	//router.RegisterRouter(route)
	route.ServeHTTP(w, req)
	mockResp := `{"code":1,"message":"succeed"}`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, mockResp, w.Body.String())
}

func TestRegister(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"name":"afreto","email":"kongandmarx@163.com","password":"qpzm2745","verificationCode":967706}`))
	//router.RegisterRouter(route)
	route.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
