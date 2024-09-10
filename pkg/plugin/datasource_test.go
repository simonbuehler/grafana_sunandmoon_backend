package plugin_test

import (
	"context"
	"testing"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/simonbuehler/sunandmoon_backend/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestQueryDataMoonAltitude(t *testing.T) {
	// Create a new Datasource instance
	ds := &plugin.Datasource{
		Latitude:  45.0,
		Longitude: 8.9718784,
	}

	// Create a sample backend.QueryDataRequest
	req := &backend.QueryDataRequest{
		Queries: []backend.DataQuery{
			{
				RefID: "A",
				JSON:  []byte(`{"latitude": "45", "longitude": "8.9718784", "target": ["moon_altitude"]}`),
				TimeRange: backend.TimeRange{
					From: time.Now().Add(-24 * time.Hour),
					To:   time.Now(),
				},
				Interval: time.Duration(30 * time.Minute),
			},
		},
	}

	// Call the QueryData method
	resp, err := ds.QueryData(context.Background(), req)

	// Assert no error occurred
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// Ensure the response contains at least one frame
	assert.Len(t, resp.Responses, 1)

	// Check the frame contains the correct fields
	frame := resp.Responses["A"].Frames[0]
	assert.Equal(t, "Moon altitude", frame.Name)
	assert.Equal(t, 2, len(frame.Fields)) // Time and Value fields

	// Validate field data
	timeField := frame.Fields[0]
	valueField := frame.Fields[1]

	assert.Equal(t, "Time", timeField.Name)
	assert.Equal(t, "Value", valueField.Name)

	// Check some value in moon altitude calculation
	assert.IsType(t, time.Time{}, timeField.At(0))
	assert.IsType(t, float64(0), valueField.At(0))
}

func TestCheckHealth(t *testing.T) {
	// Initialize Datasource with valid latitude and longitude
	ds := &plugin.Datasource{
		Latitude:  45.0,
		Longitude: 8.9718784,
	}

	// Call CheckHealth
	resp, err := ds.CheckHealth(context.Background(), nil)

	// Assert no error occurred
	assert.NoError(t, err)
	assert.Equal(t, backend.HealthStatusOk, resp.Status)
	assert.Equal(t, "Datasource added successfully.", resp.Message)

	// Test with invalid latitude
	ds.Latitude = -200.0
	resp, err = ds.CheckHealth(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, backend.HealthStatusError, resp.Status)
	assert.Contains(t, resp.Message, "Latitude not in range")
}

func TestGetLatLon(t *testing.T) {
	// Simulierte Datasource-Einstellungen mit Standardwerten für Latitude und Longitude
	ds := &plugin.Datasource{
		Latitude:  51.1657, // Default Latitude (z.B. Deutschland)
		Longitude: 10.4515, // Default Longitude (z.B. Deutschland)
	}

	t.Run("should return default lat/lon when query has no overrides", func(t *testing.T) {
		// Simulierte Query ohne Overrides
		query := backend.DataQuery{
			JSON: []byte(`{}`), // Keine Latitude/Longitude in der Query
		}

		lat, lon, err := ds.GetLatLon(query)
		assert.NoError(t, err)

		// Prüfe, dass die Default-Latitude und -Longitude der Datasource verwendet werden
		assert.Equal(t, ds.Latitude, lat)
		assert.Equal(t, ds.Longitude, lon)
	})

	t.Run("should use query overrides for lat/lon", func(t *testing.T) {
		// Simulierte Query mit Overrides für Latitude und Longitude
		query := backend.DataQuery{
			JSON: []byte(`{"latitude": "40.7128", "longitude": "-74.0060"}`), // New York City Koordinaten
		}

		lat, lon, err := ds.GetLatLon(query)
		assert.NoError(t, err)

		// Prüfe, dass die Latitude/Longitude aus der Query verwendet werden
		assert.Equal(t, 40.7128, lat)
		assert.Equal(t, -74.0060, lon)
	})

	t.Run("should return an error for invalid lat/lon", func(t *testing.T) {
		// Simulierte Query mit ungültigen Werten für Latitude und Longitude
		query := backend.DataQuery{
			JSON: []byte(`{"latitude": "invalid", "longitude": "invalid"}`),
		}

		_, _, err := ds.GetLatLon(query)
		assert.Error(t, err)
	})
}
