package hap

import (
	"fmt"
	"strconv"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
	"github.com/jeffh/airthings"

	customCharacteristics "github.com/jeffh/airthings/hap/characteristic"
)

type AirthingsAccessory struct {
	Device airthings.Device
	*accessory.A

	TemperatureSensor *service.TemperatureSensor
	Battery           *service.BatteryService
	HumiditySensor    *service.HumiditySensor
	Co2Sensor         *service.CarbonDioxideSensor
	AirQualitySensor  *service.AirQualitySensor

	AqCo2Level        *characteristic.CarbonDioxideLevel
	Co2SensorLevel    *characteristic.CarbonDioxideLevel
	VocDensity        *characteristic.VOCDensity
	RadonShortTermAvg *customCharacteristics.RadonShortTermAverage
	Pressure          *customCharacteristics.PressureCharacteristic
}

// New returns a new AirthingsAccessor for Home Accessory Protocol
func New(dev airthings.Device) *AirthingsAccessory {
	var group AirthingsAccessory
	group.Device = dev
	group.A = accessory.New(accessory.Info{
		Name:         fmt.Sprintf("%s - %s", dev.DeviceType.String(), dev.Segment.Name),
		SerialNumber: dev.SerialNumber,
		Manufacturer: "Airthings",
		Model:        dev.DeviceType.String(),
		Firmware:     "Unknown",
	}, accessory.TypeSensor)
	group.Id, _ = strconv.ParseUint(dev.SerialNumber, 10, 64)

	for _, sensorType := range dev.Sensors {
		switch sensorType {
		case airthings.Temperature:
			group.TemperatureSensor = service.NewTemperatureSensor()
			group.AddS(group.TemperatureSensor.S)
		case airthings.BatteryPercentage:
			group.Battery = service.NewBatteryService()
			group.AddS(group.Battery.S)
		case airthings.CO2:
			if group.AirQualitySensor == nil {
				group.AirQualitySensor = service.NewAirQualitySensor()
				group.AddS(group.AirQualitySensor.S)
			}
			group.Co2Sensor = service.NewCarbonDioxideSensor()
			group.AqCo2Level = characteristic.NewCarbonDioxideLevel()
			group.Co2SensorLevel = characteristic.NewCarbonDioxideLevel()
			group.AirQualitySensor.AddC(group.AqCo2Level.C)
			group.Co2Sensor.AddC(group.Co2SensorLevel.C)
			group.AddS(group.Co2Sensor.S)
			group.AirQualitySensor.AddS(group.Co2Sensor.S)
		case airthings.Humidity:
			if group.AirQualitySensor == nil {
				group.AirQualitySensor = service.NewAirQualitySensor()
				group.AddS(group.AirQualitySensor.S)
			}
			group.HumiditySensor = service.NewHumiditySensor()
			group.AddS(group.HumiditySensor.S)
			group.AirQualitySensor.AddS(group.HumiditySensor.S)
		case airthings.VOC:
			if group.AirQualitySensor == nil {
				group.AirQualitySensor = service.NewAirQualitySensor()
				group.AddS(group.AirQualitySensor.S)
			}
			group.VocDensity = characteristic.NewVOCDensity()
			group.VocDensity.SetMaxValue(65535)
			group.AirQualitySensor.AddC(group.VocDensity.C)
		case airthings.RadonShortTermAverage:
			if group.AirQualitySensor == nil {
				group.AirQualitySensor = service.NewAirQualitySensor()
				group.AddS(group.AirQualitySensor.S)
			}
			group.RadonShortTermAvg = customCharacteristics.NewRadonShortTermAverage()
			group.AirQualitySensor.AddC(group.RadonShortTermAvg.C)
		case airthings.Pressure:
			if group.AirQualitySensor == nil {
				group.AirQualitySensor = service.NewAirQualitySensor()
				group.AddS(group.AirQualitySensor.S)
			}
			group.Pressure = customCharacteristics.NewPressure()
			group.AirQualitySensor.AddC(group.Pressure.C)
		}
	}
	return &group
}

// Update takes samples fetched from GetLatestSamples and updates the accessory's data
func (a *AirthingsAccessory) Update(samples map[airthings.SensorType]interface{}) {
	airQuality := characteristic.AirQualityUnknown
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for _, sensorType := range a.Device.Sensors {
		switch sensorType {
		case airthings.Temperature:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				a.TemperatureSensor.CurrentTemperature.SetValue(value)
			}
		case airthings.BatteryPercentage:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				a.Battery.StatusLowBattery.SetValue(BatteryCharacteristics(value))
				a.Battery.BatteryLevel.SetValue(int(value))
			}
		case airthings.CO2:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				aq, level := CO2Characteristics(value)
				airQuality = max(airQuality, aq)
				a.Co2Sensor.CarbonDioxideDetected.SetValue(level)
				a.AqCo2Level.SetValue(value)
				a.Co2SensorLevel.SetValue(value)
			}
		case airthings.Humidity:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				a.HumiditySensor.CurrentRelativeHumidity.SetValue(value)
				airQuality = max(airQuality, HumidityAirQualityCharacteristics(value))
			}
		case airthings.VOC:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				airQuality = max(airQuality, VOCAirQualityCharacteristic(value))
				if temp, ok := samples[airthings.Temperature]; ok {
					if temp, ok := temp.(float64); ok {
						if pressure, ok := samples[airthings.Pressure]; ok {
							if pressure, ok := pressure.(float64); ok {
								a.VocDensity.SetValue(VOCDensityCharacteristic(value, temp, pressure))
							}
						}
					}
				}
			}
		case airthings.RadonShortTermAverage:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				a.RadonShortTermAvg.SetValue(value)
				airQuality = max(airQuality, RadonAirQualityCharacteristic(value))
			}
		case airthings.Pressure:
			value, ok := samples[sensorType]
			if !ok {
				continue
			}
			if value, ok := value.(float64); ok {
				a.Pressure.SetValue(value)
			}
		}
	}

	a.AirQualitySensor.AirQuality.SetValue(airQuality)
}
