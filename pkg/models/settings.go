package models

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// PluginSettings enthält jetzt Latitude und Longitude
type PluginSettings struct {	
	Latitude  *float64              `json:"latitude"`  // Latitude optional
	Longitude *float64              `json:"longitude"` // Longitude optional
}

// LoadPluginSettings lädt die Plugin-Einstellungen und validiert Latitude/Longitude
func LoadPluginSettings(source backend.DataSourceInstanceSettings) (*PluginSettings, error) {
	settings := PluginSettings{}
	// JSON-Daten aus der Datenquelle unmarshallen
	err := json.Unmarshal(source.JSONData, &settings)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal PluginSettings json: %w", err)
	}

	// Validierung der Latitude und Longitude
	if settings.Latitude != nil && (*settings.Latitude < -90 || *settings.Latitude > 90) {
		return nil, fmt.Errorf("Latitude not in range -90 to +90: %f", *settings.Latitude)
	}

	if settings.Longitude != nil && (*settings.Longitude < -360 || *settings.Longitude > 360) {
		return nil, fmt.Errorf("Longitude not in range -360 to +360: %f", *settings.Longitude)
	}

	return &settings, nil
}
