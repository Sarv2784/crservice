package entity

type HourlyUnits struct {
	Time       string `json:"time"`
	WaveHeight string `json:"wave_height"`
}

type Hourly struct {
	Time       []string  `json:"time"`
	WaveHeight []float64 `json:"wave_height"`
}

type MarineForecast struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationTimeMS     float64     `json:"generationtime_ms"`
	UTCOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               Hourly      `json:"hourly"`
}
