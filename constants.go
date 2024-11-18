package airthings

type DeviceType string

const (
	DevWave                 DeviceType = "WAVE"
	DevWaveMist             DeviceType = "WAVE_MIST"
	DevWave2                DeviceType = "WAVE_GEN2"
	DevWaveMini             DeviceType = "WAVE_MINI"
	DevWavePlus             DeviceType = "WAVE_PLUS"
	DevWaveCO2              DeviceType = "WAVE_CO2"
	DevWaveViewPlus         DeviceType = "VIEW_PLUS"
	DevWaveViewPlusBusiness DeviceType = "VIEW_PLUS_BUSINESS"
	DevWaveViewPollution    DeviceType = "VIEW_POLLUTION"
	DevWaveViewRadon        DeviceType = "VIEW_RADON"
	DevWaveViewCO2          DeviceType = "VIEW_CO2"
	DevWaveTernCO2          DeviceType = "TERN_CO2"
	DevWaveHub              DeviceType = "HUB"
	DevWaveHome             DeviceType = "HOME"
	DevWavePro              DeviceType = "PRO"
	DevWaveCloudBerry       DeviceType = "CLOUDBERRY"
	DevWaveAirtight         DeviceType = "AIRTIGHT"
	DevAggregatedGroup      DeviceType = "AGGREGATED_GROUP"
	DevZoneGroup            DeviceType = "ZONE_GROUP"
	DevBalanceControl       DeviceType = "BALANCE_CONTROL"
	DevInletAirControl      DeviceType = "INLET_AIR_CONTROL"
	DevVentController       DeviceType = "VENT_CONTROLLER"
	DevAirly                DeviceType = "AIRLY"
	DevAirlyNO2             DeviceType = "AIRLY_NO2"
	DevAirlyCO              DeviceType = "AIRLY_CO"
	DevAirlyNO              DeviceType = "AIRLY_NO"
	DevBreezometerWeather   DeviceType = "BREEZOMETER_WEATHER"
	DevBacNet               DeviceType = "BACNET"
	DevUnknown              DeviceType = "UNKNOWN"
)

func (dt DeviceType) String() string {
	switch dt {
	case DevWave:
		return "Wave"
	case DevWaveMist:
		return "Wave Mist"
	case DevWave2:
		return "Wave 2"
	case DevWaveMini:
		return "Wave Mini"
	case DevWavePlus:
		return "Wave Plus"
	case DevWaveCO2:
		return "Wave CO2"
	case DevWaveViewPlus:
		return "Wave View Plus"
	case DevWaveViewPlusBusiness:
		return "Wave View Plus Business"
	case DevWaveViewPollution:
		return "Wave View Pollution"
	case DevWaveViewRadon:
		return "Wave View Radon"
	case DevWaveViewCO2:
		return "Wave View CO2"
	case DevWaveTernCO2:
		return "Wave Tern CO2"
	case DevWaveHub:
		return "Wave Hub"
	case DevWaveHome:
		return "Wave Home"
	case DevWavePro:
		return "Wave Pro"
	case DevWaveCloudBerry:
		return "Wave CloudBerry"
	case DevWaveAirtight:
		return "Wave Airtight"
	case DevAggregatedGroup:
		return "Aggregated Group"
	case DevZoneGroup:
		return "Zone Group"
	case DevBalanceControl:
		return "Balance Controller"
	case DevInletAirControl:
		return "Inlet Air Controller"
	case DevVentController:
		return "Vent Controller"
	case DevAirly:
		return "Airly"
	case DevAirlyNO2:
		return "Airly NO2"
	case DevAirlyCO:
		return "Airly CO"
	case DevAirlyNO:
		return "Airly NO"
	case DevBreezometerWeather:
		return "Breezometer Weather"
	case DevBacNet:
		return "Bac Net"
	case DevUnknown:
		return "Unknown Device"
	default:
		return "Unknown"
	}
}

type SensorType string

const (
	RadonShortTermAverage               SensorType = "radonShortTermAvg"
	RadonLongTermAverage                           = "radonLongTermAvg"
	Temperature                                    = "temp"
	OutdoorTemperature                             = "outdoorTemp"
	Humidity                                       = "humidity"
	OutdoorHumidity                                = "outdoorHumidity"
	CO2                                            = "co2"
	VOC                                            = "voc"
	Pressure                                       = "pressure"
	OutdoorPressure                                = "outdoorPressure"
	PressureDifference                             = "pressureDifference"
	PressureDifferenceStandardDeviation            = "pressureDiffStdDev"
	PressureDifferenceMin                          = "pressureDiffMin"
	PressureDifferenceMax                          = "pressureDiffMax"
	Light                                          = "light"
	BatteryPercentage                              = "batteryPercentage"
	BatteryVoltage                                 = "batteryVoltage"
	Orientation                                    = "orientation"
	PM1                                            = "pm1"
	OutdoorPM1                                     = "outdoorPm1"
	PM25                                           = "pm25"
	OutdoorPM25                                    = "outdoorPm25"
	PM10                                           = "pm10"
	OutdoorPM10                                    = "outdoorPm10"
	Mold                                           = "mold"
	StaleAir                                       = "staleAir"
	TransmissionEfficiency                         = "transmissionEfficiency"
	VirusSurvivalRate                              = "virusSurvivalRate"
	VirusRisk                                      = "virusRisk"
	WindSpeed                                      = "windSpeed"
	WindDirection                                  = "windDirection"
	WindGust                                       = "windGust"
	DewPoint                                       = "dewPoint"
	CloudCover                                     = "cloudCover"
	Visibility                                     = "visibility"
	PrecipitationProbability                       = "precipitation_probability"
	TotalPrecipitation                             = "total_precipitation"
	OutdoorWeather                                 = "outdoorWeather"
	HourlyRadonStandardDeviation                   = "hourlyRadonStandardDeviation"
	HourlyRadon                                    = "hourlyRadon"
	EnergyWastage                                  = "energyWastage"
	EnergyScenarios                                = "energyScenarios"
	HistoricVentilationConfidence                  = "historicVentilationConfidence"
	DaytimeBaseline                                = "daytimeBaseline"
	DaytimePeak                                    = "daytimePeak"
	NightBaseline                                  = "nightBaseline"
	HistoricVentilation                            = "historicVentilation"
	VentilationRunningConfidence                   = "ventilationRunningConfidence"
	OccupantsUpperBound                            = "occupantsUpper"
	OccupantsLowerBound                            = "occupantsLower"
	Occupants                                      = "occupants"
	RelativeOccupants                              = "relativeOccupants"
	VentilationAmount                              = "ventilationAmount"
	HistoricVentilationRunning                     = "historicVentilationRunning"
	VentilationRunning                             = "ventilationRunning"
	RelativeVentilationRate                        = "relativeVentilationRate"
	Aggregated                                     = "aggregated"
	SLA                                            = "sla"
	PressureAtMinHeight                            = "pressureAtMinHeight"
	PressureAtMaxHeight                            = "pressureAtMaxHeight"
	RegulationPressure                             = "regulationPressure"
	RegulationHeight                               = "regulationHeight"
	ZeroPressureHeight                             = "zeroPressureHeight"
	TotalPowerLost                                 = "totalPowerLost"
	MoistGuard                                     = "moistGuard"
	PotentialPowerSaved                            = "potentialPowerSaved"
	PotentialPowerSavedPercent                     = "potentialPowerSavedPercent"
	ZeroHeightPercent                              = "zeroHeightPercent"
	Zone                                           = "zone"
	ControlSignal                                  = "controlSignal"
	ControlStatus                                  = "controlStatus"
	ReturnState                                    = "returnState"
	AppliedGain                                    = "appliedGain"
	LastBestControlSignal                          = "lastBestControlSignal"
	LastBestSignalError                            = "lastBestSignalError"
	LastBestControlSignalGain                      = "lastBestControlSignalGain"
	LastBestControlSignalRecorded                  = "lastBestControlSignalRecorded"
	Messages                                       = "messages"
	BalanceControl                                 = "balanceControl"
	ControlSignalSlot01                            = "controlSignalSlot01"
	ControlSignalSlot02                            = "controlSignalSlot02"
	ControlSignalSlot03                            = "controlSignalSlot03"
	ControlSignalSlot04                            = "controlSignalSlot04"
	ControlSignalSlot05                            = "controlSignalSlot05"
	ControlSignalSlot06                            = "controlSignalSlot06"
	ControlSignalSlot07                            = "controlSignalSlot07"
	ControlSignalSlot08                            = "controlSignalSlot08"
	InletAirControl                                = "inletAirControl"
	PowerVoltage                                   = "powerVoltage"
	RSRP                                           = "rsrp"
	VentController                                 = "ventController"
	SubsamplesCount                                = "subsamplesCount"
	Subsamples                                     = "subsamples"
	BalanceInfo                                    = "balanceInfo"
	OutdoorNO2                                     = "outdoorNo2"
	OutdoorO3                                      = "outdoorO3"
	OutdoorSo2                                     = "outdoorSo2"
	OutdoorCo                                      = "outdoorCo"
	OutdoorNo                                      = "outdoorNo"
	Airly                                          = "airly"
	AirlyNo2                                       = "airlyNo2"
	AirlyCo                                        = "airlyCo"
	AirlyNo                                        = "airlyNo"
	BacNet                                         = "bacnet"

	Relay = "relayDeviceType" // indicates the device that proxied this information (eg - a hub)
)
