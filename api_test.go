package airthings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDevices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/devices", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"devices": [
				{
					"id": "12345678",
					"deviceType": "WAVE_PLUS",
					"sensors": ["temp", "humidity", "radonShortTermAvg", "co2", "voc", "pressure"],
					"segment": {
						"id": "home",
						"name": "Home",
						"started": "2024-01-01T00:00:00Z",
						"active": true
					},
					"location": {
						"id": "bedroom",
						"name": "Bedroom"
					}
				}
			]
		}`))
	}))
	defer server.Close()

	client := &Client{
		Endpoint: server.URL,
		http:     server.Client(),
	}

	// Test listing devices
	devices, err := client.ListDevices(ListDevicesOptions{})
	assert.NoError(t, err)
	assert.Len(t, devices, 1)

	if len(devices) == 0 {
		return
	}

	device := devices[0]
	assert.Equal(t, "12345678", device.SerialNumber)
	assert.Equal(t, DevWavePlus, device.DeviceType)
	assert.Equal(t, []SensorType{
		Temperature,
		Humidity,
		RadonShortTermAverage,
		CO2,
		VOC,
		Pressure,
	}, device.Sensors)
}

func TestGetLatestSamples(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/devices/12345678/latest-samples", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"temp": 22.4,
				"humidity": 45.2,
				"co2": 800,
				"voc": 150,
				"pressure": 1013.2,
				"radonShortTermAvg": 50
			}
		}`))
	}))
	defer server.Close()

	client := &Client{
		Endpoint: server.URL,
		http:     server.Client(),
	}

	samples, err := client.GetLatestSamples(GetLatestSamplesOptions{
		SerialNumber: "12345678",
	})
	assert.NoError(t, err)
	assert.Len(t, samples, 6)

	assert.Equal(t, 22.4, samples[Temperature])
	assert.Equal(t, 45.2, samples[Humidity])
	assert.Equal(t, float64(800), samples[CO2])
	assert.Equal(t, float64(150), samples[VOC])
	assert.Equal(t, 1013.2, samples[Pressure])
	assert.Equal(t, float64(50), samples[RadonShortTermAverage])
}
