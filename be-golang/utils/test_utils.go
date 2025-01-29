package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

// CreateTestRequest membuat HTTP request untuk pengujian
func CreateTestRequest(method, url, body string) *http.Request {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// ExecuteTestRequest menjalankan request pada aplikasi Fiber
func ExecuteTestRequest(app *fiber.App, req *http.Request) (*http.Response, []byte) {
	resp, _ := app.Test(req, -1) // -1 untuk tidak membatasi timeout
	body, _ := io.ReadAll(resp.Body)
	return resp, body
}

// ExecuteTestRequestWithUserContext menjalankan request dengan context user (untuk private routes)
func ExecuteTestRequestWithUserContext(app *fiber.App, req *http.Request, userData map[string]interface{}) (*http.Response, []byte) {
	reqCtx := req.Context()
	ctx := context.WithValue(reqCtx, "userData", userData)
	req = req.WithContext(ctx)

	resp, _ := app.Test(req, -1)
	body, _ := io.ReadAll(resp.Body)
	return resp, body
}
