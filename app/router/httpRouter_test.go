package router

import (
	"crservice/app/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeRouter(t *testing.T) {
	// Define mock configuration
	mockConfig := &config.Config{
		App: config.App{
			Name:    "TestApp",
			Version: "1.0.0",
		},
		Server: config.HttpServer{
			Port: 8080,
			Host: "127.0.0.1",
		},
		GRPCServer: config.GRPCServer{
			Host: "127.0.0.1",
			Port: 50051,
		},
		Database: config.Database{
			User:     "testuser",
			Password: "testpass",
			Name:     "testdb",
			Host:     "localhost",
			Port:     5432,
		},
		MarineForecast: config.MarineForecast{
			Url: "http://mock.url/marine",
		},
	}

	// Initialize the router
	router := InitializeRouter(mockConfig)

	// Test cases
	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid GET request to /weatherforecast/marinewaveheight/",
			method:         http.MethodGet,
			url:            "/weatherforecast/marinewaveheight/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Wave height data", // Example expected response snippet
		},
		{
			name:           "Invalid route",
			method:         http.MethodGet,
			url:            "/invalid-route/",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found", // Default 404 message from Gin
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create an HTTP request
			req, _ := http.NewRequest(tt.method, tt.url, nil)

			// Create a ResponseRecorder to capture the response
			w := httptest.NewRecorder()

			// Serve the request using the initialized router
			router.ServeHTTP(w, req)

			// Validate the response status code
			assert.Equal(t, tt.expectedStatus, w.Code, "Unexpected status code")

			// Validate the response body
			assert.Contains(t, w.Body.String(), tt.expectedBody, "Unexpected response body")
		})
	}
}
