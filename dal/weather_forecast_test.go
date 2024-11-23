package dal

import (
	"context"
	"crservice/app/config"
	"crservice/entity"
	"crservice/model"
	"reflect"
	"testing"
)

func TestNewWeatherForecastDAL(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want *WeatherForecastDAL
	}{
		{
			name: "Valid Config",
			args: args{
				cfg: &config.Config{
					MarineForecast: config.MarineForecast{
						Url: "http://mock.url/marine",
					},
				},
			},
			want: &WeatherForecastDAL{
				config: &config.Config{
					MarineForecast: config.MarineForecast{
						Url: "http://mock.url/marine",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeatherForecastDAL(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeatherForecastDAL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherForecastDAL_GetMarineWaveHeight(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		config *config.Config
	}
	type args struct {
		ctx    *context.Context
		params model.MarineWaveHeightForecastReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.MarineForecast
		wantErr bool
	}{
		{
			name: "Valid Request",
			fields: fields{
				config: &config.Config{
					MarineForecast: config.MarineForecast{
						Url: "https://marine-api.open-meteo.com/v1/marine",
					},
				},
			},
			args: args{
				ctx: &ctx,
				params: model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				},
			},
			want: &entity.MarineForecast{
				Latitude:  54.5,
				Longitude: 10.25,
			},
			wantErr: false,
		},
		{
			name: "Nil Context",
			fields: fields{
				config: &config.Config{
					MarineForecast: config.MarineForecast{
						Url: "https://marine-api.open-meteo.com/v1/marine",
					},
				},
			},
			args: args{
				ctx: nil,
				params: model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &WeatherForecastDAL{
				config: tt.fields.config,
			}
			got, err := d.GetMarineWaveHeight(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMarineWaveHeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMarineWaveHeight() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMarineWaveHeightForecastParams(t *testing.T) {
	type args struct {
		params model.MarineWaveHeightForecastReq
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "Valid Parameters",
			args: args{
				params: model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				},
			},
			want: map[string]string{
				"latitude":  "54.5",
				"longitude": "10.25",
				"hourly":    "wave_height",
			},
			wantErr: false,
		},
		{
			name: "Empty Parameters",
			args: args{
				params: model.MarineWaveHeightForecastReq{},
			},
			want:    map[string]string{"hourly": "wave_height"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateMarineWaveHeightForecastParams(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateMarineWaveHeightForecastParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMarineWaveHeightForecastParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMarineWaveHeightResponse(t *testing.T) {
	type args struct {
		err      error
		response []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.MarineForecast
		want1   *entity.MarineForecast
		wantErr bool
	}{
		{
			name: "Valid Response",
			args: args{
				err: nil,
				response: []byte(`{
					"latitude": 54.5,
					"longitude": 10.25,
					"generationtime_ms": 0.5,
					"utc_offset_seconds": 0,
					"timezone": "GMT",
					"timezone_abbreviation": "GMT",
					"elevation": 10,
					"hourly_units": {
						"time": "iso8601",
						"wave_height": "m"
					},
					"hourly": {
						"time": ["2024-11-16T00:00"],
						"wave_height": [1.2]
					}
				}`),
			},
			want: &entity.MarineForecast{
				Latitude:  54.5,
				Longitude: 10.25,
			},
			want1:   nil,
			wantErr: false,
		},
		{
			name: "Invalid JSON",
			args: args{
				err:      nil,
				response: []byte(`invalid json`),
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseMarineWaveHeightResponse(tt.args.err, tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMarineWaveHeightResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMarineWaveHeightResponse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseMarineWaveHeightResponse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
