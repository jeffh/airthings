package hap

import (
	"testing"

	"github.com/brutella/hap/characteristic"
	"github.com/jeffh/airthings"
	"github.com/stretchr/testify/assert"
)

func TestNewAirthingsAccessory(t *testing.T) {
	device := airthings.Device{
		SerialNumber: "12345678",
		DeviceType:   airthings.DevWavePlus,
		Sensors: []airthings.SensorType{
			airthings.Temperature,
			airthings.Humidity,
			airthings.CO2,
			airthings.VOC,
			airthings.Pressure,
			airthings.RadonShortTermAverage,
		},
		Segment: airthings.DeviceSegment{
			Name: "Test Device",
		},
	}

	acc := New(device)
	assert.NotNil(t, acc)
	assert.Equal(t, "Wave Plus - Test Device", acc.A.Info.Name.Value())
	assert.Equal(t, "12345678", acc.A.Info.SerialNumber.Value())
	assert.Equal(t, "Airthings", acc.A.Info.Manufacturer.Value())
	assert.Equal(t, "Wave Plus", acc.A.Info.Model.Value())

	// Verify all expected sensors are created
	assert.NotNil(t, acc.TemperatureSensor)
	assert.NotNil(t, acc.HumiditySensor)
	assert.NotNil(t, acc.Co2Sensor)
	assert.NotNil(t, acc.AirQualitySensor)
	assert.NotNil(t, acc.VocDensity)
	assert.NotNil(t, acc.Pressure)
	assert.NotNil(t, acc.RadonShortTermAvg)
}

func TestUpdateAirthingsAccessory(t *testing.T) {
	device := airthings.Device{
		SerialNumber: "12345678",
		DeviceType:   airthings.DevWavePlus,
		Sensors: []airthings.SensorType{
			airthings.Temperature,
			airthings.Humidity,
			airthings.CO2,
			airthings.VOC,
			airthings.Pressure,
			airthings.RadonShortTermAverage,
		},
		Segment: airthings.DeviceSegment{
			Name: "Test Device",
		},
	}

	acc := New(device)
	samples := map[airthings.SensorType]interface{}{
		airthings.Temperature:           22.5,
		airthings.Humidity:              45.0,
		airthings.CO2:                   700.0,
		airthings.VOC:                   150.0,
		airthings.Pressure:              1013.0,
		airthings.RadonShortTermAverage: 50.0,
	}

	acc.Update(samples)

	assert.Equal(t, 22.5, acc.TemperatureSensor.CurrentTemperature.Value())
	assert.Equal(t, 45.0, acc.HumiditySensor.CurrentRelativeHumidity.Value())
	assert.Equal(t, float64(700), acc.Co2SensorLevel.Value())
	assert.Equal(t, float64(700), acc.AqCo2Level.Value())
	assert.Equal(t, characteristic.CarbonDioxideDetectedCO2LevelsNormal, acc.Co2Sensor.CarbonDioxideDetected.Value())
	assert.Equal(t, 1013.0, acc.Pressure.Value())
	assert.Equal(t, 50.0, acc.RadonShortTermAvg.Value())

	// Test air quality calculation
	assert.Equal(t, characteristic.AirQualityExcellent, acc.AirQualitySensor.AirQuality.Value())

	// Test with poor air quality values
	samples = map[airthings.SensorType]interface{}{
		airthings.CO2:                   1200.0, // High CO2
		airthings.VOC:                   2500.0, // High VOC
		airthings.Humidity:              75.0,   // High humidity
		airthings.RadonShortTermAverage: 200.0,  // High radon
	}

	acc.Update(samples)
	assert.Equal(t, characteristic.AirQualityPoor, acc.AirQualitySensor.AirQuality.Value())
}
