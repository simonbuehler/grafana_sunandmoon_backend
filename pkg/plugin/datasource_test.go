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

func TestGetLatLon(t *testing.T) {
	// Initialize Datasource
	ds := &plugin.Datasource{
		Latitude:  45.0,
		Longitude: 8.9718784,
	}

	// Create a sample backend.QueryDataRequest
	query := backend.DataQuery{
		JSON: []byte(`{"latitude": "50.0", "longitude": "9.0", "target": ["moon_illumination"]}`),
	}

	// Parse the query and check latitude/longitude overrides
	lat, lon, err := ds.GetLatLon(query)

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert the overridden values are correct
	assert.Equal(t, 50.0, lat)
	assert.Equal(t, 9.0, lon)
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
