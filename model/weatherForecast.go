package model

import (
	"crservice/entity"
)

type MarineForecast struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationTimeMs     float64     `json:"generation_time_ms"`
	UTCOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               Hourly      `json:"hourly"`
}

type HourlyUnits struct {
	Time       string `json:"time"`
	WaveHeight string `json:"wave_height"`
}

type Hourly struct {
	Time       []string  `json:"time"`
	WaveHeight []float64 `json:"wave_height"`
}

type MarineWaveHeightForecastReq struct {
	Latitude  string `form:"latitude"`
	Longitude string `form:"longitude"`
}

func ParseMarineForecastEntityToModel(marineForecastEntity *entity.MarineForecast) *MarineForecast {
	if marineForecastEntity == nil {
		return nil
	}

	return &MarineForecast{
		Latitude:             marineForecastEntity.Latitude,
		Longitude:            marineForecastEntity.Longitude,
		GenerationTimeMs:     marineForecastEntity.GenerationTimeMS,
		UTCOffsetSeconds:     marineForecastEntity.UTCOffsetSeconds,
		Timezone:             marineForecastEntity.Timezone,
		TimezoneAbbreviation: marineForecastEntity.TimezoneAbbreviation,
		Elevation:            marineForecastEntity.Elevation,
		HourlyUnits: HourlyUnits{
			Time:       marineForecastEntity.HourlyUnits.Time,
			WaveHeight: marineForecastEntity.HourlyUnits.WaveHeight,
		},
		Hourly: Hourly{
			Time:       marineForecastEntity.Hourly.Time,
			WaveHeight: marineForecastEntity.Hourly.WaveHeight,
		},
	}
}
