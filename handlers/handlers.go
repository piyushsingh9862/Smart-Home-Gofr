// device_handlers.go
package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "smart/models"
    "github.com/gofr-dev/gofr"
)

// DeviceHandlers represents handlers for smart home devices
type DeviceHandlers struct {
    gofr.Controller
    Database DeviceDatabase
}

// DeviceDatabase is an interface for database operations related to devices
type DeviceDatabase interface {
    GetDevices(ctx context.Context) ([]models.Device, error)
    CreateDevice(ctx context.Context, device models.Device) error
    UpdateDevice(ctx context.Context, device models.Device) error
    DeleteDevice(ctx context.Context, deviceID string) error
}

// GetDevices retrieves a list of devices
func (h *DeviceHandlers) GetDevices(w http.ResponseWriter, r *http.Request) {
    devices, err := h.Database.GetDevices(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.RespondJSON(w, devices)
}

// CreateDevice creates a new smart home device
func (h *DeviceHandlers) CreateDevice(w http.ResponseWriter, r *http.Request) {
    var newDevice models.Device
    if err := json.NewDecoder(r.Body).Decode(&newDevice); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Database.CreateDevice(r.Context(), newDevice); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.RespondJSON(w, newDevice)
}

// UpdateDevice updates an existing smart home device
func (h *DeviceHandlers) UpdateDevice(w http.ResponseWriter, r *http.Request) {
    var updatedDevice models.Device
    if err := json.NewDecoder(r.Body).Decode(&updatedDevice); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Database.UpdateDevice(r.Context(), updatedDevice); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.RespondJSON(w, updatedDevice)
}

// DeleteDevice deletes an existing smart home device
func (h *DeviceHandlers) DeleteDevice(w http.ResponseWriter, r *http.Request) {
    var deleteDevice models.Device
    if err := json.NewDecoder(r.Body).Decode(&deleteDevice); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Database.DeleteDevice(r.Context(), deleteDevice.ID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    h.RespondJSON(w, map[string]string{"message": "Device deleted successfully"})
}
