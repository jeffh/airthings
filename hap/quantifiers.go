package hap

import (
	"github.com/brutella/hap/characteristic"
)

func CO2Characteristics(value float64) (airQuality, level int) {
	switch {
	case value >= 1000:
		return characteristic.AirQualityPoor, characteristic.CarbonDioxideDetectedCO2LevelsAbnormal
	case value >= 800:
		return characteristic.AirQualityFair, characteristic.CarbonDioxideDetectedCO2LevelsNormal
	default:
		return characteristic.AirQualityExcellent, characteristic.CarbonDioxideDetectedCO2LevelsNormal
	}
}

func BatteryCharacteristics(value float64) (level int) {
	if value < 10 {
		return characteristic.StatusLowBatteryBatteryLevelLow
	} else {
		return characteristic.StatusLowBatteryBatteryLevelNormal
	}
}

func HumidityAirQualityCharacteristics(value float64) (airQuality int) {
	switch {
	case value >= 70, value < 25:
		return characteristic.AirQualityPoor
	case value >= 60, value < 30:
		return characteristic.AirQualityFair
	default:
		return characteristic.AirQualityExcellent
	}
}

func VOCAirQualityCharacteristic(value float64) (airQuality int) {
	switch {
	case value >= 2000:
		return characteristic.AirQualityPoor
	case value >= 250:
		return characteristic.AirQualityFair
	default:
		return characteristic.AirQualityExcellent
	}
}

func VOCDensityCharacteristic(voc, temperature, pressure float64) (density float64) {
	return voc * (78 / 22.41 * ((temperature + 273) / 273) * (1013 / pressure))
}

func RadonAirQualityCharacteristic(value float64) (airQuality int) {
	switch {
	case value >= 150:
		return characteristic.AirQualityPoor
	case value >= 100:
		return characteristic.AirQualityFair
	default:
		return characteristic.AirQualityExcellent
	}
}
