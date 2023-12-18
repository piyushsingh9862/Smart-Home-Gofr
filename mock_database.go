// mock_database.go
package handlers

import (
    "context"

    "your_project_name/models"
)

// MockDatabase is a mock implementation of DeviceDatabase for testing
type MockDatabase struct {
    Devices []models.Device
}

// GetDevices retrieves a list of devices (mock implementation)
func (m *MockDatabase) GetDevices(ctx context.Context) ([]models.Device, error) {
    return m.Devices, nil
}

// CreateDevice creates a new smart home device (mock implementation)
func (m *MockDatabase) CreateDevice(ctx context.Context, device models.Device) error {
    m.Devices = append(m.Devices, device)
    return nil
}

// UpdateDevice updates an existing smart home device (mock implementation)
func (m *MockDatabase) UpdateDevice(ctx context.Context, device models.Device) error {
    for i, d := range m.Devices {
        if d.ID == device.ID {
            m.Devices[i] = device
            return nil
        }
    }
    return nil
}

// DeleteDevice deletes an existing smart home device (mock implementation)
func (m *MockDatabase) DeleteDevice(ctx context.Context, deviceID
