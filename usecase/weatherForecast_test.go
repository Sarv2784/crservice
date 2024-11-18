package usecase

import (
	"context"
	"crservice/dal"
	"crservice/entity"
	"crservice/mocks"
	"crservice/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestNewWeatherForecastUseCase(t *testing.T) {
	type args struct {
		dal *dal.WeatherForecastDAL
	}
	tests := []struct {
		name string
		args args
		want *WeatherForecastUseCase
	}{
		{
			name: "Valid DAL",
			args: args{
				dal: &dal.WeatherForecastDAL{},
			},
			want: &WeatherForecastUseCase{
				dal: &dal.WeatherForecastDAL{},
			},
		},
		{
			name: "Nil DAL",
			args: args{
				dal: nil,
			},
			want: &WeatherForecastUseCase{
				dal: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeatherForecastUseCase(tt.args.dal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeatherForecastUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeatherForecastUseCase_GetMarineWaveHeight(t *testing.T) {
	mockDAL := mocks.NewIWeatherForecastDAL(t)
	ctx := context.Background()
	tests := []struct {
		name           string
		mockSetup      func()
		ctx            context.Context
		request        *model.MarineWaveHeightForecastReq
		expectedResult *model.MarineForecast
		expectedError  error
	}{
		{
			name: "Valid request",
			mockSetup: func() {
				mockDAL.On("GetMarineWaveHeight", mock.Anything, model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				}).Return(&entity.MarineForecast{
					Latitude:  54.5,
					Longitude: 10.25,
				}, nil).Once()
			},
			ctx: ctx,
			request: &model.MarineWaveHeightForecastReq{
				Latitude:  "54.5",
				Longitude: "10.25",
			},
			expectedResult: &model.MarineForecast{
				Latitude:  54.5,
				Longitude: 10.25,
			},
			expectedError: nil,
		},
		{
			name: "DAL returns error",
			mockSetup: func() {
				mockDAL.On("GetMarineWaveHeight", mock.Anything, model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				}).Return(nil, assert.AnError).Once()
			},
			ctx: ctx,
			request: &model.MarineWaveHeightForecastReq{
				Latitude:  "54.5",
				Longitude: "10.25",
			},
			expectedResult: nil,
			expectedError:  assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			useCase := NewWeatherForecastUseCase(mockDAL)

			result, err := useCase.GetMarineWaveHeight(&tt.ctx, tt.request)
			assert.Equal(t, tt.expectedResult, result)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
