type DeviceHandlers struct {
	gofr.Controller
}

// GetDevices retrieves a list of devices
func (h *DeviceHandlers) GetDevices(w http.ResponseWriter, r *http.Request) {
	// Fetch devices from MongoDB
	collection := mongoClient.Database("smart_home_db").Collection("devices")

	// Define options to limit the fields returned
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"_id": 0})

	cursor, err := collection.Find(r.Context(), bson.D{}, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Decode the results into a slice of devices
	var devices []models.Device
	if err := cursor.All(r.Context(), &devices); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the list of devices as a JSON response
	h.RespondJSON(w, devices)
}

// CreateDevice creates a new smart home device
func (h *DeviceHandlers) CreateDevice(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into a device object
	var newDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&newDevice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new device into MongoDB
	collection := mongoClient.Database("smart_home_db").Collection("devices")
	_, err := collection.InsertOne(r.Context(), newDevice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the created device as a JSON response
	h.RespondJSON(w, newDevice)
}

// UpdateDevice updates an existing smart home device
func (h *DeviceHandlers) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into an updated device object
	var updatedDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&updatedDevice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the existing device in MongoDB
	collection := mongoClient.Database("smart_home_db").Collection("devices")
	filter := bson.M{"id": updatedDevice.ID}
	update := bson.M{"$set": updatedDevice}

	_, err := collection.UpdateOne(r.Context(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the updated device as a JSON response
	h.RespondJSON(w, updatedDevice)
}

// DeleteDevice deletes an existing smart home device
func (h *DeviceHandlers) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into a device object
	var deleteDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&deleteDevice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the existing device from MongoDB
	collection := mongoClient.Database("smart_home_db").Collection("devices")
	filter := bson.M{"id": deleteDevice.ID}

	_, err := collection.DeleteOne(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success message as a JSON response
	h.RespondJSON(w, map[string]string{"message": "Device deleted successfully"})
}