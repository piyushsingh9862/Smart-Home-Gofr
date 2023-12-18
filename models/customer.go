package customer

import "your_project_name/models"

type Customer struct {
    ID      string         `json:"id"`
    Name    string         `json:"name"`
    Email   string         `json:"email"`
    Devices []models.Device `json:"devices"`
}