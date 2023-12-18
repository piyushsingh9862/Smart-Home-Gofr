// device_handlers_test.go
package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "smart/models"
    "github.com/stretchr/testify/assert"
)

func TestGetDevices(t *testing.T) {
    // Create a test server
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mock MongoDB response for fetching devices
        devices := []models.Device{
            {ID: "1", Name: "Device1", Type: "Light", Owner: "user1"},
            {ID: "2", Name: "Device2", Type: "Thermostat", Owner: "user1"},
        }

        // Respond with the mock devices as JSON
        jsonDevices, err := json.Marshal(devices)
        if err != nil {
            t.Fatal(err)
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonDevices)
    }))
    defer ts.Close()

    // Initialize the DeviceHandlers with the test server URL
    handlers := &DeviceHandlers{}
    handlers.Init(ts.URL)

    // Create a request to the test server
    req, err := http.NewRequest("GET", "/devices", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Perform the request
    rr := httptest.NewRecorder()
    handlers.GetDevices(rr, req)

    // Assert HTTP status code and response body
    assert.Equal(t, http.StatusOK, rr.Code, "Unexpected status code")
    assert.Contains(t, rr.Body.String(), "Device1", "Response should contain Device1")
    assert.Contains(t, rr.Body.String(), "Device2", "Response should contain Device2")
}

// TODO: Implement similar tests for CreateDevice, UpdateDevice, and DeleteDevice

