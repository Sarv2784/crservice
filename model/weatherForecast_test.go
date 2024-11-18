package model

import (
	"crservice/entity"
	"reflect"
	"testing"
)

func TestParseMarineForecastEntityToModel(t *testing.T) {
	type args struct {
		marineForecastEntity *entity.MarineForecast
	}
	tests := []struct {
		name string
		args args
		want *MarineForecast
	}{
		{
			name: "Valid Entity",
			args: args{
				marineForecastEntity: &entity.MarineForecast{
					Latitude:             54.5,
					Longitude:            10.25,
					GenerationTimeMS:     0.5,
					UTCOffsetSeconds:     3600,
					Timezone:             "GMT",
					TimezoneAbbreviation: "GMT",
					Elevation:            10.0,
					HourlyUnits: entity.HourlyUnits{
						Time:       "iso8601",
						WaveHeight: "m",
					},
					Hourly: entity.Hourly{
						Time:       []string{"2024-11-16T00:00"},
						WaveHeight: []float64{1.2},
					},
				},
			},
			want: &MarineForecast{
				Latitude:             54.5,
				Longitude:            10.25,
				GenerationTimeMs:     0.5,
				UTCOffsetSeconds:     3600,
				Timezone:             "GMT",
				TimezoneAbbreviation: "GMT",
				Elevation:            10.0,
				HourlyUnits: HourlyUnits{
					Time:       "iso8601",
					WaveHeight: "m",
				},
				Hourly: Hourly{
					Time:       []string{"2024-11-16T00:00"},
					WaveHeight: []float64{1.2},
				},
			},
		},
		{
			name: "Nil Entity",
			args: args{
				marineForecastEntity: nil,
			},
			want: nil,
		},
		{
			name: "Empty Entity",
			args: args{
				marineForecastEntity: &entity.MarineForecast{},
			},
			want: &MarineForecast{
				Latitude:             0,
				Longitude:            0,
				GenerationTimeMs:     0,
				UTCOffsetSeconds:     0,
				Timezone:             "",
				TimezoneAbbreviation: "",
				Elevation:            0,
				HourlyUnits: HourlyUnits{
					Time:       "",
					WaveHeight: "",
				},
				Hourly: Hourly{
					Time:       nil,
					WaveHeight: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseMarineForecastEntityToModel(tt.args.marineForecastEntity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMarineForecastEntityToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
